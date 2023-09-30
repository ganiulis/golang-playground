package middleware

import (
	"fmt"
	"go-project/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(cfg Config) {
	conn := postgres.Open(fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Host,
		cfg.User,
		cfg.Password,
		cfg.DBName,
		cfg.Port,
		cfg.SSLMode,
	))

	db, err := gorm.Open(conn, &gorm.Config{})

	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		panic(err)
	}

	fmt.Println("Migrated database")

	DB = db
}
