package sheets

import (
	"context"
	"fmt"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

var srv *sheets.Service

// Initialize the Sheets service
func InitSheetsService() error {
	if srv != nil {
		return nil
	}

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

// Get all sheet names from the spreadsheet
func GetSheetNames(spreadsheetId string) ([]string, error) {
	if err := InitSheetsService(); err != nil {
		return nil, err
	}

	resp, err := srv.Spreadsheets.Get(spreadsheetId).Do()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve sheet names: %v", err)
	}

	var sheetNames []string
	for _, sheet := range resp.Sheets {
		sheetNames = append(sheetNames, sheet.Properties.Title)
	}

	return sheetNames, nil
}
