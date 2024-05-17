package main

import (
	"fmt"
	"log"
	"os"

	"GoSheetToJson/sheets"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	spreadsheetId := os.Getenv("SPREADSHEET_ID")
	if spreadsheetId == "" {
		log.Fatalf("SPREADSHEET_ID not set in .env file")
	}

	outputFile := "output.json"

	// Initialize Sheets service
	if err := sheets.InitSheetsService(); err != nil {
		log.Fatalf("Unable to initialize Sheets service: %v", err)
	}

	sheetNames, err := sheets.GetSheetNames(spreadsheetId)
	if err != nil {
		log.Fatalf("Unable to get sheet names: %v", err)
	}

	fmt.Printf("Sheet names: %v\n", sheetNames)
	fmt.Printf("Google Sheets API initialized successfully. Output file: %s\n", outputFile)

	data := make(map[string]interface{})
	for _, sheetName := range sheetNames {
		sheetData, err := sheets.ReadSheet(spreadsheetId, sheetName)
		if err != nil {
			log.Fatalf("Unable to read Google Sheet data: %v", err)
		}
		data[sheetName] = sheetData
	}

	if err := sheets.SaveToJson(outputFile, data); err != nil {
		log.Fatalf("Unable to save JSON file: %v", err)
	}

	fmt.Printf("Successfully saved data to %s\n", outputFile)
}
