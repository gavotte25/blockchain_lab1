package server

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/gavotte25/blockchain_lab1/utils"
)

type Block struct {
	Timestamp     int64
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
}

func (block *Block) SetHash() {
	timeBytes := utils.ConvertTimestampToByte(block.Timestamp)
	var hashInput []byte
	if block.PrevBlockHash != nil {
		hashInput = append(block.PrevBlockHash, block.HashTransactions()...)
	} else {
		hashInput = block.HashTransactions()
	}
	hashInput = append(hashInput, timeBytes...)
	hashOutput := sha256.Sum256(hashInput)
	block.Hash = hashOutput[:]
}

func (block *Block) HashTransactions() []byte {
	return CreateMerkleTree(block.Transactions).Root.Hash
}

func NewBlock(transactions []*Transaction, prevBlock *Block) *Block {
	CurrentTimestamp := time.Now().UTC().Unix()
	block := &Block{
		Timestamp:    CurrentTimestamp,
		Transactions: transactions,
	}
	if prevBlock != nil {
		block.PrevBlockHash = prevBlock.Hash
	}
	block.SetHash()
	return block
}

func (block *Block) GetNumberOfTransactionOnBlock() int {
	return len(block.Transactions)
}

func (block *Block) GetTransaction(index int) *Transaction {
	if index < 0 || index > len(block.Transactions)-1 {
		return nil
	} else {
		return block.Transactions[index]
	}
}

func (b *Block) AddTransaction(NewTransaction Transaction) bool {
	b.Transactions = append(b.Transactions, &NewTransaction)
	b.SetHash()
	return true
}

func (block *Block) PrintInfo() {
	fmt.Printf("Block address: %x\n", block.Hash)
	fmt.Printf("Block size: %d bytes\n", block.GetBlockSize())
	fmt.Printf("Created timestamp (UTC+0): %s\n", utils.GetTimestampFormat(block.Timestamp))
	fmt.Printf("Number of Transactions: %d\n", len(block.Transactions))
	// utils.GetLog("INFO", fmt.Sprintf("Block address: %x", block.Hash))
	// utils.GetLog("INFO", fmt.Sprintf("Block size: %d bytes", block.GetBlockSize()))
	// utils.GetLog("INFO", fmt.Sprintf("Created timestamp (UTC+0): %s", utils.GetTimestampFormat(block.Timestamp)))
	// utils.GetLog("INFO", fmt.Sprintf("Number of Transactions: %d", len(block.Transactions)))
}

func (block *Block) PrintTransaction() {
	for idx, transaction := range block.Transactions {
		utils.GetLog("info", fmt.Sprintf("-- Transaction index: %d", idx))
		utils.GetLog("info", fmt.Sprintf("-- Created timestamp (UTC+0): %s",
			utils.GetTimestampFormat(transaction.Timestamp)))
		utils.GetLog("info", fmt.Sprintf("-- Data: %s", string(transaction.Data)))
	}
}

func (block *Block) VerifyTransaction(tx *Transaction, merklePath *MerklePath) bool {
	if tx == nil || merklePath == nil || merklePath.Steps == nil || merklePath.Direction == nil || len(merklePath.Steps) == 0 {
		return false
	}
	var hash [32]byte
	if merklePath.Direction[0] {
		hash = sha256.Sum256(append(tx.Hash(), merklePath.Steps[0]...))
	} else {
		hash = sha256.Sum256(append(merklePath.Steps[0], tx.Hash()...))
	}
	for i := 1; i < len(merklePath.Steps); i++ {
		if merklePath.Direction[i] {
			hash = sha256.Sum256(append(hash[:], merklePath.Steps[i]...))
		} else {
			hash = sha256.Sum256(append(merklePath.Steps[i], hash[:]...))
		}
	}
	timeBytes := utils.ConvertTimestampToByte(block.Timestamp)
	var hashInput []byte
	if block.PrevBlockHash != nil {
		hashInput = append(block.PrevBlockHash, hash[:]...)
	} else {
		hashInput = hash[:]
	}
	hashInput = append(hashInput, timeBytes...)
	hashOutput := sha256.Sum256(hashInput)
	return hashOutput == [32]byte(block.Hash)
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

func (block *Block) SaveBlockAsJSON(dir string) error {
	fileName := utils.GetStringEncode(block.Hash)
	filePath := filepath.Join(dir, fileName+".json")
	jsonData, err := json.Marshal(block)
	if err != nil {
		return errors.New("Can't marshall file " + filePath + " reason: " + err.Error())
	}
	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return errors.New("Can't read file " + filePath + " reason: " + err.Error())
	}
	return nil
}

func LoadBlockFromJSON(fileName string, dir string) (*Block, error) {
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

	var block Block
	err = json.Unmarshal(jsonData, &block)
	if err != nil {
		return nil, errors.New("Can't unmarshal file " + filePath + " reason: " + err.Error())
	}

	return &block, nil
}
