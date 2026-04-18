package handlers

import (
	"context"
	"log"
	"net/url"
	"os"
	"strings"

	// "encoding/json"
	"fmt"
	"math/rand"
	"time"

	"url-shortner-cli/config"
	"url-shortner-cli/models"

	"github.com/denisbrodbeck/machineid"
	"github.com/joho/godotenv"
)


func AddURL(url string){

	if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using system environment variables")
    }
	userId, Iderr := machineid.ID()
	if Iderr != nil {
		log.Fatal(Iderr)
	}

	if !IsValidHTTPSURL(url){
		fmt.Println("Invalid URL")
		return
	}
	collection := config.DB.Database("urlshortener").Collection("urls")

	code := generateCode()

	doc := models.URL{
		ShortCode: code,
		UserID: userId,
		LongURL: url,
		CreatedAt: time.Now(),
	}

	_, err := collection.InsertOne(context.TODO(), doc)

	if err != nil{
		fmt.Println("Error adding URL:", err)
		return
	}


	fmt.Println("URL added successfully !")
	Server := os.Getenv("SERVER")
	fmt.Println(Server + code)
}

// functions
func generateCode() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	code := make([]byte, 6)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}
	return string(code)
}


func IsValidHTTPSURL(input string) bool {
	if input == "" {
		return false
	}

	// Basic check: must start with https://
	if !strings.HasPrefix(strings.ToLower(input), "https://") {
		return false
	}

	// Parse the URL properly
	parsedURL, err := url.Parse(input)
	if err != nil {
		return false
	}

	// Additional validations
	if parsedURL.Scheme != "https" {
		return false
	}

	if parsedURL.Host == "" {
		return false // must have a domain/host
	}

	return true
}
