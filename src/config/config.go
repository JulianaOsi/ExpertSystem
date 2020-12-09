package config

import (
	"expertSystem/src/store"
)

var DBPassword = ""

type Config struct {
	DB *store.ConfigDB
}

func ReadConfig() *Config {
	db := &store.ConfigDB{
		Host:     "127.0.0.1",
		Port:     "5432",
		Name:     "expert",
		User:     "postgres",
		Password: DBPassword,
	}

	return &Config{DB: db}
}
