package handlers

import (
	"context"
	"fmt"
	"url-shortner-cli/config"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func RemoveURL(shortCode string){

	collection := config.DB.Database("urlshortener").Collection("urls")

	_, err := collection.DeleteOne(context.TODO(), bson.D{{Key: "short_code", Value: shortCode}})

	if err != nil{
		fmt.Println("Error removing URL:", err)
		return
	}

	fmt.Println("URL removed successfully")
}
