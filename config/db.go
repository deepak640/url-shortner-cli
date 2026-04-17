package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var DB *mongo.Client


func Connect(){
	opts := options.Client().ApplyURI("mongodb://localhost:27017")

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
