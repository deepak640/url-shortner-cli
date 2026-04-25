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
	"strconv"
	"strings"

	"github.com/denisbrodbeck/machineid"
)

func AddURL(rawURL string, customCode string, expiresIn string, MaxClick int) {
	if !IsValidHTTPSURL(rawURL) {
		fmt.Println("Invalid URL. Must start with https://")
		return
	}

	userId, err := machineid.ID()
	if err != nil {
		log.Printf("failed to get machine id: %v", err)
		return
	}

	data := map[string]any{
		"UserID": userId,
		"URL":    rawURL,
	}

	if customCode != "" {
		data["CustomCode"] = customCode
	}

	if MaxClick > 0 {
		data["MaxClicks"] = strconv.Itoa(MaxClick)
	}

	if expiresIn != "" {
		data["ExpiresIn"] = expiresIn
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("failed to marshal request payload: %v", err)
		return
	}

	response, err := http.Post(GetServerURL()+"shorten", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("failed to call backend API: %v", err)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("failed to read backend response: %v", err)
		return
	}

	if response.StatusCode == 400 {
		msg := strings.TrimSpace(string(body))
		if strings.Contains(strings.ToLower(msg), "custom") && strings.Contains(strings.ToLower(msg), "exist") {
			fmt.Println("Custom code already exists")
		} else if msg != "" {
			fmt.Printf("Request rejected by backend: %s\n", msg)
		} else {
			fmt.Println("Request rejected by backend")
		}
		return
	}
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		msg := strings.TrimSpace(string(body))
		if msg == "" {
			msg = response.Status
		}
		fmt.Printf("Backend API error (%s): %s\n", response.Status, msg)
		return
	}

	var result map[string]any
	if err := json.Unmarshal(body, &result); err != nil {
		log.Printf("error decoding backend JSON: %v. raw body: %s", err, string(body))
		return
	}

	shortURL, ok := result["short_url"]
	if !ok {
		fmt.Println("Backend response missing short_url")
		return
	}

	fmt.Println("✅ URL shortened successfully!")
	fmt.Println("🔗", shortURL)
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
