/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/gavotte25/blockchain_lab1/server"
	"github.com/spf13/cobra"
)

// blockInfoCmd represents the blockInfo command
var chaininfoCmd = &cobra.Command{
	Use:   "chaininfo",
	Short: "Print the information of all blocks in the blockchain",
	Long: `Simply print the information of all blocks. The information contains "Block address", "Block size", "Time stamp", "Number of transaction".

	Usage:
		blockchain_lab1 chaininfo --name (-n) myblockchain
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
			// PrintInfo
			if err == nil {
				for index, block := range bc.BlockArr {
					fmt.Printf("Block: %d\n", index)
					block.PrintInfo()
					block.PrintTransaction()
					fmt.Println()
				}
			} else {
				fmt.Printf("An error occurs when access to blockchain with name %s. %s\n", blockName, err)
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(chaininfoCmd)
	chaininfoCmd.Flags().StringP("name", "n", "", "The name of the blockchain")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// blockInfoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// blockInfoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
