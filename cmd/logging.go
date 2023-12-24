/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// loggingCmd represents the logging command
var loggingCmd = &cobra.Command{
	Use:   "login",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Client started")
		reader := bufio.NewReader(os.Stdin)
		wallet := new(Wallet)
		wallet.init(loggingEnabled)
	},
}

func init() {
	rootCmd.AddCommand(loggingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loggingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loggingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
