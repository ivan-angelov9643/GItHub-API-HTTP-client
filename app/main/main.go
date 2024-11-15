package main

import (
	"fmi-go-homework-1/app"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <filename>", os.Args[0])
	}
	filename := os.Args[1]

	usernames, err := app.ReadUsernames(filename)
	if err != nil {
		log.Fatalf("Failed to read usernames: %v", err)
	}

	app.GenerateReport(usernames)
}
