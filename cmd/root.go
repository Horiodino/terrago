/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	rgname string
)

// rootCmd represents the base command when called without any subcommands
var terrago = &cobra.Command{
	Use:  "terrago",
	Long: `terrago is a CLI tool for Terraform`,

	Run: func(cmd *cobra.Command, args []string) {
		template := `resource "azurerm_resource_group" "{{.RgName}}" {
			name     = "{{.RgName}}"
			location = "eastus"
		}`
		temp, err := template.New("template").Parse(template)
		if err != nil {
			log.Fatal(err)
		}
		//define the data to be passed to the template
		data := struct {
			RGname string
		}{
			RGname: rgname,
		}

		//create a new file and write the template to it
		var result string
		if err := temp.Execute(&result, data); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		file, err := os.Create("main.tf")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		if _, err := file.WriteString(result); err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		fmt.Println("Terraform template created")
	},
}

func init() {

	terrago.PersistentFlags().StringVarP(&rgname, "rgname", "n ", "", "Resource Group Name")
	//here we are making the flag required for the command to run successfully
	terrago.MarkPersistentFlagRequired("rgname")
}
