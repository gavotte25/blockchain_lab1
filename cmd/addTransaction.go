/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/gavotte25/blockchain_lab1/server"
	"github.com/spf13/cobra"
)

// addTransactionCmd represents the addTransaction command
// It is used to add new transactions to an existing blockchain.
var addTransactionCmd = &cobra.Command{
	Use: "addtransaction",

	Short: "Add new transactions to an existing blockchain",
	Long: `The addtransaction command is used to add new transactions to an existing blockchain.
	It takes the name of the blockchain as an argument and the data for the transactions to be added.
	If the blockchain with the given name does not exist, it displays an error message.
	Otherwise, it reads the content of the blockchain, adds the new transactions, and writes the updated blockchain back to file.

	Usage:
		blockchain_lab1 addtransaction --name (-n) myblockchain --data (-d) "A send to B an amount of 2 BTC" "B send to C an amount of 3 BTC"
`,
	Run: func(cmd *cobra.Command, args []string) {
		blockName, _ := cmd.Flags().GetString("name")
		// Concate to create a json file
		blockNameJSON := blockName + ".json"
		// Check the existence of a blockchain with blockName
		isExist, _ := server.IsFileInDatabasePath(blockNameJSON)
		if !isExist {
			fmt.Printf("Blockchain with name %s does not exist. Use \"createchain\" before adding a block\n", blockName)
		} else {
			// If exists, read all the content of blockchain
			var err error
			bc, err := server.ReadJSONFromFile(blockNameJSON)
			// Add array transaction data by using function in server
			if err == nil {
				transactions, _ := cmd.Flags().GetStringArray("data")
				bc.AddTransactions(transactions)
			}
			// Write the blockchain back to file
			error := server.WriteJSONToFile(blockNameJSON, bc)
			if error == nil {
				fmt.Printf("Adding new transactions to the last block of the blockchain with name %s successfully.\n", blockName)
			} else {
				fmt.Println("An error occurs when adding new transaction.")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(addTransactionCmd)
	addTransactionCmd.Flags().StringP("name", "n", "", "The name of the blockchain")
	addTransactionCmd.Flags().StringArrayP("data", "d", []string{}, "The array of transactions")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addTransactionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addTransactionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
