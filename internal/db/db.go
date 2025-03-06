package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB(sqlitePath string) {
	var err error
	DB, err = sql.Open("sqlite3", sqlitePath)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Unable to reach database:", err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS packs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		size INTEGER NOT NULL UNIQUE
	);`

	if _, err = DB.Exec(createTable); err != nil {
		log.Fatal("Failed to create packs table:", err)
	}

	SeedDefaultPackSizes()
}

// SeedDefaultPackSizes inserts default pack sizes into the database if none exist.
// This should be replaced with a proper migration system in a production environment.
func SeedDefaultPackSizes() {
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM packs").Scan(&count)
	if err != nil {
		log.Fatal("Failed to check pack count:", err)
	}

	if count == 0 {
		_, err = DB.Exec(`
		INSERT INTO packs (size) VALUES 
		(250), (500), (1000), (2000), (5000)
		`)
		if err != nil {
			log.Fatal("Failed to seed default pack sizes:", err)
		}
		log.Println("Seeded default pack sizes.")
	} else {
		log.Println("Default pack sizes already exist.")
	}
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
