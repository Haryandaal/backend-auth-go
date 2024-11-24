package utils

import (
	"encoding/json"
	"os"
)

func ReadJSON(filePath string, target interface{}) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, target)
}

func WriteJSON(filePath string, target interface{}) error {
	data, err := json.MarshalIndent(target, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0644)
}
