package config

import "os"

var Env Config

type Config struct {
	AppName   string
	AppEnv    string
	AppUrl    string
	SecretKey string
	DBPath    string
}

func InitConfig() {
	Env = Config{
		AppName:   os.Getenv("APP_NAME"),
		AppEnv:    os.Getenv("APP_ENV"),
		AppUrl:    os.Getenv("APP_URL"),
		SecretKey: os.Getenv("SECRET_KEY"),
		DBPath:    os.Getenv("DB_PATH"),
	}
}
