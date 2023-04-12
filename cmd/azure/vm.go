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

	// Create a virtual machine client
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

				//Disk is also important as it will define the storage of the vm
				//here we are using the os disk to define the storage of the vm
				OsDisk: &compute.OSDisk{
					Name:         &osDiskName,
					Caching:      compute.CachingTypesNone,
					CreateOption: compute.DiskCreateOptionTypes(osDiskCreateOption),
					// we not giving the size of the disk so it will take the default size
					DiskSizeGB:   &osDiskSizeGB,
					//here we are using the managed disk to define the storage of the vm
					// like which type of storage we are using standard or premium
					ManagedDisk: &compute.ManagedDiskParameters{
						StorageAccountType: compute.StorageAccountTypes(storageAccountType),
					},
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
									//path is the path where the public key will be stored
									Path: to.StringPtr(fmt.Sprintf("/home/%s/.ssh/authorized_keys", username)),
									//keydata is the actual public key which will be used to login to the vm
									KeyData: to.StringPtr(PublicKey),
						},
			},

		//here we are using the network profile to define the network interface
		//we are using the default vnet and subnet

		NetworkProfile: &compute.NetworkProfile{

			//defining the name of the network interface  


			NetworkInterfaces: &[]compute.NetworkInterfaceReference{
				//here we are using the network interface reference to define the network interface
				{

					//here we are using the network interface reference properties to define the network interface
					ID: to.StringPtr(fmt.Sprintf("/subscriptions", subscriptionID, "/resourceGroups", rgName, "/providers/Microsoft.Network/networkInterfaces", vmname, "-nic")),
					//now lets create the network interface
					NetworkInterfaceReferenceProperties: &compute.NetworkInterfaceReferenceProperties{
						//here we are using the primary to define the primary network interface
						//it simply means that this network interface will be used to connect to the vm
						Primary: to.BoolPtr(true),
						//here we are using the network interface properties to define the ip configuration
						NetworkInterfaceReferenceProperties: &compute.NetworkInterfaceReferenceProperties{
							IPConfigurations: &[]compute.IPConfiguration{
								{
									//now subnet is also important as it will define the ip address
									Subnet: &compute.APIEntityReference{
										ID: to.StringPtr(fmt.Sprintf("/subscriptions", subscriptionID, "/resourceGroups", rgName, "/providers/Microsoft.Network/virtualNetworks", vmname, "-vnet/subnets/default")),
									},

									//now here we are using the ip configuration properties to define the public ip address
									IPConfigurationProperties: &compute.IPConfigurationProperties{
										//here we are using the public ip address to define the public ip address
										PublicIPAddress: &compute.PublicIPAddress{
											ID: to.StringPtr(fmt.Sprintf("/subscriptions", subscriptionID, "/resourceGroups", rgName, "/providers/Microsoft.Network/publicIPAddresses", vmname, "-pip")),
											//now lets create the public ip address

											//publicIPAddressPropertiesFormat is used to define the public ip address
											//here we are using the publicIPAddressPropertiesFormat to define the public ip address
											PublicIPAddressPropertiesFormat: &compute.PublicIPAddressPropertiesFormat{
												PublicIPAddressVersion:   compute.IPv4,
												PublicIPAllocationMethod: compute.Static,
											},
										},
										},

									//NetworkSecurity Group
									NetworkSecurityGroup: &compute.NetworkSecurityGroup{
										ID: to.StringPtr(fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Network/networkSecurityGroups/%s-nsg", subscriptionID, rgName, vmname)),
										//now lets create the network security group
										networkSecurityGroupPropertiesFormat: &compute.NetworkSecurityGroupPropertiesFormat{
											SecurityRules: &[]compute.SecurityRule{
												{
													//here we are using the security rule properties to define the security rule
													//1st rule is to allow the ssh traffic
													SecurityRulePropertiesFormat: &compute.SecurityRulePropertiesFormat{
													//2nd protocol is to allow the ssh traffic
													Protocol: 			   compute.SecurityRuleProtocolTCP,
													//port range is to allow the ssh traffic
													DestinationPortRange: to.StringPtr(),
													//source address prefix is to allow the ssh traffic
													//it means that the traffic will be allowed from any ip address
													//but we can also specify the ip address from where the traffic will be allowed
													//here we are using the * to allow the traffic from any ip address
													//but it is not recommended to use the * as it will allow the traffic from any ip address
													SourceAddressPrefix:  to.StringPtr("*"),

													//now here we are using the access to allow the traffic
													//here we are using the allow to allow the traffic
													Access: compute.SecurityRuleAccessAllow,
													//priority is to allow the ssh traffic
													//why we are using the priority
													//because if we have multiple rules then the priority will decide which rule will be applied first

													Priority: to.Int32Ptr(100),
													//direction is to allow the ssh traffic
													//here we are using the inbound to allow the traffic from outside to inside
													Direction: compute.SecurityRuleDirectionInbound,
												},
										},
									},
								},
							},
						},
					},
				},
			},

			// now we created all the things required to create the vm
			//such as network interface, public ip address, network security group

			//now lets create the vm
		},
	},
}