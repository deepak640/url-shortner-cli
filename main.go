package main

import (
	"fmt"
	"os"
	"flag"

	"ziplink/handlers"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  ziplink shorten --url <url> [--custom <code>] [--expiry <hours>]")
		fmt.Println("  ziplink list")
		fmt.Println("  ziplink remove <code>")
		return
	}

	switch os.Args[1] {
	case "shorten":
		shortenCmd := flag.NewFlagSet("shorten", flag.ExitOnError)
		urlFlag := shortenCmd.String("url", "", "The URL to shorten (required)")
		customFlag := shortenCmd.String("custom", "", "Custom short code (optional)")
		expiryFlag := shortenCmd.Int("expiry", 0, "Expiry in hours (optional)")
		MaxClickFlag := shortenCmd.Int("clicks", 0, "Maximum clicks (optional)")


		shortenCmd.Parse(os.Args[2:])

		if *urlFlag == "" {
			fmt.Println("Error: --url flag is required")
			fmt.Println("Usage: ziplink shorten --url <url> [--custom <code>] [--expiry <hours>]")
			return
		}
		handlers.AddURL(*urlFlag, *customFlag, *expiryFlag, *MaxClickFlag)
	case "remove":
		if len(os.Args) < 3 {
			fmt.Println("Usage: ziplink remove <short-code>")
			return
		}
		handlers.RemoveURL(os.Args[2])
	case "list":
		handlers.ListUrl()
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		fmt.Println("Run `ziplink` with no arguments to see usage.")
	}
}
