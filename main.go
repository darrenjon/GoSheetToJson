package main

import (
	"flag"
	"fmt"
)

func main() {
	spreadsheetId := flag.String("id", "", "Google Sheet ID")
	outputFile := flag.String("output", "output.json", "Output JSON File")
	flag.Parse()

	if *spreadsheetId == "" {
		fmt.Println("Please provide the Google Sheet ID")
		return
	}

	fmt.Printf("Spreadsheet ID: %s\n", *spreadsheetId)
	fmt.Printf("Output File: %s\n", *outputFile)
}
