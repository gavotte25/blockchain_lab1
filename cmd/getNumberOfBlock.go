/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/gavotte25/blockchain_lab1/server"
	"github.com/spf13/cobra"
)

// getNumberOfBlockCmd represents the getNumberOfBlock command
// code implementation
var getNumberOfBlockCmd = &cobra.Command{
	Use:   "numblock",
	Short: "Get the number of blocks in a blockchain",
	Long: `Get the number of blocks in a blockchain. This command retrieves the total number of blocks in a specified blockchain.
	
		Example:
		$ blockchain_lab1 getNumberOfBlock --name (-n) myblockchain`,
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
				fmt.Printf("Number of blocks on Chain: %d\n", bc.GetNumberOfBlocks()) // Print number of transaction
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(getNumberOfBlockCmd)
	getNumberOfBlockCmd.Flags().StringP("name", "n", "", "The name of the blockchain")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getNumberOfBlockCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getNumberOfBlockCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
