/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/gavotte25/blockchain_lab1/server"
	"github.com/spf13/cobra"
)

// deleteChainCmd represents the deleteChain command
// deletechainCmd represents the deletechain command
var deletechainCmd = &cobra.Command{
	Use:   "deletechain",
	Short: "Delete a blockchain by name",
	Long: `Delete a blockchain by name. This command allows you to delete a specific blockchain from the system.
	 Provide the name of the blockchain as an argument.

	Example:
	blockchain_lab1 deletechain --name (-n) myblockchain

	This will delete the blockchain with the name "myblockchain" from the system.`,
	Run: func(cmd *cobra.Command, args []string) {
		blockName, _ := cmd.Flags().GetString("name")
		server.DeleteBlockchain(blockName)
	},
}

func init() {
	rootCmd.AddCommand(deletechainCmd)
	deletechainCmd.Flags().StringP("name", "n", "", "The name of the blockchain you want to delete")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteChainCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteChainCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
