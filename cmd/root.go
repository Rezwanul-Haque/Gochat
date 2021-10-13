package cmd

import (
	"fmt"
	"gochat/infra/clients/fireauth"
	"gochat/infra/config"
	"gochat/infra/logger"
	"os"

	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use:   "gochat code",
		Short: "implementing gochat architecture in golang",
	}
)

func init() {
	RootCmd.AddCommand(serveCmd)
}

// Execute executes the root command
func Execute() {
	config.LoadConfig()
	fireauth.Init()

	logger.Info("about to start the application")

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
