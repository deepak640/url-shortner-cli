package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/denisbrodbeck/machineid"
)

func RemoveURL(shortCode string) {
	userId, err := machineid.ID()
	if err != nil {
		log.Fatal(err)
	}

	data := map[string]interface{}{
		"UserID": userId,
		"Code":   shortCode,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	res, err := http.Post(GetServerURL()+"remove", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalf("Error decoding JSON: %v. Raw body: %s", err, string(body))
	}

	fmt.Println(result["message"])
}
