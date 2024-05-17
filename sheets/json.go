package sheets

import (
	"encoding/json"
	"fmt"
	"os"
)

// Convert data to JSON format and save to a file
func SaveToJson(filename string, data interface{}) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("unable to create file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("unable to encode data to JSON: %v", err)
	}

	return nil
}
