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

func (bc *Blockchain) VerifyTransaction(tx *Transaction, merklePath [][]byte) bool {
	txLocation := bc.GetTransactionLocation(1)
	blockIndex := txLocation[0]
	txIndex := txLocation[1]
	if blockIndex < 0 || txIndex < 0 {
		return false
	}
	block := bc.GetBlock(blockIndex)
	if block == nil {
		return false
	}
	return block.VerifyTransaction(tx, merklePath)
}

// GetTransactionLocation returns array of block index and transaction index in that block. Return {-1, -1} if can't find
func (bc *Blockchain) GetTransactionLocation(transactionId int) [2]int {
	// TODO
	return [2]int{-1, -1}
}

func (bc *Blockchain) GetBlock(index int) *Block {
	if index < 0 || index > len(bc.BlockArr)-1 {
		return nil
	} else {
		return bc.BlockArr[index]
	}
}

func (bc *Blockchain) getLightVersion() *Blockchain {
	lightBc := new(Blockchain)
	lightBc.BlockArr = make([]*Block, len(bc.BlockArr))
	for i := 0; i < len(bc.BlockArr); i++ {
		lightBc.BlockArr[i] = bc.BlockArr[i].GetBlockHeader()
	}
	return lightBc
}

func (bc *Blockchain) GetVersionNumber() int {
	return len(bc.BlockArr)
}

func (bc *Blockchain) Append(blocks []*Block) {
	bc.BlockArr = append(bc.BlockArr, blocks...)
}
