package azure

import (
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-09-01/network"
	// "github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

var (
	loadBalancerName = "loadBalancerName"
	subscriptionID   = "subscriptionID"
)

// defining the load balancer
func loadbalancer() {
	// authenticating to the Azure API
	authorizer, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		panic(err)
	}

	// creating the client for load balancer
	lbclient := network.NewLoadBalancersClient(subscriptionID)
	lbclient.Authorizer = authorizer

	// creating the load balancer

}
