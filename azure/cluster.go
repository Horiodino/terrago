package azure

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func clustercreate(clustername string, rgname string, rbac bool, location string, nodepool int, os_type string, dnsprefix string) {
	fmt.Println("Creating cluster: ", clustername, " in resource group: ", rgname)
	t := `resource "azurerm_kubernetes_cluster" "example" {
		name                = "{{.Clustername}}"
		location            = "{{.Location}}"
		resource_group_name = "{{.RGname}}"
		dns_prefix          = "{{.dns_prefix}}"
		kubernetes_version  = "1.21.3"
	  
		agent_pool_profile {
		  name            = "default"
		  count           = "{{.nodepool}}"
		  vm_size         = "Standard_DS2_v2"
		  os_type         = "{{.os_type}}"
		  vnet_subnet_id  = "/subscriptions/subscription_id/resourceGroups/rg_name/providers/Microsoft.Network/virtualNetworks/vnet_name/subnets/aks-subnet"
		}
	  
		service_principal {
		  client_id     = "service_principal_client_id"
		  client_secret = "service_principal_client_secret"
		}
	  
		role_based_access_control {
		  enabled = "{{.RBAC}}"
		}
	  
		tags = {
		  Environment = "{{.Environment}}}"
		}`

	temp, err := template.New("template").Parse(t)
	if err != nil {
		log.Fatal(err)
	}

	// issue
	// if resourceGroup == "" {
	// 	fmt.Printf("Resource group not specified using default resource group: %s", rgname)
	// 	resourceGroup = rgname
	// 	location = "East US"
	// }

	data := struct {
		Clustername string
		Location    string
		nodepool    int
		RGname      string
		dns_prefix  string
	}{
		Clustername: clustername,
		Location:    location,
		RGname:      rgname,
		nodepool:    nodepool,
		dns_prefix:  dnsprefix,
	}

	var result string
	if err := temp.Execute(&result, data); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var file, _ err = os.Create("cluster.tf")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if _, err := file.WriteString(result); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Println("File Created Successfully", file.Name())
}
