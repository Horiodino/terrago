package azure

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	clusterName   string
	resourceGroup string
	location      string
	nodepool      int
	os_type       string
	dnsprefix     string
	rbac          bool
)

var k8sCmd = &cobra.Command{
	Use:   "k8s",
	Short: "Kubernetes commands",
	Long:  `Kubernetes commands`,

	Run: func(cmd *cobra.Command, args []string) {
		clustercreate(clusterName, resourceGroup, rbac, location, nodepool, os_type, dnsprefix)
	},
}

func init() {
	k8sCmd.Flags().StringVar(&clusterName, "name", "", "Cluster name")
	k8sCmd.MarkFlagRequired("name")
	k8sCmd.Flags().StringVar(&resourceGroup, "rg", "", "Resource group name")
	k8sCmd.Flags().StringVar(&location, "location", "", "Location for the cluster")
	k8sCmd.MarkFlagRequired("location")
	k8sCmd.Flags().StringVar(&os_type, "os_type", "", "OS type for the cluster")
	k8sCmd.MarkFlagRequired("os_type")
	k8sCmd.Flags().StringVar(&dnsprefix, "dns_prefix", "", "DNS prefix for the cluster")
	k8sCmd.MarkFlagRequired("dns_prefix")
	k8sCmd.Flags().IntVar(&nodepool, "nodepool", 0, "Number of nodes in the cluster")
	k8sCmd.MarkFlagRequired("nodepool")
	k8sCmd.Flags().BoolVar(&rbac, "rbac", false, "Enable RBAC")
	k8sCmd.MarkFlagRequired("rbac")
}

func main() {
	rootCmd := &cobra.Command{Use: "main"}
	rootCmd.AddCommand(k8sCmd)
	rootCmd.Execute()
}
