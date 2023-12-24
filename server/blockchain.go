package server

import (
	"log"

	"github.com/gavotte25/blockchain_lab1/utils"
)

const BlockChainDBPath = "blockchainDB" // the file name of saved blockchain's data.

type Blockchain struct {
	BlockArr       []*Block
	TransactionIdx *TransactionIndex
}

func (bc *Blockchain) Init() {
	bc.BlockArr = make([]*Block, 0)
}

func (bc *Blockchain) AddBlock(transactions []*Transaction) *Block {
	var block *Block
	if len(bc.BlockArr) == 0 {
		block = NewBlock(transactions, nil)
	} else {
		block = NewBlock(transactions, bc.BlockArr[len(bc.BlockArr)-1])
	}
	bc.BlockArr = append(bc.BlockArr, block)
	return block
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

func (bc *Blockchain) SaveMetaDataFile(dir string) error {
	var orderedHashes []string
	for _, block := range bc.BlockArr {
		hashString := utils.GetStringEncode(block.Hash)
		orderedHashes = append(orderedHashes, hashString)
	}
	err := utils.WriteFile(orderedHashes, "metadata.bc", dir)
	if err != nil {
		log.Panicln("SaveMetaDataFile to ", dir+"metadata.bc", " failed. Reason: ", err.Error())
		return err
	}

	return nil
}
