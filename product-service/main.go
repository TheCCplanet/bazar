package main

import (
	"bazar/config"
	"bazar/internal/db"
	"bazar/internal/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	err := db.InitDB(config.LoadDBConfig())
	if err != nil {
		log.Println("Data base is Down 💤:\n", err)
		return

	}
	log.Println("DataBase is succesfuly up 🟩")

	if err = db.RunMigrations(); err != nil {
		log.Println("Migration Error:", err)
		return
	}
	os.MkdirAll("./uploads", os.ModePerm)

	http.HandleFunc("/products", handlers.ProductHandler)
	http.HandleFunc("/upload", handlers.UploadHandler)
	log.Println("Product Service succesfuly runed 📦")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
