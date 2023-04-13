package azure

import (
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/resources"
	"github.com/Azure/azure-sdk-for-go/services/containerservice/mgmt/2022-07-01/containerservice"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/Azure/go-autorest/autorest/to"
)

func CreateCluster(clustername string, newlocation string, nodecount int) {
	// Create a new authorizer

	if newlocation == "" {
		Create_RGroup(clustername, newlocation)
	}
	authorizer, err := auth.NewAuthorizerFromCLI()
	if err != nil {
		panic(err)
	}

	//kubernetes client

	//containerserive.NewManagedClustersClient("subscriptionID") will create a new client for the kubernetes cluster
	kubernetesClient := containerservice.NewManagedClustersClient("subscriptionID")
	kubernetesClient.Authorizer = authorizer

	netprofile = Create_NetworkProfile(clustername, newlocation)

	//create a new cluster
	Cluste := containerservice.ManagedCluster{
		Location: &newlocation,
		ManagedClusterProperties: &containerservice.ManagedClusterProperties{
			KubernetesVersion: to.StringPtr("1.21.2"),
			DNSPrefix:         to.StringPtr(clustername),
			AgentPoolProfiles: &[]containerservice.ManagedClusterAgentPoolProfile{
				{
					Name:   to.StringPtr("agentpool"),
					Count:  to.Int32Ptr(int32(nodecount)),
					VMSize: containerservice.StandardDS2V2,
					OsType: containerservice.Linux,
				},
			},

			//now network profile
			// we already defined network profile in the networkprofile.go file
			// so we will use that function here

			NetworkProfile: netprofile,

			// now defining storage profile
			// here we are using the default storage account

			StorageProfile: &containerservice.ManagedClusterStorageProfile{
				StorageAccountType: containerservice.ManagedClusterStorageAccountTypes("Standard_LRS"),
			},
			},

			// now we will create the cluster
			// we will use the kubernetesClient.CreateOrUpdate() function to create the cluster

		},


		_, err = kubernetesClient.CreateOrUpdate(context.Background(), "resourcegroupname", clustername, Cluster)
	if err != nil {
		panic(err)
	}



	fmt.Printf("Cluster %s is being created", clustername, time.Now().Format(time.RFC850))

	//now we will wait for the cluster to be created
	// we will use the kubernetesClient.Get() function to get the status of the cluster
	// this will tell us the status of the cluster
	_, err = kubernetesClient.Get(context.Background(), "resourcegroupname", clustername)
}
