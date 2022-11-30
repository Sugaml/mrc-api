package config

import (
	"fmt"
	"os"
)

type Config struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSource string `mapstructure:"DB_SOURCE"`
}

const (
	DB_HOST     = "127.0.0.1"
	DB_DRIVER   = "postgres"
	DB_USER     = "root"
	DB_PASSWORD = "secret"
	DB_NAME     = "simple"
	DB_PORT     = "5432"
)

func LoadConfig() (config Config, err error) {
	config.DBDriver = os.Getenv("DB_DRIVER")
	if config.DBDriver == "" {
		config.DBDriver = DB_DRIVER
	}
	if os.Getenv("DB_HOST") != "" {
		config.DBSource = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"))

	} else {
		config.DBSource = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DB_HOST, DB_PORT, DB_USER, DB_NAME, DB_PASSWORD)
	}
	return config, err
}
