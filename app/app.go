package app

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gavotte25/blockchain_lab1/client"
)

func clearConsole() {
	fmt.Print("\033[H\033[2J")
}

func Run() bool {
	fmt.Println("Welcome to the application!")
	fmt.Println("Please select an option:")
	fmt.Println("1. Login")
	fmt.Println("2. Exit")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		fmt.Println("Login selected")
		wallet := client.NewWallet()
		for {
			clearConsole()
			state := Client(wallet)
			if !state {
				break
			}
		}
	case 2:
		fmt.Println("Exit selected")
		return false
	default:
		fmt.Println("Invalid choice")
		return false
	}
	return true
}

func Client(w *client.Wallet) bool {
	fmt.Println("Welcome to the client application!")
	fmt.Println("Please select an option:")
	fmt.Println("1. Make transaction")
	fmt.Println("2. Get transaction")
	fmt.Println("3. Verify transaction")
	fmt.Println("4. Print block information")
	fmt.Println("5. Exit")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		fmt.Println("Make transaction selected")
		fmt.Printf("Type the content of the transaction here: ")
		reader := bufio.NewReader(os.Stdin)
		info, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err.Error())
		} else {
			isSuccess := w.MakeTransaction(info)
			if isSuccess {
				fmt.Println("Transaction is made successfully. Press Enter to continue")
				fmt.Scanln()
			} else {
				fmt.Println("Transaction is made unsuccessfully. Press Enter to continue")
				fmt.Scanln()
			}
		}

	case 2:
		fmt.Println("Get transaction selected")
		// Add your get transaction logic here
	case 3:
		fmt.Println("Verify transaction selected")
		// Add your verify transaction logic here
	case 4:
		fmt.Println("Print block information selected")
		// Add your print block information logic here
	case 5:
		fmt.Println("Exit selected")
		w.Finish()
		return false
	default:
		fmt.Println("Invalid choice")
	}

	return true
}
