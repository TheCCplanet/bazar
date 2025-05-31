package db

import (
	"bazar/internal/model"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB
var err error

func InitDB(cfg model.DBConfig) error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	return DB.Ping()
}
