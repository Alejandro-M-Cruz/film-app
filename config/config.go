package config

import "os"

type Config struct {
    AppEnv string
    AppUrl string
    DBPath string
}

func InitConfig() *Config {
    return &Config{
        AppEnv: os.Getenv("APP_ENV"),
        AppUrl: os.Getenv("APP_URL"),
        DBPath: os.Getenv("DB_PATH"),
    }
}
