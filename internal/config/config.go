package config

import (
	"os"
)

var Config struct {
	Port     string
	SQLiteDB string
}

func LoadConfig() {
	Config.Port = getEnv("PORT", "8080")
	Config.SQLiteDB = getEnv("SQLITE_DB_PATH", "./packs.db")
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
