package app

import (
	"flag"
	"harmonee/app/collections"
	"harmonee/app/controllers"
	"harmonee/app/database"
	"harmonee/app/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func Run() {
	var server = controllers.Server{}
	var appConfig = collections.AppConfig{}
	var dbConfig = collections.DBConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error on loading .env file")
	}

	appConfig.AppName = GetEnv("APP_NAME", "harmonee")
	appConfig.AppEnv = GetEnv("APP_ENV", "development")
	appConfig.AppUrl = GetEnv("APP_URL", "localhost")
	appConfig.AppPort = GetEnv("APP_PORT", "9999")

	dbConfig.DBHost = GetEnv("DB_HOST", "localhost")
	dbConfig.DBUser = GetEnv("DB_USER", "root")
	dbConfig.DBPassword = GetEnv("DB_PASSWORD", "")
	dbConfig.DBName = GetEnv("DB_NAME", "harmonee_pos")
	dbConfig.DBPort = GetEnv("DB_PORT", "3306")

	flag.Parse()
	arg := flag.Arg(0)
	if arg != "" {
		database.InitCommands(dbConfig)
	} else {
		routes.InitializeRoutes(&server)
		database.InitializeDatabase(dbConfig)
		server.Run(":" + appConfig.AppPort)
	}
}
