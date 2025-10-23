package configs

import (
	"fmt"
	"log"
	"os"

	// "github.com/malailiyati/todoList/internals/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")
	ssl := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, pass, name, port, ssl)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect DB: %v", err)
	}

	log.Println("Connected to PostgreSQL")
	return db
}

// func MigrateDB(db *gorm.DB) {
// 	err := db.AutoMigrate(&models.Category{}, &models.Todo{})
// 	if err != nil {
// 		log.Fatalf("Migration failed: %v", err)
// 	}
// 	log.Println("Migration completed")
// }
