package main

import (
	"flag"
	"fmt"
	"log"

	"GoSheetToJson/sheets"
)

func main() {
	spreadsheetId := flag.String("id", "", "Google Sheet ID")
	outputFile := flag.String("output", "output.json", "Output JSON File")
	flag.Parse()

	if *spreadsheetId == "" {
		log.Fatalf("Please provide the Google Sheet ID")
	}

	// Initialize Sheets service
	if err := sheets.InitSheetsService(); err != nil {
		log.Fatalf("Unable to initialize Sheets service: %v", err)
	}

	sheetNames, err := sheets.GetSheetNames(*spreadsheetId)
	if err != nil {
		log.Fatalf("Unable to get sheet names: %v", err)
	}

	fmt.Printf("Sheet names: %v\n", sheetNames)
	fmt.Printf("Google Sheets API initialized successfully. Output file: %s\n", *outputFile)
}
