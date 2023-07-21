package cmd

import (
	"fmt"
	"os"

	metrices "github.com/Horiodino/terrago/cluster-metrices"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "terrago",
	Short: "A sample Cobra CLI tool",
}

var clusterInfoCmd = &cobra.Command{
	Use:   "clusterinfo",
	Short: "Print cluster information",
	Run: func(cmd *cobra.Command, args []string) {
		metrices.Getinfo()
		metrices.Display()
	},
}
var Nodeinfo = &cobra.Command{
	Use:   "nodeinfo",
	Short: "Print node information",
	Run: func(cmd *cobra.Command, args []string) {
		metrices.Cpu()
		for _, node := range metrices.NodeInfoList {
			fmt.Println("")
			fmt.Println("|-------------------------------------|")
			fmt.Println("Node Name: ", node.Name)
			fmt.Println("CPU Usage: ", node.CPU)
			fmt.Println("Memory Usage: ", node.Memory)
			fmt.Println("Disk Usage: ", node.Disk)
			fmt.Println("CPU Temperature: ", node.CpuTemp)
			fmt.Println("IP: ", node.IP)
			fmt.Println("|-------------------------------------|")
		}
	},
}

var Nsinfo = &cobra.Command{
	Use:   "nsinfo",
	Short: "Print namespace information",
	Run: func(cmd *cobra.Command, args []string) {
		metrices.NamespacesInfo()
		metrices.Displaynsinfo()

	},
}

var Objects = &cobra.Command{
	Use:   "objects",
	Short: "Print objects information",
	Run: func(cmd *cobra.Command, args []string) {
		metrices.ClusterInfo()
		metrices.ObjectDisplay()
	},
}

func init() {
	rootCmd.AddCommand(clusterInfoCmd)
	rootCmd.AddCommand(Nodeinfo)
	rootCmd.AddCommand(Objects)
	rootCmd.AddCommand(Nsinfo)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
