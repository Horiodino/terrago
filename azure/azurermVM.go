package azure

import (
	"github.com/spf13/cobra"
)

var (
	name      string
	rgname    string
	location  string
	os_type   string
	size      string
	disk_size int
	username  string
	password  string
)

var vmcmd = &cobra.Command{
	Use:   "vm",
	Short: "Create a VM",
	Long:  `Create a VM`,
	Run: func(cmd *cobra.Command, args []string) {
		vmcreate(name, rgname, location, os_type, size, disk_size, username, password)
	},
}

func vmcreate(name string, rgname string, location string, os_type string, size string, disk_size int, username string, password string) {

}
