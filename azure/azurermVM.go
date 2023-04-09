package azure

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	vm_name      string
	vm_rgname    string
	vm_location  string
	vm_os_type   string
	vm_size      string
	vm_disk_size int
	vm_username  string
	vm_password  string
)

var vmcmd = &cobra.Command{
	Use:   "vm",
	Short: "Create a VM",
	Long:  `Create a VM`,
	Run: func(cmd *cobra.Command, args []string) {
		vmcreate(vm_name, vm_rgname, vm_location, vm_os_type, vm_size, vm_disk_size, vm_username, vm_password)
	},
}

func vmcreate(name string, rgname string, location string, os_type string, size string, disk_size int, username string, password string) {

	fmt.Println("Creating VM: ", name, " in resource group: ", rgname, "size")

	t := `resource "azurerm_virtual_machine" "example" {
		name                  = "{{.name}}"
		location              = "{{.location}}"
		resource_group_name   = "{{.rgname}}"
		network_interface_ids = [azurerm_network_interface.example.id]
		vm_size               = "{{.size}}"
		os_profile            = "{{.os_type}}"
		`

}
