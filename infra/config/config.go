package config

import (
	"encoding/json"
	"fmt"
	"gochat/infra/logger"
	"log"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

type AppConfig struct {
	Name string
	Port string
}

type FireBaseConfig struct {
	CredentialFilePath string
}

type Config struct {
	App      *AppConfig
	FireBase *FireBaseConfig
}

var config Config

func App() *AppConfig {
	return config.App
}

func Firebase() *FireBaseConfig {
	return config.FireBase
}

func LoadConfig() {
	setDefaultConfig()

	_ = viper.BindEnv("CONSUL_URL")
	_ = viper.BindEnv("CONSUL_PATH")

	consulURL := viper.GetString("CONSUL_URL")
	consulPath := viper.GetString("CONSUL_PATH")

	if consulURL != "" && consulPath != "" {
		_ = viper.AddRemoteProvider("consul", consulURL, consulPath)

		viper.SetConfigType("json")
		err := viper.ReadRemoteConfig()

		if err != nil {
			log.Println(fmt.Sprintf("%s named \"%s\"", err.Error(), consulPath))
		}

		config = Config{}

		if err := viper.Unmarshal(&config); err != nil {
			panic(err)
		}

		if r, err := json.MarshalIndent(&config, "", "  "); err == nil {
			fmt.Println(string(r))
		}
	} else {
		logger.Info("CONSUL_URL or CONSUL_PATH missing! Serving with default config...")
	}
}

func setDefaultConfig() {
	config.App = &AppConfig{
		Name: "gochat",
		Port: "8080",
	}

	config.FireBase = &FireBaseConfig{
		CredentialFilePath: "fb-svc-key.json",
	}
}
