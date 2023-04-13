package azure

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-09-01/network"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

func network_profile() {
	// Create a new authorizer
	authorizer, err := auth.NewAuthorizerFromCLI()
	if err != nil {
		panic(err)
	}

	//network profile client
	networkClient := network.NewProfilesClient("subscriptionID")
	networkClient.Authorizer = authorizer

	//create a new network profile

}
