package database

import (
	"fmt"
	"harmonee/app/collections"
	"harmonee/app/database/seeders"
	"harmonee/app/models"
	"log"
	"os"

	"github.com/urfave/cli"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Connect *gorm.DB

func InitializeDatabase(dbConfig collections.DBConfig) {
	fmt.Println("Initialize database...")
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBHost, dbConfig.DBPort, dbConfig.DBName)
	Connect, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
}

func dbMigrate() {
	for _, model := range models.RegisterModels() {
		err := Connect.Debug().AutoMigrate(model.Model)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Database migrated successfully")
}

func InitCommands(dbConfig collections.DBConfig) {
	InitializeDatabase(dbConfig)

	cmdApp := cli.NewApp()
	cmdApp.Commands = []cli.Command{
		{
			Name: "db:migrate",
			Action: func(c *cli.Context) error {
				dbMigrate()
				return nil
			},
		},
		{
			Name: "db:seed",
			Action: func(c *cli.Context) error {
				err := seeders.DBSeed(Connect)
				if err != nil {
					log.Fatal(err)
				}
				return nil
			},
		},
	}

	err := cmdApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
