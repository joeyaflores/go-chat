package main

import (
	"log"
	"os"

	"github.com/joeyaflores/go-chat/internal/db"
	"github.com/joeyaflores/go-chat/internal/rbmq"
	"github.com/joeyaflores/go-chat/internal/web"
	"github.com/joho/godotenv"
)

func main() {
	log.Printf("Starting server...")
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	uri := os.Getenv("MONGODB_URI")
	db.Connect(uri)

	rabbitmqURI := os.Getenv("RABBITMQ_URI")
	go rbmq.Start(rabbitmqURI)

	web.Start()
}
