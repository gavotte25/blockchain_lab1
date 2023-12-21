/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/gavotte25/blockchain_lab1/client"
	"github.com/gavotte25/blockchain_lab1/server"
)

func main() {
	// bc := server.InitBlockchain()
	// flag := bc.AddBlock("Send 1 BTC to Ivan")
	// if !flag {
	// 	message := fmt.Errorf("add block is not meet the constraint of the minimum %d transactions", server.MinBlockTransactions)
	// 	log.Fatal(message)
	// }
	// // bc.AddBlock("Send 2 more BTC to Ivan")
	// // bc.AddBlock("Send 10 more BTC to Ivan")

	// fmt.Printf("Number of Transactions on Chain: %d\n", bc.GetNumberOfTransactionsOnChain())
	// fmt.Printf("Number of blocks on Chain: %d\n", bc.GetNumberOfBlocks())
	// for index, block := range bc.BlockArr {
	// 	fmt.Printf("Block: %d\n", index)
	// 	block.PrintInfo()
	// 	//block.PrintTransaction()
	// }

	server.Start()
	client.Start()
}
