/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
// access_key string = os.Getenv("AWS_ACCESS_KEY_ID")
// secret_key string = os.Getenv("AWS_SECRET_ACCESS_KEY")
// cfgFile  string

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
		fmt.Println(viper.GetString("access"))
		fmt.Println(viper.GetString("secret"))

	},
}

func init() {
	rootCmd.AddCommand(ssmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	// rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	// viper.BindPFlag("username", rootCmd.PersistentFlags().Lookup("username"))
	// viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ssmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
