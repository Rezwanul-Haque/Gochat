package authc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"gochat/app/serializers"
	"gochat/infra/config"
	"gochat/infra/errors"
	"gochat/infra/logger"
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"
	"strings"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

type authClient struct {
	authc *auth.Client
	httpc *http.Client
}

func connectFirebase() {
	var err error
	ctx = context.Background()

	absPath, err := filepath.Abs(config.Auth().Firebase.ServiceAccountFilePath)
	if err != nil {
		panic("unable to load service account keys file")
	}

	logger.Info("absPath: " + absPath)

	opts := option.WithCredentialsFile(absPath)
	app, err := firebase.NewApp(ctx, nil, opts)
	if err != nil {
		panic(fmt.Sprintf("error initializing app: %v", err))
	}

	logger.Info("firebase connection established...")

	//Firebase Auth
	auth, err := app.Auth(ctx)
	if err != nil {
		panic(fmt.Sprintf("firebase auth load error: %+v", err))
	}

	myAuthClient = &authClient{
		authc: auth,
		httpc: authHttpClient(),
	}
}

func (ac authClient) Signup(email string, password string) (*serializers.LoginResp, *errors.RestErr) {
	payload := &serializers.LoginReq{
		Email:             email,
		Password:          password,
		ReturnSecureToken: true,
	}

	byteData, _ := json.Marshal(payload)

	req := prepareFirebaseURL(config.Auth().Firebase.SignUpWithEmailAndPasswordUrl, byteData, "POST", CONTENT_TYPE_JSON)

	res, err := ac.httpc.Do(&req)
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

	var resp serializers.LoginResp

	if err := json.Unmarshal(body, &resp); err != nil {
		logger.Error("error occurred while unmarshalling", err)
		restErr := errors.NewInternalServerError(errors.ErrSomethingWentWrong)
		return nil, restErr
	}

	logger.InfoAsJson("firebase response after signup", resp)

	if resp.IDToken == "" || resp.Email == "" || resp.RefreshToken == "" {
		return nil, errors.NewConflictError("user with this email already exists")
	}

	return &resp, nil
}

func (ac authClient) Login(email string, password string) (*serializers.LoginResp, *errors.RestErr) {
	payload := &serializers.LoginReq{
		Email:             email,
		Password:          password,
		ReturnSecureToken: true,
	}

	byteData, _ := json.Marshal(payload)

	req := prepareFirebaseURL(config.Auth().Firebase.SignInWithEmailAndPasswordUrl, byteData, "POST", CONTENT_TYPE_JSON)

	res, err := ac.httpc.Do(&req)
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

	var resp serializers.LoginResp

	if err := json.Unmarshal(body, &resp); err != nil {
		logger.Error("error occurred while unmarshalling", err)
		restErr := errors.NewInternalServerError(errors.ErrSomethingWentWrong)
		return nil, restErr
	}

	if resp.IDToken == "" || resp.Email == "" || resp.RefreshToken == "" {
		return nil, errors.NewUnauthorizedError("email/password is not correct")
	}

	logger.InfoAsJson("firebase response after login", resp)

	return &resp, nil
}

func (ac authClient) RefreshToken(rtoken string) (*serializers.RefreshTokenResp, *errors.RestErr) {
	payload := &serializers.RefreshTokenReq{
		GranTType:    "refresh_token",
		RefreshToken: rtoken,
	}

	data := url.Values{}
	data.Set("grant_type", payload.GranTType)
	data.Set("refresh_token", payload.RefreshToken)
	encodedData := data.Encode()

	reqURL, _ := url.Parse(config.Auth().Firebase.RefreshTokenUrl)

	// adding query params
	q := url.Values{}
	q.Add("key", config.Auth().Firebase.ApiKey)
	reqURL.RawQuery = q.Encode()

	req, err := http.NewRequest("POST", reqURL.String(), strings.NewReader(encodedData))
	if err != nil {
		logger.Error("firebase requesting error", err)
		restErr := errors.NewInternalServerError(errors.ErrSomethingWentWrong)
		return nil, restErr
	}
	req.Header.Add("Content-Type", CONTENT_TYPE_URL_ENCODING)
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	res, err := ac.httpc.Do(req)
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

		var resp serializers.RefreshTokenResp

		if err := json.Unmarshal(body, &resp); err != nil {
			logger.Error("error occurred while unmarshalling", err)
			restErr := errors.NewInternalServerError(errors.ErrSomethingWentWrong)
			return nil, restErr
		}

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

func (fc authClient) VerifyToken(idToken string) *errors.RestErr {
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
	q.Add("key", config.Auth().Firebase.ApiKey)
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
