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
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
