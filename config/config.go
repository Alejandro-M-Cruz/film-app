package config

import (
	"os"
	"strings"
)

var Env env

type env struct {
	AppName          string
	AppEnv           string
	AppDebug         bool
	AppPort          string
	AppUrl           string
	CorsAllowOrigins []string
	SecretKey        string
	DBPath           string
}

func InitConfig() {
	Env = env{
		AppName:          os.Getenv("APP_NAME"),
		AppEnv:           os.Getenv("APP_ENV"),
		AppDebug:         os.Getenv("APP_DEBUG") == "true",
		AppPort:          os.Getenv("APP_PORT"),
		AppUrl:           os.Getenv("APP_URL"),
		CorsAllowOrigins: strings.Split(os.Getenv("CORS_ALLOW_ORIGINS"), ","),
		SecretKey:        os.Getenv("SECRET_KEY"),
		DBPath:           os.Getenv("DB_PATH"),
	}
}
