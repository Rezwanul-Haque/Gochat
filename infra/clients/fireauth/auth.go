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
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
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

	req := prepareFirebaseURL(config.Firebase().SignUpWithEmailAndPasswordUrl, byteData, "POST")

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

	req := prepareFirebaseURL(config.Firebase().SignInWithEmailAndPasswordUrl, byteData, "POST")

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

func (fc fireauthClient) VerifyToken(idToken string) *errors.RestErr {
	token, err := fc.authc.VerifyIDToken(ctx, idToken)
	if err != nil {
		logger.ErrorAsJson("error verifying ID token", err)
		return errors.NewUnauthorizedError("access forbidden")
	}

	logger.InfoAsJson("Verified ID token", token)
	return nil
}

func prepareFirebaseURL(baseUrl string, body []byte, method string) http.Request {
	reqURL, _ := url.Parse(baseUrl)

	// adding query params
	q := url.Values{}
	q.Add("key", config.Firebase().ApiKey)
	reqURL.RawQuery = q.Encode()

	req := http.Request{
		Method: method,
		URL:    reqURL,
		Header: map[string][]string{
			"Content-Type": {"application/json"},
		},
		Body:          ioutil.NopCloser(bytes.NewReader(body)),
		ContentLength: int64(len(body)),
	}

	return req
}
