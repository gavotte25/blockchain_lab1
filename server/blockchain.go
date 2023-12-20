package server

import (
	"fmt"
	"time"

	"github.com/gavotte25/blockchain_lab1/utils"
)

const MinBlockTransactions = 1          // the minimum transactions of block is 3 (for easily debug)
const BlockChainDBPath = "blockchainDB" // the file name of saved blockchain's data.
const DataBasePath = "./database"

type Blockchain struct {
	BlockArr       []*Block
	TransactionIdx *TransactionIndex
}

func InitBlockchain() *Blockchain {
	// Init the first block
	Transaction := "Init transaction"
	PrevBlockHash := []byte{}
	block := NewBlock(Transaction, PrevBlockHash)
	txIndex := &TransactionIndex{Index: make(map[string]TransactionLocation)}
	stringTxIndex := utils.GetStringEncode(block.Transactions[0].Hash())
	txIndex.Index[stringTxIndex] = TransactionLocation{BlockIndex: 0, TransactionIndex: 0}
	blockchain := &Blockchain{BlockArr: []*Block{block}, TransactionIdx: txIndex}
	return blockchain
}

func (bc *Blockchain) SetTransactionIndex(blockIndex int, txIndex int) {
	if blockIndex >= 0 && txIndex >= 0 {
		tx := bc.BlockArr[blockIndex].Transactions[txIndex]
		stringTxHash := utils.GetStringEncode(tx.Hash())
		bc.TransactionIdx.Index[stringTxHash] = TransactionLocation{BlockIndex: blockIndex, TransactionIndex: txIndex}
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
	bc.SetTransactionIndex(NumberOfBlocks, 0)
	return true
}

// Add a list of transactions
func (bc *Blockchain) AddTransactions(transactionData []string) {
	// Get the last block in the blockchain
	prevBlock := bc.BlockArr[len(bc.BlockArr)-1]
	txIndex := 1
	for _, transaction := range transactionData {
		// Try to add the transaction to the current block
		// -- True: successful
		// -- False: over the limit size of block
		n := bc.GetNumberOfBlocks()
		CurrentTimestamp := time.Now().UTC().Unix()
		NewTransaction := Transaction{Data: []byte(transaction), Timestamp: CurrentTimestamp}
		flag := prevBlock.AddTransaction(NewTransaction)

		// If adding the transaction exceeds the block size limit, close the current block then create a new one
		if !flag {
			bc.AddBlock(transaction) // the init transaction with txIndex=0
			prevBlock = bc.BlockArr[len(bc.BlockArr)-1]
			txIndex = 1
		} else {
			bc.SetTransactionIndex(n-1, txIndex)
			txIndex += 1
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
	txLocation := bc.GetTransactionLocation(tx)
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
func (bc *Blockchain) GetTransactionLocation(tx *Transaction) [2]int {
	// TODO
	txHash := utils.GetStringEncode(tx.Hash())
	txIndex := bc.TransactionIdx.Index[txHash]
	return [2]int{txIndex.BlockIndex, txIndex.TransactionIndex}
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

func (bc *Blockchain) RebuildTransactionIndex(blockIndex int) {
	block := bc.GetBlock(blockIndex)
	for txIndex, _ := range block.Transactions {
		bc.SetTransactionIndex(blockIndex, txIndex)
	}
}

func (bc *Blockchain) SaveBlockChainAsJSON(truncate bool) error {
	utils.GetLog("info", "wipe old data")
	if truncate {
		utils.GetLog("info", fmt.Sprintf("-- truncate folder %s ...\n", DataBasePath))
		err := utils.WipeFolder(DataBasePath)
		if err != nil {
			return err
		} else {
			utils.GetLog("info", "-- truncate successfully.")
		}
	}
	utils.GetLog("info", "write block data")

	var orderedFiles []string
	for _, block := range bc.BlockArr {
		err := block.SaveBlockAsJSON()
		blockFileName := utils.GetStringEncode(block.Hash)
		orderedFiles = append(orderedFiles, blockFileName)
		if err != nil {
			utils.GetLog("error", fmt.Sprintf("--Failed to saved block %s \n", blockFileName))
			return err
		}
	}
	utils.GetLog("info", fmt.Sprintf("--Successful saved blockchain data at folder %s\n", DataBasePath))

	utils.GetLog("info", "write blockchain metadata")
	err := utils.WriteFile(orderedFiles, "metadata.bc", DataBasePath)
	if err != nil {
		utils.GetLog("error", fmt.Sprintf("--Failed to saved metadata %s \n", "metadata.bc"))
		return err
	} else {
		utils.GetLog("info", fmt.Sprintf("--Successful saved metadata at folder %s\n", DataBasePath+"/metadata.bc"))
	}

	return nil
}

func LoadBlockChainFromFile() *Blockchain {
	utils.GetLog("info", "load blockchain metadata ...")
	arr := utils.ReadFile("metadata.bc", DataBasePath)
	if arr == nil {
		utils.GetLog("error", fmt.Sprintf("--Failed to load metadata %s \n", DataBasePath+"/metadata.bc"))
		return nil
	}

	var bc Blockchain
	txIndex := &TransactionIndex{Index: make(map[string]TransactionLocation)}
	bc.TransactionIdx = txIndex
	for blockIndex, blockFile := range arr {
		block := loadBlockFromJSON(blockFile)
		bc.BlockArr = append(bc.BlockArr, block)
		bc.RebuildTransactionIndex(blockIndex)

	}
	return &bc
}
