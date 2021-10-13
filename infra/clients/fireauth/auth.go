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

type fireauthClient struct {
	authc *auth.Client
}

var (
	myAuthClient fireauthClient
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
	auth, err := app.Auth(context.Background())
	if err != nil {
		panic(fmt.Sprintf("firebase auth load error: %+v", err))
	}

	myAuthClient = fireauthClient{
		authc: auth,
	}
}

func FireAuth() fireauthClient {
	return myAuthClient
}

func (fc fireauthClient) Login() {
	fmt.Println("login")
}

func (fc fireauthClient) Signup(payload map[string]interface{}) (*auth.UserRecord, error) {
	params := (&auth.UserToCreate{}).
		Email(payload["Email"].(string)).
		PhoneNumber(payload["Phone"].(string)).
		Password(payload["Password"].(string)).
		DisplayName(payload["DisplayName"].(string)).
		PhotoURL(payload["ProfilePic"].(string))

	u, err := fc.authc.CreateUser(ctx, params)
	if err != nil {
		logger.Error("error creating user: %v", err)
		return nil, err
	}
	logger.InfoAsJson("Successfully created user", u)
	return u, nil
}
