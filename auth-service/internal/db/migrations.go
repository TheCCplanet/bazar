package db

import "log"

func RunMigrations() error {
	CreateUsersTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	username VARCHAR(255) UNIQUE NOT NULL,
	password_hash TEXT NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP
	);`

	_, err := DB.Exec(CreateUsersTableQuery)
	if err != nil {
		log.Println("DataBase Error: Faild to Create users Table", err)
	}
	return err
}
