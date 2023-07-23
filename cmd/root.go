package cmd

import (
	"fmt"
	"os"
	"os/exec"

	metrices "github.com/Horiodino/terrago/cluster-metrices"
	"github.com/Horiodino/terrago/networks"
	"github.com/fatih/color"
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
var Getendpoint = &cobra.Command{
	Use:   "getendpoint",
	Short: "Get endpoint",
	Run: func(cmd *cobra.Command, args []string) {
		networks.GetEndpoints()
	},
}
var Startserver = &cobra.Command{
	Use:   "init",
	Short: "deploys metrics server",
	Run: func(cmd *cobra.Command, args []string) {
		metriceserver()
	},
}

func init() {
	rootCmd.AddCommand(clusterInfoCmd)
	rootCmd.AddCommand(Nodeinfo)
	rootCmd.AddCommand(Objects)
	rootCmd.AddCommand(Nsinfo)
	rootCmd.AddCommand(Startserver)
	rootCmd.AddCommand(Network)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func metriceserver() {
	color.Blue("Deploying metrics server...")
	cmd := exec.Command("kubectl", "apply", "-f", "https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	fmt.Println()
	metriceserverstate()

	color.Blue("now you can run 'terrago clusterinfo' to get cluster information")

}
