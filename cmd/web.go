package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	metrices "github.com/Horiodino/terrago/cluster-metrices"
	"github.com/spf13/cobra"
)

var displayedNodes = make(map[string]bool)
var NISLICE = metrices.NodeInfoList
var allData AllData

var WebCmd = &cobra.Command{
	Use:   "dashboard",
	Short: "Print cluster information",
	Run: func(cmd *cobra.Command, args []string) {
		metrices.Getinfo()
		metrices.Cpu()
		allData = AllData{
			Monitoring: metrices.Monitoring{
				ClusterName: metrices.M.ClusterName,
				Cpu:         metrices.M.Cpu,
				Cores:       metrices.M.Cores,
				Nodes:       metrices.M.Nodes,
				Totalmemory: metrices.M.Totalmemory,
				Usedmemory:  metrices.M.Usedmemory,
				Disk:        metrices.M.Disk,
				Totaldisk:   metrices.M.Totaldisk,
				Billing:     metrices.M.Billing,
			},
			NodeInfoList: NISLICE,
		}

		go func() {
			for {
				metrices.Getinfo()
				metrices.Cpu()
				updateMonitoringData()
				time.Sleep(5 * time.Second)
			}
		}()

		http.HandleFunc("/data", handleData)
		http.Handle("/", http.FileServer(http.Dir(".")))
		fmt.Println("Listening on port http://localhost:8810")
		http.ListenAndServe(":8810", nil)
	},
}

type AllData struct {
	Monitoring     metrices.Monitoring
	NodeInfoList   []metrices.NodeInfo
	ResourcesList  []metrices.Resources
	NamespacesList []metrices.Namespacestruct
}

func handleData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// If displayedNodes is empty, send allData directly
	if len(displayedNodes) == 0 {
		json.NewEncoder(w).Encode(allData)
		return
	}

	// Filter out the nodes that have already been displayed
	newNodes := make([]metrices.NodeInfo, 0)
	for _, node := range allData.NodeInfoList {
		if !displayedNodes[node.Name[0]] {
			newNodes = append(newNodes, node)
			displayedNodes[node.Name[0]] = true
		}
	}

	// Send only the new nodes in the response
	newData := AllData{
		Monitoring:   allData.Monitoring,
		NodeInfoList: newNodes,
	}
	json.NewEncoder(w).Encode(newData)
}

func updateMonitoringData() {
	// Simulate updating the monitoring data with new values
	allData.Monitoring.Cpu = metrices.M.Cpu
	allData.Monitoring.Usedmemory = metrices.M.Usedmemory

	// Update the NodeInfoList with the latest data
	allData.NodeInfoList = metrices.NodeInfoList
}
