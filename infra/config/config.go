package config

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

type AppConfig struct {
	Name             string
	Port             string
	MetricsPort      string
	LogLevel         string
	AuthClientType   string
	RtcClientType    string
	LoggerClientType string
}

type FireBaseConfig struct {
	ServiceAccountFilePath        string
	ApiKey                        string
	SignUpWithEmailAndPasswordUrl string
	SignInWithEmailAndPasswordUrl string
	RefreshTokenUrl               string
	Timeout                       time.Duration
}

type AuthClient struct {
	Firebase *FireBaseConfig
}

type RTCClient struct {
	Agora *AgoraConfig
}

type AgoraConfig struct {
	AppID            string
	AppCertificate   string
	DefaultExpiresIn uint32
	TokenTypes       []string
}

type Config struct {
	App  *AppConfig
	Auth AuthClient
	RTC  RTCClient
}

var config Config

func App() *AppConfig {
	return config.App
}

func Auth() AuthClient {
	return config.Auth
}

func RTC() RTCClient {
	return config.RTC
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
		log.Println("CONSUL_URL or CONSUL_PATH missing! Serving with default config...")
	}
}

func setDefaultConfig() {
	config.App = &AppConfig{
		Name:             "gochat",
		Port:             "8080",
		MetricsPort:      "9080",
		LogLevel:         "info",
		AuthClientType:   "firebase",
		RtcClientType:    "agora",
		LoggerClientType: "zap",
	}

	config.Auth.Firebase = &FireBaseConfig{
		ServiceAccountFilePath:        "fb-svc-key.json",
		ApiKey:                        "<firevase-api-key>",
		SignUpWithEmailAndPasswordUrl: "https://identitytoolkit.googleapis.com/v1/accounts:signUp",
		SignInWithEmailAndPasswordUrl: "https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword",
		RefreshTokenUrl:               "https://securetoken.googleapis.com/v1/token",
		Timeout:                       10,
	}

	config.RTC.Agora = &AgoraConfig{
		AppID:            "agora-project-app-id",
		AppCertificate:   "agora-project-app-certificate",
		DefaultExpiresIn: 86400, // default expires in 86400 seconds(24 hour)
		TokenTypes:       []string{"userAccount", "uid"},
	}
}
