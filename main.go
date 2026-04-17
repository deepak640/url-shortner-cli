package main

import (
	"fmt"
	"os"
	"url-shortner-cli/config"
	"url-shortner-cli/handlers"

)


func main(){
	config.Connect()
	if len(os.Args) < 2 {
		fmt.Println("Please provide a URL")
		return
	}
	switch os.Args[1] {
	case "shorten":
		handlers.AddURL(os.Args[2])
	case "remove":
		handlers.RemoveURL(os.Args[2])
	case "list":
		handlers.ListUrl()
	default:
		fmt.Println("Invalid command")
	}
}
