package main

import (
	"flag"
	"fmt"
	"os"

	"ziplink/handlers"
)

func printUsage() {
	fmt.Println("Ziplink - A fast and simple URL shortener CLI")
	fmt.Println("\nUsage:")
	fmt.Println("  ziplink <command> [arguments]")
	fmt.Println("\nCommands:")
	fmt.Println("  shorten    Create a new short URL (Requires --url)")
	fmt.Println("  list       List all URLs created from this machine")
	fmt.Println("  remove     Delete a shortened URL by its short code")
	fmt.Println("\nFlags:")
	fmt.Println("  -h, --help Show this help message")
	fmt.Println("\nExamples:")
	fmt.Println("  ziplink shorten --url https://github.com")
	fmt.Println("  ziplink shorten --url https://google.com --custom mygoogle --expiry 7d")
	fmt.Println("\nUse \"ziplink <command> --help\" for detailed information on flags.")
}

func printShortenUsage() {
	fmt.Println("Usage: ziplink shorten --url <url> [flags]")
	fmt.Println("\nDescription:")
	fmt.Println("  Generates a short link for a given HTTPS URL. You can optionally set")
	fmt.Println("  a custom code, a click limit, or an expiration time.")
	fmt.Println("\nFlags:")
	fmt.Println("  --url <url>       (Required) The destination URL. Must start with https://")
	fmt.Println("  --custom <code>   (Optional) Use a specific string for the short link instead of a random one.")
	fmt.Println("  --clicks <num>    (Optional) The link will stop working after this many clicks.")
	fmt.Println("  --expiry <time>   (Optional) How long the link stays active.")
	fmt.Println("                    Format: <number>[unit]. If no unit is given, hours is used.")
	fmt.Println("\n                    Expiry Units Logic:")
	fmt.Println("                      h - Hours:  1 unit = 1 hour (Default). Example: 24h")
	fmt.Println("                      d - Days:   1 unit = 24 hours. Example: 7d")
	fmt.Println("                      m - Months: 1 unit = 30 days (720 hours). Example: 1m")
	fmt.Println("\n                    Example: --expiry 48h, --expiry 30d, --expiry 2")
	fmt.Println("  -h, --help        Show this help message")
}

func main() {
	if len(os.Args) < 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		printUsage()
		return
	}

	switch os.Args[1] {
	case "shorten":
		shortenCmd := flag.NewFlagSet("shorten", flag.ExitOnError)
		urlFlag := shortenCmd.String("url", "", "")
		customFlag := shortenCmd.String("custom", "", "")
		expiryFlag := shortenCmd.String("expiry", "", "")
		MaxClickFlag := shortenCmd.Int("clicks", 0, "")

		shortenCmd.Usage = printShortenUsage

		if len(os.Args) == 2 {
			printShortenUsage()
			return
		}

		shortenCmd.Parse(os.Args[2:])

		if *urlFlag == "" {
			fmt.Println("❌ Error: The --url flag is required.")
			printShortenUsage()
			return
		}
		handlers.AddURL(*urlFlag, *customFlag, *expiryFlag, *MaxClickFlag)

	case "remove":
		if len(os.Args) < 3 || os.Args[2] == "-h" || os.Args[2] == "--help" {
			fmt.Println("Usage: ziplink remove <code>")
			fmt.Println("\nDescription:")
			fmt.Println("  Removes a shortened URL mapping from the server.")
			fmt.Println("\nExample:")
			fmt.Println("  ziplink remove a1b2c3")
			return
		}
		handlers.RemoveURL(os.Args[2])

	case "list":
		if len(os.Args) > 2 && (os.Args[2] == "-h" || os.Args[2] == "--help") {
			fmt.Println("Usage: ziplink list")
			fmt.Println("\nDescription:")
			fmt.Println("  Shows all active short URLs associated with your Machine ID.")
			return
		}
		handlers.ListUrl()

	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		printUsage()
	}
}
