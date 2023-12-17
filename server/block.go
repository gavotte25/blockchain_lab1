package server

import (
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/gavotte25/blockchain_lab1/utils"
)

const maxBlockSize = 128 * 1024 // the maximum size of block is 128 KB

type Block struct {
	Timestamp     int64
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
}

func (b *Block) SetHash() {
	timeBytes := utils.ConvertTimestampToByte(b.Timestamp)
	hashInput := append(b.HashTransactions(), timeBytes...)
	hashOutput := sha256.Sum256(hashInput)
	b.Hash = hashOutput[:]
}

func (b *Block) HashTransactions() []byte {
	var hashInput []byte
	for _, tx := range b.Transactions {
		hashTx := sha256.Sum256(tx.Data)
		hashInput = append(hashInput, hashTx[:]...)
	}
	hashOutput := sha256.Sum256(hashInput)
	return hashOutput[:]
}

func NewBlock(TransactionData string, prevBlockHash []byte) *Block {
	hash := []byte{} // Assuming this is an empty byte slice; you need to calculate the hash
	CurrentTimestamp := time.Now().UTC().Unix()
	TransactionByte := []byte(TransactionData)
	block := &Block{
		Timestamp: CurrentTimestamp,
		Transactions: []*Transaction{{Data: TransactionByte,
			Timestamp: CurrentTimestamp},
		},
		PrevBlockHash: prevBlockHash,
		Hash:          hash,
	}

	block.SetHash()
	return block
}

func (block *Block) GetNumberOfTransactionOnBlock() int {
	return len(block.Transactions)
}

func (b *Block) AddTransaction(NewTransaction Transaction) bool {
	//TransactionData := []byte(TransactionRaw)
	// Check if adding the new transaction exceeds the block size limit
	currentBlockSize := len(b.HashTransactions())
	newTransactionSize := len(NewTransaction.Data)

	if currentBlockSize+newTransactionSize > maxBlockSize {
		return false
	}

	b.Transactions = append(b.Transactions, &NewTransaction)
	b.SetHash()
	return true
}

func (block *Block) PrintInfo() {
	fmt.Printf("Block address: %x\n", block.Hash)
	fmt.Printf("Block size: %d bytes\n", len(block.HashTransactions()))
	fmt.Printf("Created timestamp (UTC+0): %s\n", utils.GetTimestampFormat(block.Timestamp))
	fmt.Printf("Number of Transactions: %d\n", len(block.Transactions))
}

func (block *Block) PrintTransaction() {
	for idx, transaction := range block.Transactions {
		fmt.Printf("-- Transaction number: %d\n", idx)
		fmt.Printf("-- Created timestamp (UTC+0): %s\n", utils.GetTimestampFormat(transaction.Timestamp))
		fmt.Printf("-- Data: %x\n", transaction.Data)
	}
}
