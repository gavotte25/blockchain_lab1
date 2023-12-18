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
var getNumberOfTransactionOnChainCmd = &cobra.Command{
	Use:   "numtransaction",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
