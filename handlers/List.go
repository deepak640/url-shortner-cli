package handlers

import (
	"context"
	"fmt"
	"log"
	"os"
	"url-shortner-cli/config"
	"url-shortner-cli/models"

	"github.com/denisbrodbeck/machineid"
	"github.com/olekukonko/tablewriter"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func ListUrl() {
	userId, Iderr := machineid.ID()
	if Iderr != nil {
		log.Fatal(Iderr)
	}

	cursor, err := config.DB.Database("urlshortener").Collection("urls").Find(context.TODO(), bson.D{{Key: "user_id", Value: userId}})

	if err != nil {
		log.Fatal("Error listing URLs:", err)
	}
	defer cursor.Close(context.TODO())

	var urls []models.URL
	if err = cursor.All(context.TODO(), &urls); err != nil {
		log.Fatal("Error decoding URLs:", err)
	}

	if len(urls) == 0 {
		fmt.Println("No URLs found.")
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Short Code", "Short Url", "Long Url", "Created At"})
	table.SetBorder(true)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	for _, url := range urls {
		table.Append([]string{
			url.ShortCode,
			"http://localhost:8080/" + url.ShortCode,
			url.LongURL,
			url.CreatedAt.Format("2006-01-02"),
		})
	}
	table.Render()
}

