package config

import (
	"github.com/joho/godotenv"
	"julo_test/lib/db"

	"os"
)

type Config struct{}

func (c *Config) InitEnv() error {
	return godotenv.Load(".env")
}

func (c *Config) GetDBConfig() db.ConfigDB {
	return db.ConfigDB{
		Driver:   os.Getenv("DB_DRIVER"),
		DBName:   os.Getenv("DB_NAME"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PWD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
	}
}

func (c *Config) CatchError(err error) {
	if err != nil {
		panic(any(err))
	}
}
