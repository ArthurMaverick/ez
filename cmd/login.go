/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// loginCmd represents the login command
var (
	access   string
	secret   string
	loginCmd = &cobra.Command{
		Use:   "login",
		Short: "sign in to use a cloud services ( only aws available )",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(access)
		},
	}
)

func init() {
	loginCmd.Flags().StringVarP(&access, "access", "a", "", "Template Name")
	loginCmd.Flags().StringVarP(&secret, "secret", "s", "", "Iac Provider")

	rootCmd.AddCommand(loginCmd)

	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()

	value := viper.Get("AWS_SECRET_KEY")
	access = value.(string)
	// i := viper.Get("AWS_SECRET_KEY")
	// access = i.(string)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
