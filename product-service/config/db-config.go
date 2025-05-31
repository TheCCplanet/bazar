package config

import (
	"bazar/internal/model"
)

func LoadDBConfig() model.DBConfig {
	return model.DBConfig{
		Host:     "localhost",
		Port:     "5433",
		User:     "product_user",
		Password: "secret123",
		DBName:   "product_db",
	}
}
