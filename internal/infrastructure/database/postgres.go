// infrastructure/database/postgres.go

package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"quickflow/config"
)

var DB *gorm.DB

func InitDatabase(cfg *config.Config) {
	var err error
	dsn := cfg.GetDSN()

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connection established")
}

func GetDB() *gorm.DB {
	return DB
}
