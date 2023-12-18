/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/gavotte25/blockchain_lab1/server"
	"github.com/spf13/cobra"
)

// getNumberOfTransactionOnChainCmd represents the getNumberOfTransactionOnChain command
// getNumberOfTransactionOnChainCmd represents the "numtransaction" command.
// It retrieves the number of transactions on a specified blockchain.
var getNumberOfTransactionOnChainCmd = &cobra.Command{
	Use:   "numtransaction",
	Short: "Retrieve the number of transactions on a blockchain",
	Long: `The "numtransaction" command retrieves the number of transactions on a specified blockchain.
It requires the name of the blockchain as a parameter.

	Usage:
		blockchain_lab1 numtransaction --name (-n) myblockchain

	This command will check if a blockchain with the specified name exists.
	If it exists, it will read the blockchain content and print the number of transactions on the chain.`,
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
			if err == nil {
				fmt.Printf("Number of Transactions on Chain: %d\n", bc.GetNumberOfTransactionsOnChain()) // Print number of transaction
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(getNumberOfTransactionOnChainCmd)
	getNumberOfTransactionOnChainCmd.Flags().StringP("name", "n", "", "The name of the blockchain")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getNumberOfTransactionOnChainCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getNumberOfTransactionOnChainCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
