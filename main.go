package main

import (
	"github.com/gavotte25/blockchain_lab1/server"
	"github.com/gavotte25/blockchain_lab1/utils"
	// "github.com/gavotte25/blockchain_lab1/client"
	// "github.com/gavotte25/blockchain_lab1/server"
)

func main() {
	// bc := server.InitBlockchain()
	// flag := bc.AddBlock("Send 1 BTC to Ivan")
	// if !flag {
	// 	message := fmt.Errorf("add block is not meet the constraint of the minimum %d transactions", server.MinBlockTransactions)
	// 	log.Fatal(message)
	// }
	utils.GetLog("warning", "hello world")
	// TxList := [10]string{"Send 1 BTC to Ivan",
	// 	"Send 2 BTC to Ivan",
	// 	"Send 3 BTC to Ivan",
	// 	"Send 4 BTC to Ivan",
	// 	"Send 5 BTC to Ivan",
	// 	"Send 6 BTC to Ivan",
	// 	"Send 7 BTC to Ivan",
	// 	"Send 8 BTC to Ivan",
	// 	"Send 9 BTC to Ivan",
	// 	"Send 10 BTC to Ivan",
	// }
	// bc.AddTransactions(TxList[:])

	// bc.SaveBlockChainAsJSON(true)
	bc := server.LoadBlockChainFromFile()
	bc.BlockArr[0].PrintInfo()
	// server.Start()
	// client.Start()
}
