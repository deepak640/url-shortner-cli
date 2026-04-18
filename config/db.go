package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var DB *mongo.Client

func Connect() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017" // Default value
	}

	opts := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(opts)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Database unreachable: %v", err)
	}

	DB = client
}
