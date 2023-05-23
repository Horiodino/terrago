package azure

import (
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-09-01/network"
	// "github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	// "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwo/rk/v2"
)

// define the network profile
func networkProfile() {
	// authenticating to the Azure API
	authorizer, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		panic(err)
	}

	// creating the client for network
	netclient := network.NewNetworkProfilesClient(subscriptionID)
	netclient.Authorizer = authorizer

}
