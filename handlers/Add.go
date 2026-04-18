package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/denisbrodbeck/machineid"
	"github.com/joho/godotenv"
)


func AddURL(url string){
	if !IsValidHTTPSURL(url){
		fmt.Println("Invalid URL")
		return
	}

	if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using system environment variables")
  }
	userId, Iderr := machineid.ID()
	if Iderr != nil {
		log.Fatal(Iderr)
	}

	Server := os.Getenv("SERVER")


	data := map[string]interface{}{
		"UserID": userId,
		"URL": url,
	}

	jsonData, err2 := json.Marshal(data)

	if err2 != nil {
		log.Fatal(err2)
	}

	response , err3 := http.Post(Server + "shorten", "application/json", bytes.NewBuffer(jsonData))
	if err3 != nil {
		log.Fatal(err3)
	}

	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		log.Fatalf("Server error (%s): %s", response.Status, string(body))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalf("Error decoding JSON: %v. Raw body: %s", err, string(body))
	}
	fmt.Println("URL added successfully !")

	fmt.Println(result["short_url"])
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
