package main

import "time"

type Blockchain struct {
	blocks []*Block
}

func InitBlockchain() *Blockchain {
	// Init the first block
	Transaction := "Init transaction"
	PrevBlockHash := []byte{}
	block := NewBlock(Transaction, PrevBlockHash)
	blockchain := &Blockchain{blocks: []*Block{block}}
	return blockchain
}

func (bc *Blockchain) AddBlock(transaction_data string) {
	PrevBlock := bc.blocks[len(bc.blocks)-1]
	block := NewBlock(transaction_data, PrevBlock.Hash)
	bc.blocks = append(bc.blocks, block)
}

// Add a list of transactions
func (bc *Blockchain) AddTransactions(transactionData []string) {
	// Get the last block in the blockchain
	prevBlock := bc.blocks[len(bc.blocks)-1]

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
			bc.blocks = append(bc.blocks, newBlock)
		}

	}
}

func (bc *Blockchain) GetNumberOfBlocks() int {
	return len(bc.blocks)
}

func (bc *Blockchain) GetNumberOfTransactionsOnChain() int {
	var result int = 0
	for _, i := range bc.blocks {
		result += i.GetNumberOfTransactionOnBlock()
	}
	return result
}
