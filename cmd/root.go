package cmd

import (
	"fmt"
	"gochat/infra/clients/authc"
	"gochat/infra/config"
	"gochat/infra/logger"
	"os"

	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use:   "gochat code",
		Short: "implementing gochat video app in golang",
	}
)

func init() {
	RootCmd.AddCommand(serveCmd)
}

// Execute executes the root command
func Execute() {
	config.LoadConfig()
	logger.NewLoggerClient()
	authc.NewAuthClient() // initialize admin sdk

	logger.Info("about to start the application")

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
