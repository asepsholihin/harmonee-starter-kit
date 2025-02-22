package database

import (
	"harmonee-pos/collection"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func ConnectDatabase() {
	var dbConfig = collection.DBConfig{}

	//db config
	dbConfig.DBHost = getEnv("DB_HOST", "localhost")
	dbConfig.DBUser = getEnv("DB_USER", "root")
	dbConfig.DBPassword = getEnv("DB_PASSWORD", "")
	dbConfig.DBName = getEnv("DB_NAME", "harmonee_pos")
	dbConfig.DBPort = getEnv("DB_PORT", "3306")

	dbLogging := getEnv("DB_LOGGING", "FALSE")
	if dbLogging == "FALSE" {
		dbConfig.DBLogging = false
	} else {
		dbConfig.DBLogging = true
	}

	dsn := "" + dbConfig.DBUser + ":" + dbConfig.DBPassword + "@tcp(" + dbConfig.DBHost + ":" + dbConfig.DBPort + ")/" + dbConfig.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
}
