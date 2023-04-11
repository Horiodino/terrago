package cmd

import (
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-05-01/resources"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/spf13/cobra"
)

func NewCreateResourceGroupCmd() *cobra.Command {
	var rgName string

	cmd := &cobra.Command{
		Use:   "create-resource-group",
		Short: "Create a new resource group in Azure.",

		RunE: func(cmd *cobra.Command, args []string) error {
			//azure needs a subscription ID to create a resource group
			//we can get this from the environment variables
			//we also need to autenticate to azure using the environment variables

			//now we are using all the predefined functions from the azure sdk
			authoriser, err := 

		},
	}
}
