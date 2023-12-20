package server

import (
	"crypto/sha256"

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
