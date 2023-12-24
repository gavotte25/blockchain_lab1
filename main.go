package main

import (
	"bufio"
	"fmt"
	"os"

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
	// bc := server.loadBlockChainFromFile()
	// bc.BlockArr[0].PrintInfo()
	args := os.Args
	if len(args) > 1 {
		switch option := args[1]; option {
		case "server":
			server.Start()
			fmt.Println("Press enter to stop server")
			bufio.NewReader(os.Stdin).ReadString('\n')
		case "client":
			debuggingEnabled := false
			if len(args) > 2 {
				mode := args[2]
				debuggingEnabled = mode == "debug"
			}
			client.Start(debuggingEnabled)
		}
	} else {
		server.Start()
		client.Start(false)
	}
}
