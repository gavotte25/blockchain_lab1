package server

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

var (
	database_path = "./database"
)

func IsFileInDatabasePath(filename string) (bool, error) {
	filePath := filepath.Join(database_path, filename)
	_, err := os.Stat(filePath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func ReadJSONFromFile(filename string) (*Blockchain, error) {
	filePath := filepath.Join(database_path, filename)
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	jsonData, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var data Blockchain
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func WriteJSONToFile(filename string, data *Blockchain) error {
	filePath := filepath.Join(database_path, filename)
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func DeleteJSONFile(filename string) error {
	filePath := filepath.Join(database_path, filename)
	err := os.Remove(filePath)
	if err != nil {
		return err
	}
	return nil
}
