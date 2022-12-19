/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/ArthurMaverick/ez/package/Infra/network"
	"github.com/spf13/cobra"
)

var (
	IpAdress string
)

// netCmd represents the net command
var netCmd = &cobra.Command{
	Use:   "net",
	Short: "This command get info about ip adress",
	Long:  `This command get info about ip adress, for example: ez net -i YOUR_IP_ADRESS`,
	Run: func(cmd *cobra.Command, args []string) {
		network.GetIpInfo(IpAdress)
	},
}

func init() {
	netCmd.Flags().StringVarP(&IpAdress, "ip", "i", "", "Ip adress to get info")
	rootCmd.AddCommand(netCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// netCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// netCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
