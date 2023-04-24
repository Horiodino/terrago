package azure

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/containerservice/mgmt/2019-11-01/containerservice"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/Azure/go-autorest/autorest/to"
)

var (
	clusterName = "clusterName"
)

func clustercreate() {
	//authenticating to the Azure API
	authorizer, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	//creating the client for cluster
	//authenticating to the Azure API
	//creating the cluster client
	clusterClient := containerservice.NewManagedClustersClient(subscriptionID)
	clusterClient.Authorizer = authorizer

	//creating the cluster
	_, err = clusterClient.CreateOrUpdate(context.Background(), resourceGroupName, clusterName, containerservice.ManagedCluster{
		Location: to.StringPtr("location"),
		//declaring the cluster properties
		ManagedClusterProperties: &containerservice.ManagedClusterProperties{
			KubernetesVersion: to.StringPtr("1.15.7"),
			DNSPrefix:         to.StringPtr("dnsPrefix"),
			AgentPoolProfiles: &[]containerservice.ManagedClusterAgentPoolProfile{
				{
					Name:              to.StringPtr("agentPoolName"),
					Count:             to.Int32Ptr(1),
					VMSize:            containerservice.StandardDS2V2,
					OsDiskSizeGB:      to.Int32Ptr(30),
					OsType:            containerservice.Linux,
					VnetSubnetID:      to.StringPtr("vnetSubnetID"),
					EnableAutoScaling: to.BoolPtr(true),
					// AgentPoolType:     containerservice.VirtualMachineScaleSets,
					AvailabilityZones: &[]string{"1", "2", "3"},
				},
			},
			MaxAgentPools : to.Int32Ptr(1),
			LinuxProfile: &containerservice.LinuxProfile{
				AdminUsername: to.StringPtr("adminUsername"),
				SSH: &containerservice.SSHConfiguration{
					PublicKeys: &[]containerservice.SSHPublicKey{
						{
							KeyData: to.StringPtr("sshPublicKey"),
						},
					},
				},
			},

			ServicePrincipalProfile : &containerservice.ManagedClusterServicePrincipalProfile{
				ClientID: to.StringPtr("clientID"),
				Secret: to.StringPtr("clientSecret"),
			},

			AddonProfiles : &map[string]*containerservice.ManagedClusterAddonProfile{
				"omsagent": {
					Enabled: to.BoolPtr(true),
					Config: &map[string]string{
						"logAnalyticsWorkspaceResourceID": "logAnalyticsWorkspaceResourceID",
					},
				},
			NodeResourceGroup : to.StringPtr("nodeResourceGroup"),

			EnableRBAC : to.BoolPtr(true),
			EnablePodSecurityPolicy : to.BoolPtr(true),
			NetworkProfile : &containerservice.ContainerServiceNetworkProfile{
				NetworkPlugin: containerservice.Azure,
				NetworkPolicy: containerservice.Azure,
			},
			AadProfile : &containerservice.ManagedClusterAADProfile{
				ClientAppID: to.StringPtr("clientAppID"),
				ServerAppID: to.StringPtr("serverAppID"),
				TenantID: to.StringPtr("tenantID"),
				AdminGroupObjectIDs: &[]string{"adminGroupObjectID"},
			},
			APIServerAccessProfile : &containerservice.ManagedClusterAPIServerAccessProfile{
				EnablePrivateCluster: to.BoolPtr(true),
			},
		},
		
		Identity : &containerservice.ManagedClusterIdentityProfile{
			"systemAssigned": {
				PrincipalID: to.StringPtr("principalID"),
				TenantID: to.StringPtr("tenantID"),
				Type : ResourceIdentityType("systemAssigned"),

			},
		},
		ID : to.StringPtr("id"),
		Name : to.StringPtr("name"),
		Type : to.StringPtr("type"),
		Location : to.StringPtr("location"),
	},
	)// here fix error 

	if err != nil {
		log.Fatal(err)
	}

	// as the cluster creation is an asynchronous operation, we need to wait for the cluster to be created
	// we will wait for the cluster to be in the Succeeded state
	// we will wait for a maximum of 4 minutes as the cluster creation can take up to 3 minutes
	// if the cluster is not in the Succeeded state after 4 minutes, we will return an error

	// we need to check if our cluster is in the created then we will return nil
	// else we will return an error

	//checking if the cluster is in the created state

	//here we will get the cluster details
	cluster, err := clusterClient.Get(context.Background(), resourceGroupName, clusterName)
	if err != nil {
		log.Fatal(err)
	}

	//here we will check if the cluster is in the Succeeded state
	if cluster.ManagedClusterProperties.ProvisioningState == containerservice.Succeeded {
		return nil
	}

}

// getting all the clusters in the resource group
func clusterlist() {
	//cuthenticatuing to the Azure API
	authorizer, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	//creating the cluster client
	clusterClient := containerservice.NewManagedClustersClient(subscriptionID)
	clusterClient.Authorizer = authorizer

	//getting all the clusters in the resource group
	clusters, err := clusterClient.ListByResourceGroup(context.Background(), resourceGroupName)
	if err != nil {
		log.Fatal(err)
	}

	//printing the cluster details
	for _, cluster := range clusters.Values() {
		//printing the cluster name
		fmt.Println(*cluster.Name)

		//printing the cluster location
		fmt.Println(*cluster.Location)

		//printing the cluster kubernetes version
		fmt.Println(*cluster.ManagedClusterProperties.KubernetesVersion)
	}
}


// getting the cluster details

func clusterget() {
	//authenticating to the Azure API
	authorizer, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	//creating the cluster client
	clusterClient := containerservice.NewManagedClustersClient(subscriptionID)
	clusterClient.Authorizer = authorizer

	//getting the cluster details
	cluster, err := clusterClient.Get(context.Background(), resourceGroupName, clusterName)
	if err != nil {
		log.Fatal(err)
	}

	//printing the cluster details
	fmt.Println(*cluster.Name)
	fmt.Println(*cluster.Location)
	fmt.Println(*cluster.ManagedClusterProperties.KubernetesVersion)

	//printing the agent pool details
	for _, agentPool := range *cluster.ManagedClusterProperties.AgentPoolProfiles {
		fmt.Println(*agentPool.Name)
		fmt.Println(*agentPool.Count)
		fmt.Println(agentPool.VMSize)
		fmt.Println(*agentPool.OsDiskSizeGB)
		fmt.Println(agentPool.OsType)
		fmt.Println(*agentPool.VnetSubnetID)
		fmt.Println(*agentPool.EnableAutoScaling)
		fmt.Println(*agentPool.AvailabilityZones)
	}

	//printing the linux profile details
	fmt.Println(*cluster.ManagedClusterProperties.LinuxProfile.AdminUsername)
}

// deleting the cluster from the resource group and also from the whole subscription

func clusterdelete() {

	//authenticating to the Azure API
	authorizer, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	//creating the cluster client
	clusterClient := containerservice.NewManagedClustersClient(subscriptionID)
	clusterClient.Authorizer = authorizer

	//deleting the cluster

	_, err = clusterClient.Delete(context.Background(), resourceGroupName, clusterName) // here fix error regarding delete
	if err != nil {
		log.Fatal(err)
	}

	//checking if the cluster is deleted
	_, err = clusterClient.Get(context.Background(), resourceGroup(), clusterName)
	if err != nil {
		log.Fatal(err)
	}
	//if the cluster is not deleted, we will return an error above
}

// also we need to get the kubeconfig file for the cluster
// we will get the kubeconfig file for the cluster and we will save it in a file

// getting the kubeconfig file for the cluster

func getkubeconfig() {
	// authenticating to the Azure API
	authorizer, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	// creating the cluster client
	clusterClient := containerservice.NewManagedClustersClient(subscriptionID)
	clusterClient.Authorizer = authorizer

	// getting the kubeconfig file for the cluster
	accessProfile, err := clusterClient.GetAccessProfile(context.Background(), resourceGroupName, clusterName, "clusterUser")
	if err != nil {
		log.Fatal(err)
	}

	// saving the kubeconfig file in a file
	err = ioutil.WriteFile("kubeconfig", *accessProfile.KubeConfig, 0644)
	if err != nil {
		log.Fatal(err)
	}
}




// we will get all the cluster resources info such as the cpu usage, memory usage, disk usage,
// network usage, pod usage, container usage, and the incoming traffic and outgoing traffic
// fo this we will use kubeconfig file for the cluster

// getting the cluster resources info

func clusterresources() {

}
