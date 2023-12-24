package cli

import (
	"bufio"
	"flag"
	"log"
	"os"

	"github.com/gavotte25/blockchain_lab1/client"
	"github.com/gavotte25/blockchain_lab1/server"
)

type CLI struct {
	bc *server.Blockchain
}

func (cli *CLI) Run() {
	cli.validateArgs()

	loginCmd := flag.NewFlagSet("login", flag.ExitOnError)
	makeTransactionCmd := flag.NewFlagSet("make", flag.ExitOnError)
	finishCmd := flag.NewFlagSet("finish", flag.ExitOnError)

	transactionContent := makeTransactionCmd.String("content", "", "Transaction content")

	if os.Args[1] == "login" {
		err := loginCmd.Parse(os.Args[2:])
		if err == nil && loginCmd.Parsed() {
			cli.login()
		} else {
			log.Println("An error occurs when logging in")
		}
	} else {
		log.Println("You must login first")
	}

	// switch os.Args[1] {
	// case "login":
	// 	err := loginCmd.Parse(os.Args[2:])
	// case "make":
	// 	err := makeTransactionCmd.Parse(os.Args[2:])
	// case "finish":
	// 	err := finishCmd.Parse(os.Args[2:])
	// default:
	// 	cli.printUsage()
	// 	os.Exit(1)
	// }

	// if makeTransactionCmd.Parsed() {
	// 	if *transactionContent == "" {
	// 		makeTransactionCmd.Usage()
	// 		os.Exit(1)
	// 	}
	// 	cli.makeTransaction(*transactionContent)
	// }

	// if loginCmd.Parsed() {
	// 	cli.login()
	// }

	// if finishCmd.Parsed() {
	// 	cli.finish()
	// }
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) printUsage() {
	usage := "Usage:\n" +
		"  login - Login to the system\n" +
		"  make - Make a transaction\n" +
		"  finish - Finish the work\n"
	os.Stderr.WriteString(usage)
}

func (cli *CLI) login() {
	log.Println("Client started")
	reader := bufio.NewReader(os.Stdin)
	wallet := new(client.Wallet)
	wallet.init(true)

}
