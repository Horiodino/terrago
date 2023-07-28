package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	metrices "github.com/Horiodino/terrago/cluster-metrices"
	"github.com/spf13/cobra"
)

var allData AllData

var WebCmd = &cobra.Command{
	Use:   "dashboard",
	Short: "Print cluster information",
	Run: func(cmd *cobra.Command, args []string) {
		metrices.Getinfo()
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
			NodeInfoList: []metrices.NodeInfo{},
		}

		go func() {
			for {
				metrices.Getinfo()
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
	json.NewEncoder(w).Encode(allData)
}

func updateMonitoringData() {
	// Simulate updating the monitoring data with new values
	allData.Monitoring.Cpu = metrices.M.Cpu
	allData.Monitoring.Usedmemory = metrices.M.Usedmemory
}

// func updateNodeInfo() {
// 	// Simulate updating the node information with new values
// 	for i := range allData.NodeInfoList {
// 		allData.NodeInfoList[i].CPU = randFloat(1, 2)
// 		allData.NodeInfoList[i].Memory = randFloat(10, 20)
// 	}
// }

// func randFloat(min, max float64) float64 {
// 	return min + rand.Float64()*(max-min)
// }
