package app

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/gavotte25/blockchain_lab1/client"
	"github.com/gavotte25/blockchain_lab1/server"
)

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func clearConsole() {
	CallClear()
}

func Run() bool {
	clearConsole()
	fmt.Println("Welcome to the blockchain application!")
	fmt.Println("1. Sign up")
	fmt.Println("0. Exit")
	fmt.Println("Please select an option:")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		//fmt.Println("Login selected")
		wallet := client.NewWallet()
		for {
			clearConsole()
			message, state := Client(wallet)
			if !state {
				break
			}
			fmt.Println(message)
			fmt.Scanln()
		}
	case 0:
		fmt.Println("Your session is over. See you later.")
		return false
	default:
		fmt.Println("Invalid choice.")
		return false
	}
	return true
}

func drawTableHeader(option2 bool) {
	if !option2 {
		fmt.Printf("| %-4s | %-10s | %-20s |\n", "No.", "Id", "Content")
		fmt.Println("---------------------------------------------------------------------------")
	} else {
		fmt.Printf("| %-4s | %-10s | %-20s | %-10s |\n", "No.", "Id", "Content", "Status")
		fmt.Println("---------------------------------------------------------------------------")
	}
}

func drawOnePair(no int, tx *server.Transaction, option2 bool, verified_status string) {
	if !option2 {
		fmt.Printf("| %-4d | %-10d | %-20s |\n", no, tx.Timestamp, strings.TrimSpace(string(tx.Data[:])))
	} else {
		fmt.Printf("| %-4d | %-10d | %-20s | %-10s |\n", no, tx.Timestamp, strings.TrimSpace(string(tx.Data[:])), verified_status)
	}
}

func drawTable(transactions []*server.Transaction, unverifiedTransaction []*server.Transaction, option2 bool) {
	if option2 {
		drawTableHeader(true)
		for id, content := range transactions {
			status := "Verified"
			for _, transaction := range unverifiedTransaction {
				if content.Timestamp == transaction.Timestamp {
					status = "UNVERIFIED"
				}
			}
			drawOnePair(id, content, true, status)
		}

	} else {
		drawTableHeader(false)
		for id, content := range transactions {
			drawOnePair(id, content, false, "")
		}
	}

}

func Client(w *client.Wallet) (string, bool) {
	fmt.Println("==== LOGGIN' IN ====")
	fmt.Println("1. Make transaction")
	fmt.Println("2. Verify transaction")
	fmt.Println("3. View block information")
	fmt.Println("0. Exit")
	fmt.Printf("Please select an option: ")
	var message string
	var choice int
	//var choice int
	fmt.Scanln(&choice)
	//fmt.Println(err)
	//fmt.Printf("Your choice is %d\n", choice)
	switch choice {
	case 1:
		//fmt.Println("Make transaction selected")
		fmt.Printf("Type the content of the transaction: ")

		reader := bufio.NewReader(os.Stdin)
		info, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err.Error())
		} else {
			isSuccess := w.MakeTransaction(strings.TrimSpace(info))
			if isSuccess {
				//fmt.Println("Success! Your transaction is ready to verify. Press Enter to continue")
				message = "Success! Your transaction is ready to verify. Press Enter to continue"
			} else {
				message = "Your transaction is failed to make. Please try again. Press Enter to continue"
			}
		}
	case 3:
		// read the blockname
		filePath := "server/database/metadata.bc"
		data, err := os.ReadFile(filePath)
		if err != nil {
			message = "Your blockchain has no block. Make more transactions and try again. Press Enter to continue."
		}
		blockNames := strings.Split(string(data), "\n")
		dir := "server/database"
		// list the blockname
		clearConsole()
		fmt.Printf("| %-4s | %-64s |\n", "No.", "Block address")
		fmt.Println("---------------------------------------------------------------------------")
		for id, blockName := range blockNames {
			if blockName != "" {
				fmt.Printf("| %-4d | %-64s |\n", id, blockName)
			}
		}

		fmt.Print("Select the No. of block: ")
		var selection int
		fmt.Scan(&selection)
		if selection >= len(blockNames)-1 {
			message = "Your selection is out of range. Press Enter to try again."
			return message, true
		}

		// read the file of block
		block, _ := server.LoadBlockFromJSON(blockNames[selection], dir)
		transactions := block.Transactions
		unverifiedTransaction, error := w.ReadTransactionFile()
		//fmt.Println(transactions)
		if error != nil {
			unverifiedTransaction = nil
		}
		clearConsole()
		w.PrintBlock(selection)
		drawTable(transactions, unverifiedTransaction, true)
		message = "Press Enter to continue."
		return message, true

	case 2:
		// readfile
		transactions, error := w.ReadTransactionFile()
		if error != nil {
			message = "You have no transaction to verify. Press Enter to continue."
			return message, true
		}
		// draw list of transaction
		drawTable(transactions, nil, false)
		// choose a transaction based on no.
		fmt.Print("Select the No. of transaction you want to verify: ")
		var selection int
		fmt.Scan(&selection)
		if selection >= len(transactions) {
			message = "Your selection is out of range. Press Enter to try again."
			return message, true
		}
		// verify the transaction
		isSuccess := w.VerifyTransaction(transactions[selection])
		fmt.Println("Your transaction is being verified. Please be patient.")
		if !isSuccess {
			message = "Failed to verify transaction. Please try again. Press Enter to continue."
		} else {
			message = "Verify transaction successfully. Press Enter to continue."
		}
	// case 4:
	// 	fmt.Println("Print block information selected")
	// 	fmt.Printf("Type the address of block here: ")

	// 	var bIndex int
	// 	_, err := fmt.Scanln(&bIndex)
	// 	if err != nil {
	// 		log.Println(err)
	// 		fmt.Println("An error occurs. Press Enter to continue")
	// 		fmt.Scanln()
	// 	}
	// 	isSuccess := w.PrintBlock(bIndex)
	// 	if !isSuccess {
	// 		fmt.Println("Press Enter to continue")
	// 		fmt.Scanln()
	// 	}
	case 0:
		message = "Exit selected"
		w.Finish()
		return message, false
	default:
		message = "Your choice is invalid. Press Enter to try again"
		return message, true
	}

	return message, true
}
