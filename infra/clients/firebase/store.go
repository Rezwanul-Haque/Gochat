package firbase

import (
	"context"
	"fmt"
	"gochat/infra/config"
	"gochat/infra/logger"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func Init() {
	var err error
	ctx := context.Background()

	opts := option.WithCredentialsFile(config.Cloud().CredentialPath)
	app, err := firebase.NewApp(ctx, nil, opts)
	if err != nil {
		logger.Error("error initializing app: ", err)
	}

	fmt.Println(app)
}
