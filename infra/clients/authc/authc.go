package authc

import (
	"context"
	"gochat/app/domain"
	"gochat/infra/config"
	"net"
	"net/http"
	"time"
)

const (
	CONTENT_TYPE_JSON         = "application/json"
	CONTENT_TYPE_URL_ENCODING = "application/x-www-form-urlencoded"
)

var (
	myAuthClient *authClient
	ctx          context.Context
)

func NewAuthClient() domain.IAuth {
	if config.App().AuthClientType == "firebase" {
		connectFirebase()
	}
	return &authClient{}
}

func authHttpClient() *http.Client {
	timeout := config.Auth().Firebase.Timeout * time.Second
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

func Auth() *authClient {
	return myAuthClient
}
