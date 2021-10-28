package config

import (
	"encoding/json"
	"fmt"
	"gochat/infra/logger"
	"log"
	"time"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

type AppConfig struct {
	Name string
	Port string
}

type FireBaseConfig struct {
	CredentialFilePath            string
	ApiKey                        string
	SignUpWithEmailAndPasswordUrl string
	SignInWithEmailAndPasswordUrl string
	RefreshTokenUrl               string
	Timeout                       time.Duration
}

type AgoraConfig struct {
	AppID            string
	AppCertificate   string
	DefaultExpiresIn uint32
	TokenTypes       []string
}

type Config struct {
	App      *AppConfig
	FireBase *FireBaseConfig
	Agora    *AgoraConfig
}

var config Config

func App() *AppConfig {
	return config.App
}

func Firebase() *FireBaseConfig {
	return config.FireBase
}

func Agora() *AgoraConfig {
	return config.Agora
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
		CredentialFilePath:            "fb-svc-key.json",
		ApiKey:                        "firebase-web-api-key",
		SignUpWithEmailAndPasswordUrl: "https://identitytoolkit.googleapis.com/v1/accounts:signUp",
		SignInWithEmailAndPasswordUrl: "https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword",
		RefreshTokenUrl:               "https://securetoken.googleapis.com/v1/token",
		Timeout:                       10,
	}

	config.Agora = &AgoraConfig{
		AppID:            "agora-project-app-id",
		AppCertificate:   "agora-project-app-certificate",
		DefaultExpiresIn: 86400, // default expires in 86400 seconds(24 hour)
		TokenTypes:       []string{"userAccount", "uid"},
	}
}
