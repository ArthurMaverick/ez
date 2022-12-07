/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	iac "github.com/ArthurMaverick/ez/package/template"
	// "github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	TemplateValue string
	Provider      string
)

var infraCmd = &cobra.Command{
	Use:   "infra",
	Short: "generate cloudformation templates",
	Long: `How use:

	for example use. ez infra create --template vpc`,

	Run: func(cmd *cobra.Command, args []string) {
		method := iac.Template{}
		method.GetEndpoints()

		if Provider == "tf" {
			method.GenerateTerraformModules(TemplateValue)
		}
		if Provider == "cfn" {
			method.GenerateClouformationTemplates(TemplateValue)
		}

	},
}

func init() {
	infraCmd.Flags().StringVarP(&TemplateValue, "create", "c", "", "Template Name")
	infraCmd.Flags().StringVarP(&Provider, "provider", "p", "", "Iac Provider")
	rootCmd.AddCommand(infraCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cfnCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cfnCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// cfnCmd.Flags().String("create", "", "Template Name")
}
