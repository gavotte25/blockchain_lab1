package server

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/gavotte25/blockchain_lab1/utils"
	//"github.com/gavotte25/blockchain_lab1/database"
)

const maxBlockSize = 128 //* 1024 // the maximum size of block is 128 B = 2 transactions on each block

type Block struct {
	Timestamp     int64
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
}

func (block *Block) SetHash() {
	timeBytes := utils.ConvertTimestampToByte(block.Timestamp)
	hashInput := append(block.PrevBlockHash, block.HashTransactions()...)
	hashInput = append(hashInput, timeBytes...)
	hashOutput := sha256.Sum256(hashInput)
	block.Hash = hashOutput[:]
}

func (block *Block) HashTransactions() []byte {
	return CreateMerkleTree(block.Transactions).Root.Hash
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
	currentBlockSize := b.GetBlockSize()
	newTransactionSize := len(NewTransaction.Data)

	if currentBlockSize+newTransactionSize > maxBlockSize {
		return false
	}

	b.Transactions = append(b.Transactions, &NewTransaction)
	b.SetHash()
	return true
}

func (block *Block) PrintInfo() {
	utils.GetLog("info", fmt.Sprintf("Block address: %x", block.Hash))
	utils.GetLog("info", fmt.Sprintf("Block size: %d bytes", block.GetBlockSize()))
	utils.GetLog("info", fmt.Sprintf("Created timestamp (UTC+0): %s", utils.GetTimestampFormat(block.Timestamp)))
	utils.GetLog("info", fmt.Sprintf("Number of Transactions: %d", len(block.Transactions)))
}

func (block *Block) PrintTransaction() {
	for idx, transaction := range block.Transactions {
		utils.GetLog("info", fmt.Sprintf("-- Transaction index: %d", idx))
		utils.GetLog("info", fmt.Sprintf("-- Created timestamp (UTC+0): %s",
			utils.GetTimestampFormat(transaction.Timestamp)))
		utils.GetLog("info", fmt.Sprintf("-- Data: %x", transaction.Data))
	}
}

func (block *Block) VerifyTransaction(tx *Transaction, merklePath [][]byte) bool {
	if tx == nil || merklePath == nil || len(merklePath) == 0 {
		return false
	}
	hash := sha256.Sum256(append(tx.Hash(), merklePath[0]...))
	for i := 1; i < len(merklePath); i++ {
		hash = sha256.Sum256(append(hash[:], merklePath[i]...))
	}
	return hash == [32]byte(block.Hash)
}

func (block *Block) GetHash() [32]byte {
	return [32]byte(block.Hash)
}

func (block *Block) GetBlockHeader() *Block {
	header := new(Block)
	header.Hash = block.Hash
	header.PrevBlockHash = block.PrevBlockHash
	header.Timestamp = block.Timestamp
	return header
}

func (b *Block) GetBlockSize() int {
	blockSize := binary.Size(b.Timestamp)
	blockSize += len(b.PrevBlockHash)
	blockSize += len(b.Hash)
	for _, tx := range b.Transactions {
		blockSize += len(tx.Data)
		blockSize += binary.Size(tx.Timestamp)
	}
	return blockSize
}

func (block *Block) SaveBlockAsJSON() error {
	fileName := utils.GetStringEncode(block.Hash)
	filePath := filepath.Join(DataBasePath, fileName+".json")
	jsonData, err := json.Marshal(block)
	if err != nil {
		return err
	}
	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}

func loadBlockFromJSON(fileName string) *Block {
	filePath := filepath.Join(DataBasePath, fileName+".json")
	file, err := os.Open(filePath)
	if err != nil {
		return nil
	}
	defer file.Close()

	jsonData, err := io.ReadAll(file)
	if err != nil {
		return nil
	}

	var block Block
	err = json.Unmarshal(jsonData, &block)
	if err != nil {
		return nil
	}

	return &block
}
