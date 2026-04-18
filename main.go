package main

import (
	"fmt"
	"os"

	"url-shortner-cli/handlers"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  urlshortner shorten <url>    Shorten a URL")
		fmt.Println("  urlshortner list             List all your shortened URLs")
		fmt.Println("  urlshortner remove <code>    Remove a shortened URL")
		return
	}

	switch os.Args[1] {
	case "shorten":
		if len(os.Args) < 3 {
			fmt.Println("Usage: urlshortner shorten <url>")
			return
		}
		handlers.AddURL(os.Args[2])
	case "remove":
		if len(os.Args) < 3 {
			fmt.Println("Usage: urlshortner remove <short-code>")
			return
		}
		handlers.RemoveURL(os.Args[2])
	case "list":
		handlers.ListUrl()
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		fmt.Println("Run `urlshortner` with no arguments to see usage.")
	}
}
