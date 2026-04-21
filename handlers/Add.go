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
	"strings"

	"github.com/denisbrodbeck/machineid"
)

func AddURL(rawURL string, customCode string, expiresIn int) {
	if !IsValidHTTPSURL(rawURL) {
		fmt.Println("Invalid URL. Must start with https://")
		return
	}

	userId, err := machineid.ID()
	if err != nil {
		log.Fatal(err)
	}

	data := map[string]interface{}{
		"UserID": userId,
		"URL":    rawURL,
	}

	if customCode != "" {
		data["CustomCode"] = customCode
	}

	if expiresIn > 0 {
		data["ExpiresIn"] = expiresIn
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	response, err := http.Post(GetServerURL()+"shorten", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)
	if response.StatusCode == 400 {
		fmt.Println("Custom code already exists")
		return
	}
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		log.Fatalf("Server error (%s): %s", response.Status, string(body))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalf("Error decoding JSON: %v. Raw body: %s", err, string(body))
	}

	fmt.Println("✅ URL shortened successfully!")
	fmt.Println("🔗", result["short_url"])
}

// generateCode generates a random 6-character alphanumeric code.
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
	if !strings.HasPrefix(strings.ToLower(input), "https://") {
		return false
	}
	parsedURL, err := url.Parse(input)
	if err != nil {
		return false
	}
	if parsedURL.Scheme != "https" {
		return false
	}
	if parsedURL.Host == "" {
		return false
	}
	return true
}
