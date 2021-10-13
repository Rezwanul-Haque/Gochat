package fireauth

import (
	"context"
	"fmt"
	"gochat/infra/config"
	"gochat/infra/logger"
	"path/filepath"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

var authc *auth.Client

func Init() {
	var err error
	ctx := context.Background()

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
	auth, err := app.Auth(context.Background())
	if err != nil {
		panic(fmt.Sprintf("firebase auth load error: %+v", err))
	}

	authc = auth
}

func FireAuth() *auth.Client {
	return authc
}
