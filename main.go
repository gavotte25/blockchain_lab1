package main

import (
	"fmt"
)

func main() {
	bc := InitBlockchain()

	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")
	bc.AddBlock("Send 10 more BTC to Ivan")

	fmt.Printf("Number of Transactions on Chain: %d\n", bc.GetNumberOfTransactionsOnChain())
	fmt.Printf("Number of blocks on Chain: %d\n", bc.GetNumberOfBlocks())
	for index, block := range bc.blocks {
		fmt.Printf("Block: %d\n", index)
		block.PrintInfo()
		//block.PrintTransaction()
	}
}
