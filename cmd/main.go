package main

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"

	"github.com/malailiyati/todoList/internals/configs"
	"github.com/malailiyati/todoList/internals/routers"
)

func main() {
	db := configs.InitDB()
	log.Println("Database connected")

	r := routers.InitRouter(db)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running at :%s\n", port)
	r.Run(":" + port)
}
