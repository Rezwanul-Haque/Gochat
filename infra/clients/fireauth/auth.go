package fireauth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"gochat/infra/config"
	"gochat/infra/errors"
	"gochat/infra/logger"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

const (
	CONTENT_TYPE_JSON         = "application/json"
	CONTENT_TYPE_URL_ENCODING = "application/x-www-form-urlencoded"
)

type fireauthClient struct {
	authc *auth.Client
	httpc *http.Client
}

type LoginReq struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ReturnSecureToken bool   `json:"returnSecureToken"`
}

type LoginResp struct {
	IDToken      string `json:"idToken"`
	Email        string `json:"email"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
	LocalID      string `json:"localId"`
}

type RefreshTokenReq struct {
	GranTType    string `json:"grant_type"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenResp struct {
	ExpiresIn    string `json:"expires_in"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	IDToken      string `json:"id_token"`
	UserID       string `json:"user_id"`
	ProjectID    string `json:"-"`
}

var (
	myAuthClient *fireauthClient
	ctx          context.Context
)

func Init() {
	var err error
	ctx = context.Background()

	absPath, err := filepath.Abs(config.Firebase().CredentialFilePath)
	if err != nil {
		panic("unable to load service account keys file")
	}

	opts := option.WithCredentialsFile(absPath)
	app, err := firebase.NewApp(ctx, nil, opts)
	if err != nil {
		logger.Error("error initializing app: ", err)
	}

	logger.Info("firebase connection established...")

	//Firebase Auth
	auth, err := app.Auth(ctx)
	if err != nil {
		panic(fmt.Sprintf("firebase auth load error: %+v", err))
	}

	myAuthClient = &fireauthClient{
		authc: auth,
		httpc: ConnectFirebase(),
	}
}

func ConnectFirebase() *http.Client {
	timeout := config.Firebase().Timeout * time.Second
	var netTransport = &http.Transport{
		DialContext:         (&net.Dialer{Timeout: timeout, KeepAlive: time.Minute}).DialContext,
		TLSHandshakeTimeout: timeout,
		MaxIdleConnsPerHost: 10,
	}

	httpc := &http.Client{
		Timeout:   timeout,
		Transport: netTransport,
	}

	return httpc

}

func FireAuth() *fireauthClient {
	return myAuthClient
}

func (fc fireauthClient) Signup(email string, password string) (*LoginResp, *errors.RestErr) {
	payload := &LoginReq{
		Email:             email,
		Password:          password,
		ReturnSecureToken: true,
	}

	byteData, _ := json.Marshal(payload)

	req := prepareFirebaseURL(config.Firebase().SignUpWithEmailAndPasswordUrl, byteData, "POST", CONTENT_TYPE_JSON)

	res, err := fc.httpc.Do(&req)
	if err != nil {
		logger.Error("firebase requesting error", err)
		restErr := errors.NewInternalServerError(errors.ErrSomethingWentWrong)
		return nil, restErr
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Error("reading response body from firebase", err)
		restErr := errors.NewInternalServerError(errors.ErrSomethingWentWrong)
		return nil, restErr
	}

	var resp LoginResp

	json.Unmarshal(body, &resp)

	logger.InfoAsJson("firebase response after signup", resp)

	return &resp, nil
}

func (fc fireauthClient) Login(email string, password string) (*LoginResp, *errors.RestErr) {
	payload := &LoginReq{
		Email:             email,
		Password:          password,
		ReturnSecureToken: true,
	}

	byteData, _ := json.Marshal(payload)

	req := prepareFirebaseURL(config.Firebase().SignInWithEmailAndPasswordUrl, byteData, "POST", CONTENT_TYPE_JSON)

	res, err := fc.httpc.Do(&req)
	if err != nil {
		logger.Error("firebase requesting error", err)
		restErr := errors.NewInternalServerError(errors.ErrSomethingWentWrong)
		return nil, restErr
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Error("reading response body from firebase", err)
		restErr := errors.NewInternalServerError(errors.ErrSomethingWentWrong)
		return nil, restErr
	}

	var resp LoginResp

	json.Unmarshal(body, &resp)

	logger.InfoAsJson("firebase response after login", resp)

	return &resp, nil
}

func (fc fireauthClient) RefreshToken(rtoken string) (*RefreshTokenResp, *errors.RestErr) {
	payload := &RefreshTokenReq{
		GranTType:    "refresh_token",
		RefreshToken: rtoken,
	}

	data := url.Values{}
	data.Set("grant_type", payload.GranTType)
	data.Set("refresh_token", payload.RefreshToken)
	encodedData := data.Encode()

	reqURL, _ := url.Parse(config.Firebase().RefreshTokenUrl)

	// adding query params
	q := url.Values{}
	q.Add("key", config.Firebase().ApiKey)
	reqURL.RawQuery = q.Encode()

	req, err := http.NewRequest("POST", reqURL.String(), strings.NewReader(encodedData))
	if err != nil {
		logger.Error("firebase requesting error", err)
		restErr := errors.NewInternalServerError(errors.ErrSomethingWentWrong)
		return nil, restErr
	}
	req.Header.Add("Content-Type", CONTENT_TYPE_URL_ENCODING)
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	res, err := fc.httpc.Do(req)
	if err != nil {
		logger.Error("firebase requesting error", err)
		restErr := errors.NewInternalServerError(errors.ErrSomethingWentWrong)
		return nil, restErr
	}

	if res.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			logger.Error("reading response body from firebase", err)
			restErr := errors.NewInternalServerError(errors.ErrSomethingWentWrong)
			return nil, restErr
		}

		var resp RefreshTokenResp

		json.Unmarshal(body, &resp)

		logger.InfoAsJson("firebase response after renew refresh token", resp)

		return &resp, nil
	}

	logger.Info(fmt.Sprintf("firebase response status code: %+v", res.StatusCode))
	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	newStr := buf.String()
	logger.Info(fmt.Sprintf("firebase response after: %+v", newStr))

	restErr := errors.NewUnauthorizedError(errors.ErrInvalidRefreshToken)
	return nil, restErr
}

func (fc fireauthClient) VerifyToken(idToken string) *errors.RestErr {
	token, err := fc.authc.VerifyIDToken(ctx, idToken)
	if err != nil {
		logger.ErrorAsJson("error verifying ID token", err)
		return errors.NewUnauthorizedError("unauthorized action")
	}

	logger.InfoAsJson("verified ID token", token)
	return nil
}

func prepareFirebaseURL(baseUrl string, body []byte, method, contentType string) http.Request {
	reqURL, _ := url.Parse(baseUrl)

	// adding query params
	q := url.Values{}
	q.Add("key", config.Firebase().ApiKey)
	reqURL.RawQuery = q.Encode()

	req := http.Request{
		Method: method,
		URL:    reqURL,
		Header: map[string][]string{
			"Content-Type": {contentType},
		},
		Body:          ioutil.NopCloser(bytes.NewReader(body)),
		ContentLength: int64(len(body)),
	}

	return req
}
