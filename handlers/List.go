package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"text/tabwriter"
	"time"

	"github.com/denisbrodbeck/machineid"
)

type URLEntry struct {
	ID        string `json:"ID"`
	ShortCode string `json:"ShortCode"`
	UserID    string `json:"UserID"`
	LongURL   string `json:"LongURL"`
	CreatedAt string `json:"CreatedAt"`
}

func ListUrl() {
	userId, err := machineid.ID()
	if err != nil {
		log.Fatal("Error getting user ID:", err)
	}

	data := map[string]string{
		"UserID": userId,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	res, err := http.Post(GetServerURL()+"list", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	var urls []URLEntry
	if err := json.Unmarshal(body, &urls); err != nil {
		log.Fatalf("Error decoding JSON: %v. Raw body: %s", err, string(body))
	}

	if len(urls) == 0 {
		fmt.Println("No URLs found.")
		return
	}

	const (
		colorReset  = "\033[0m"
		colorCyan   = "\033[36m"
		colorYellow = "\033[33m"
		colorBold   = "\033[1m"
	)

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintf(w, "%s%s#\tShort Code\tLong URL\tCreated At%s\n", colorBold, colorCyan, colorReset)
	fmt.Fprintf(w, "%s%s---\t----------\t--------\t----------%s\n", colorBold, colorCyan, colorReset)
	for i, u := range urls {
		createdAt := u.CreatedAt
		if t, err := time.Parse(time.RFC3339Nano, u.CreatedAt); err == nil {
			createdAt = t.Local().Format("02 Jan 2006, 03:04 PM")
		}
		fmt.Fprintf(w, "%s%d%s\t%s%s%s\t%s\t%s\n",
			colorYellow, i+1, colorReset,
			colorBold, u.ShortCode, colorReset,
			u.LongURL,
			createdAt,
		)
	}
	w.Flush()
}
