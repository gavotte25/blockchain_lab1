package server

import (
	"fmt"
	"time"
)

const MinBlockTransactions = 1 // the minimum transactions of block is 3 (for easily debug)

type Blockchain struct {
	BlockArr []*Block
}

func InitBlockchain() *Blockchain {
	// Init the first block
	Transaction := "Init transaction"
	PrevBlockHash := []byte{}
	block := NewBlock(Transaction, PrevBlockHash)
	blockchain := &Blockchain{BlockArr: []*Block{block}}
	return blockchain
}

func CreateBlockchain(chainName string) {
	// Check if chainName is exists
	chainNameJSON := chainName + ".json"
	isExist, _ := IsFileInDatabasePath(chainNameJSON)
	if isExist {
		fmt.Println("Attempting to create a blockchain failed. Either the name of blockchain already exists or an error occurs.")
	} else {
		bc := InitBlockchain()
		err := WriteJSONToFile(chainNameJSON, bc)
		if err == nil {
			fmt.Printf("Create a blockchain with name %s successfully.\n", chainName)
		} else {
			fmt.Println("An error occurs when create a new blockchain.")
		}
	}
}

func DeleteBlockchain(chainName string) {
	chainNameJSON := chainName + ".json"
	isExist, _ := IsFileInDatabasePath(chainNameJSON)
	if !isExist {
		fmt.Println("Attempting to delete a blockchain failed. Either the blockchain does not exist or an error occurs.")
	} else {
		err := DeleteJSONFile(chainNameJSON)
		if err == nil {
			fmt.Printf("Delete a blockchain with name %s successfully.\n", chainName)
		} else {
			fmt.Printf("An error occurs when delete a blockchain with name %s.\n", chainName)
		}
	}
}

func (bc *Blockchain) AddBlock(transaction_data string) bool {
	NumberOfBlocks := len(bc.BlockArr)
	PrevBlock := bc.BlockArr[NumberOfBlocks-1]
	if NumberOfBlocks >= 1 && PrevBlock.GetNumberOfTransactionOnBlock() < MinBlockTransactions {
		return false
	}
	block := NewBlock(transaction_data, PrevBlock.Hash)
	bc.BlockArr = append(bc.BlockArr, block)
	return true
}

// Add a list of transactions
func (bc *Blockchain) AddTransactions(transactionData []string) {
	// Get the last block in the blockchain
	prevBlock := bc.BlockArr[len(bc.BlockArr)-1]

	for _, transaction := range transactionData {
		// Try to add the transaction to the current block
		// -- True: successful
		// -- False: over the limit size of block
		CurrentTimestamp := time.Now().UTC().Unix()
		NewTransaction := Transaction{Data: []byte(transaction), Timestamp: CurrentTimestamp}
		flag := prevBlock.AddTransaction(NewTransaction)

		// If adding the transaction exceeds the block size limit, close the current block then create a new one
		if !flag {
			newBlock := NewBlock(transaction, prevBlock.Hash)
			bc.BlockArr = append(bc.BlockArr, newBlock)
		}
	}
}

func (bc *Blockchain) GetNumberOfBlocks() int {
	return len(bc.BlockArr)
}

func (bc *Blockchain) GetNumberOfTransactionsOnChain() int {
	var result int = 0
	for _, i := range bc.BlockArr {
		result += i.GetNumberOfTransactionOnBlock()
	}
	return result
}
