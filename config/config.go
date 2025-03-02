package config

import (
	"os"
	"strings"
)

var Env config

type config struct {
	AppName          string
	AppEnv           string
	AppDebug         bool
	AppUrl           string
	CorsAllowOrigins []string
	SecretKey        string
	DBPath           string
}

func InitConfig() {
	Env = config{
		AppName:          os.Getenv("APP_NAME"),
		AppEnv:           os.Getenv("APP_ENV"),
		AppDebug:         os.Getenv("APP_DEBUG") == "true",
		AppUrl:           os.Getenv("APP_URL"),
		CorsAllowOrigins: strings.Split(os.Getenv("CORS_ALLOW_ORIGINS"), ","),
		SecretKey:        os.Getenv("SECRET_KEY"),
		DBPath:           os.Getenv("DB_PATH"),
	}
}
