/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/gavotte25/blockchain_lab1/server"
	"github.com/spf13/cobra"
)

// newchainCmd represents the newchain command
var createChainCmd = &cobra.Command{
	Use:   "createchain",
	Short: "Create a new blockchain",
	Long: `createchain is a command that creates a new blockchain.

This command creates a new blockchain with the specified name. It initializes the blockchain
with a genesis block and sets it as the current active blockchain.

Example:
	blockchain_lab1 createchain --name (-n) myblockchain

This will create a new blockchain named "myblockchain".`,
	Run: func(cmd *cobra.Command, args []string) {
		blockName, _ := cmd.Flags().GetString("name")
		server.CreateBlockchain(blockName)
	},
}

func init() {
	rootCmd.AddCommand(createChainCmd)
	createChainCmd.Flags().StringP("name", "n", "", "The name of the blockchain you want to create")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newchainCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newchainCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
