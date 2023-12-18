/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/gavotte25/blockchain_lab1/server"
	"github.com/spf13/cobra"
)

// addBlockCmd represents the addBlock command
var addBlockCmd = &cobra.Command{
	Use:   "addblock",
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
			// Add transaction data
			if err == nil {
				transaction_data, _ := cmd.Flags().GetString("data")
				flag := bc.AddBlock(transaction_data)
				if !flag {
					message := fmt.Errorf("add block does not meet the constraint of the minimum %d transactions", server.MinBlockTransactions)
					log.Fatal(message)
				}
			}
			// Write the blockchain back to file
			error := server.WriteJSONToFile(blockNameJSON, bc)
			if error == nil {
				fmt.Printf("Adding a block to blockchain with name %s successfully.\n", blockName)
			} else {
				fmt.Println("An error occurs when create a new blockchain.")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(addBlockCmd)
	addBlockCmd.Flags().StringP("name", "n", "", "The name of the blockchain")
	addBlockCmd.Flags().StringP("data", "d", "", "The data of transaction")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addBlockCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addBlockCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
