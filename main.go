package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

var srv *sheets.Service

// Initialize the Sheets service
func initSheetsService() error {
	b, err := os.ReadFile("service_account.json") // Use your service account JSON file
	if err != nil {
		return fmt.Errorf("unable to read service account file: %v", err)
	}

	conf, err := google.JWTConfigFromJSON(b, sheets.SpreadsheetsReadonlyScope)
	if err != nil {
		return fmt.Errorf("unable to parse service account file to config: %v", err)
	}

	client := conf.Client(context.Background())
	srv, err = sheets.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		return fmt.Errorf("unable to retrieve Sheets client: %v", err)
	}

	return nil
}

func main() {
	spreadsheetId := flag.String("id", "", "Google Sheet ID")
	outputFile := flag.String("output", "output.json", "Output JSON File")
	flag.Parse()

	if *spreadsheetId == "" {
		log.Fatalf("Please provide the Google Sheet ID")
	}

	// Initialize Sheets service
	if err := initSheetsService(); err != nil {
		log.Fatalf("Unable to initialize Sheets service: %v", err)
	}

	fmt.Printf("Google Sheets API initialized successfully. Output file: %s\n", *outputFile)
}
