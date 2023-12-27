package server

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/gavotte25/blockchain_lab1/utils"
)

type Transaction struct {
	Data      []byte
	Timestamp int64
}

func (tx *Transaction) Hash() []byte {
	timeBytes := utils.ConvertTimestampToByte(tx.Timestamp)
	hashInput := append(tx.Data, timeBytes...)
	hashOutput := sha256.Sum256(hashInput)
	return hashOutput[:]
}

type TransactionIndex struct {
	Index map[string]TransactionLocation
}

type TransactionLocation struct {
	BlockIndex       int
	TransactionIndex int
}

func (tx *Transaction) SaveTransactionAsJson(file_name string, dir string) error {
	fileName := file_name
	if fileName == "" {
		fileName = utils.GetStringEncode(tx.Hash())
	}

	filePath := filepath.Join(dir, fileName+".json")
	jsonData, err := json.Marshal(tx)
	if err != nil {
		return errors.New("Can't marshall file " + filePath + " reason: " + err.Error())
	}
	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return errors.New("Can't read file " + filePath + " reason: " + err.Error())
	}
	return nil
}

func LoadTransactionFromJSON(fileName string, dir string) (*Transaction, error) {
	filePath := filepath.Join(dir, fileName+".json")
	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New("Can't open file " + filePath + " reason: " + err.Error())
	}
	defer file.Close()

	jsonData, err := io.ReadAll(file)
	if err != nil {
		return nil, errors.New("Can't read file " + filePath + " reason: " + err.Error())
	}

	var tx Transaction
	err = json.Unmarshal(jsonData, &tx)
	if err != nil {
		return nil, errors.New("Can't unmarshal file " + filePath + " reason: " + err.Error())
	}

	return &tx, nil
}
