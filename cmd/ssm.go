/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	access_key string = os.Getenv("AWS_ACCESS_KEY_ID")
	secret_key string = os.Getenv("AWS_SECRET_ACCESS_KEY")
)

// ssmCmd represents the ssm command
var ssmCmd = &cobra.Command{
	Use:   "ssm",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ssm called")
	},
}

func init() {
	rootCmd.AddCommand(ssmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	ssmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ssmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
