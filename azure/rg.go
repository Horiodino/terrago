package azure

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

var (
	rg_name string
)

// rgname represents the terrago command
var rgname = &cobra.Command{
	Use:  "rgname",
	Long: `rgname is a CLI tool for creating an Azure Resource Group in Terraform`,

	Run: func(cmd *cobra.Command, args []string) {
		t := `resource "azurerm_resource_group" "{{.RgName}}" {
			name     = "{{.RgName}}"
			location = "eastus"
		}`
		temp, err := template.New("template").Parse(t)
		if err != nil {
			log.Fatal(err)
		}
		//define the data to be passed to the template
		data := struct {
			RGname string
		}{
			RGname: rg_name,
		}

		//create a new buffer and write the template to it
		var buf bytes.Buffer
		if err := temp.Execute(&buf, data); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		file, err := os.Create("main.tf")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		if _, err := buf.WriteTo(file); err != nil {
			log.Fatal(err)
		}

		fmt.Println("Terraform template created")
	},
}

func init() {
	rgname.PersistentFlags().StringVarP(&rg_name, "rgname", "n", "", "Resource Group Name")
	//here we are making the flag required for the command to run successfully
	rgname.MarkPersistentFlagRequired("rgname")
}
