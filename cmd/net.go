package cmd

import (
	"github.com/Horiodino/terrago/networks"
	"github.com/spf13/cobra"
)

var Network = &cobra.Command{
	Use:   "network",
	Short: "Print network information",
	Run: func(cmd *cobra.Command, args []string) {
		go networks.AcepptRequest("Node 2", "localhost:8011")

	},
}
