package cmd

import (
	"fmt"

	"github.com/Azure/azure-sdk-for-go/profiles/2017-03-09/compute/mgmt/compute"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-05-01/resources"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/spf13/cobra"
)

var (
	subscriptionID string
	rgName         string
	vmname         string
	vmSize         string
	location       string
	username       string
	password       string
)

var vmCmd = &cobra.Command{
	Use:   "vm",
	Short: "Manage Azure Virtual Machines",
	Long:  `Manage Azure Virtual Machines`,
	Run: func(cmd *cobra.Command, args []string) {
		createVM(vmname, vmSize, location, username, password, subscriptionID, rgName)
	},
}

func init() {
	vmCmd.AddCommand(vmCmd)
	vmCmd.PersistentFlags().StringVarP(&subscriptionID, "subscriptionID", "s", "", "Azure Subscription ID")
	vmCmd.PersistentFlags().StringVarP(&rgName, "rgName", "g", "", "Azure Resource Group Name")
	vmCmd.PersistentFlags().StringVarP(&vmname, "vmname", "n", "", "Azure Virtual Machine Name")
	vmCmd.PersistentFlags().StringVarP(&vmSize, "vmSize", "z", "", "Azure Virtual Machine Size")
	vmCmd.PersistentFlags().StringVarP(&location, "location", "l", "", "Azure Location")
	vmCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "Azure Virtual Machine Username")
	vmCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "Azure Virtual Machine Password")

}

func createVM(vmname string, vmSize string, location string, username string, password string, subscriptionID string, rgName string) {
	authorizer, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		panic(err)
	}
	groupsClient := resources.NewGroupsClient(subscriptionID)
	groupsClient.Authorizer = authorizer
	group, err := groupsClient.CreateOrUpdate(ctx, rgName, resources.Group{Location: &location})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created resource group %q", *group.Name)

	//creating the vm
	vmClient := compute.NewVirtualMachinesClient(subscriptionID)
	vmClient.Authorizer = authorizer
	future, err := vmClient.CreateOrUpdate(
		ctx,
		rgName,
		vmname,
		compute.VirtualMachine{
			Location : &location,
			//  using the virtual machine properties
			VirtualMachineProperties: &conpute.VirtualMachineProperties{
				//  using the hardware profile to define the vm size
				HardwareProfile: &compute.HardwareProfile{
					VMSize: compute.VirtualMachineSizeTypes(vmSize),
				},

				// its a linux vm so here we are using the linux configuration
				StorageProfile: &compute.StorageProfile{
					//here we are using the image reference to create the vm
					ImageReference: &compute.ImageReference{
						Publisher: to.StringPtr("Canonical"),
						Offer:     to.StringPtr("UbuntuServer"),
						//sku means the version of the os
						Sku:       to.StringPtr("16.04-LTS"),
						Version:   to.StringPtr("latest"),
					},
				},

				//os profile will have the username password and the ssh keys
				//here linuxconfiguration is used to disable the password authentication
				//as azure will not allow the password authentication and the ssh keys
				OsProfile: &compute.OSProfile{
					ComputerNAme: to.StringPtr(vmname),
					AdminUsername: to.StringPtr(username),
					AdminPassword: to.StringPtr(password),
					LinuxConfiguration: &compute.LinuxConfiguration{
						//disabling the password authentication as we are using the ssh keys
						DisablePasswordAuthentication: to.BoolPtr(false),
						//here we are using the ssh keys to login to the vm
						SSh: &compute.SShConfiguration{
							PublicKeys: &[]compute.SShPublicKey{
								{
									//now here we are using the public key to login to the vm
						},
			},
			
}
