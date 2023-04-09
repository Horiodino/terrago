package azure

import (
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

var cluster = &cobra.Command{
	Use:   "k8s",
	Short: "Kubernetes commands",
	Long:  `Kubernetes commands`,

	Run: func(cmd *cobra.Command, args []string) {
		clustercreate(clusterName, resourceGroup, rbac, location, nodepool, os_type, dnsprefix)
	},
}

func init() {
	cluster.Flags().StringVar(&clusterName, "name", "", "Cluster name")
	cluster.MarkFlagRequired("name")
	cluster.Flags().StringVar(&resourceGroup, "rg", "", "Resource group name")
	cluster.Flags().StringVar(&location, "location", "", "Location for the cluster")
	cluster.MarkFlagRequired("location")
	cluster.Flags().StringVar(&os_type, "os_type", "", "OS type for the cluster")
	cluster.MarkFlagRequired("os_type")
	cluster.Flags().StringVar(&dnsprefix, "dns_prefix", "", "DNS prefix for the cluster")
	cluster.MarkFlagRequired("dns_prefix")
	cluster.Flags().IntVar(&nodepool, "nodepool", 0, "Number of nodes in the cluster")
	cluster.MarkFlagRequired("nodepool")
	cluster.Flags().BoolVar(&rbac, "rbac", false, "Enable RBAC")
	cluster.MarkFlagRequired("rbac")
}

func main() {
	rootCmd := &cobra.Command{Use: "main"}
	rootCmd.AddCommand(cluster)
	rootCmd.Execute()
}
