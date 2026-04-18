package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/denisbrodbeck/machineid"
	"github.com/joho/godotenv"
)

func RemoveURL(shortCode string){
	if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using system environment variables")
  }
	Server := os.Getenv("SERVER")
	UserID, err := machineid.ID()
	if err != nil{
		log.Fatal(err)
		return
	}

	data := map[string]interface{}{
		"UserID":UserID,
		"Code":shortCode,
	}

	jsonData,err := json.Marshal(data)

	if err != nil{
		log.Fatal(err)
		return
	}

	res,err := http.Post(Server + "remove","application/json",bytes.NewBuffer(jsonData))
	if err != nil{
		log.Fatal(err)
		return
	}
	defer res.Body.Close()

	body,err := io.ReadAll(res.Body)
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalf("Error decoding JSON: %v. Raw body: %s", err, string(body))
	}
	fmt.Println(result["message"])
}
