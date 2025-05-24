package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB
var err error

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func InitDB(cfg DBConfig) error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)

	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	if DB.Ping() != nil {
		log.Println("Data Base is OFF")
	}
	log.Println("DataBase successfuly runing ✅")
	return DB.Ping()
}
