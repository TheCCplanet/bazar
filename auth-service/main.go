package main

import (
	"bazar/config"
	"bazar/internal/db"
	"bazar/internal/handlers"

	"log"
	"net/http"
)

func main() {

	cfg := config.Load()
	err := db.InitDB(db.DBConfig{
		Host:     cfg.Host,
		Port:     cfg.Port,
		User:     cfg.User,
		Password: cfg.Password,
		DBName:   cfg.DBName,
	})
	if err != nil {
		log.Println("innit db Error:", err)
	}

	if err = db.RunMigrations(); err != nil {
		log.Fatal("migration error:", err)
		return
	}

	http.HandleFunc("/auth", handlers.AuthHandler)
	log.Println("auth service 👮")
	err = http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("Faild to start auth Service Lisener ")
		return
	}
}
