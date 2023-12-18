/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/gavotte25/blockchain_lab1/server"
	"github.com/spf13/cobra"
)

// deleteChainCmd represents the deleteChain command
var deletechainCmd = &cobra.Command{
	Use:   "deletechain",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
