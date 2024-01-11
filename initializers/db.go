package initializers

import (
	"fmt"
	"log"
	"os"

	"go-fiber-service/models"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectPostgresDB(config *Config) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database! \n", err.Error())
		os.Exit(1)
	}

	DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	DB.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running Migrations")
	err = DB.AutoMigrate(&models.Note{})
	if err != nil {
		log.Fatal("Failed to migrations to the Note table! \n", err.Error())
	}

	log.Println("🚀 Connected Successfully to the Database")
}

func ConnectMariaDB(config *Config) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUserName, config.DBUserPassword, config.DBHost, config.DBPort, config.DBName)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database!\n", err.Error())
		os.Exit(1)
	}

	log.Println("Running Migrations")
	// Modify the model to your MariaDB model (e.g., &models.Note{})
	err = DB.AutoMigrate(&models.Note{})
	if err != nil {
		log.Fatal("Failed to migrations Database!\n", err.Error())
		os.Exit(1)
	}
	err = DB.AutoMigrate(&models.Pokemon{})
	if err != nil {
		log.Fatal("Failed to migrations to the Pokemon table! \n", err.Error())
	}

	DB.Logger = logger.Default.LogMode(logger.Info)

	log.Println("🚀 Connected Successfully to the Database")
}
