package localhost

import (
	"fmt"
	metrices "github.com/Horiodino/terrago/monitoring"
	"net/http"
	// alerts "github.com/Horiodino/terrago/Alerts"
	// container "github.com/Horiodino/terrago/monitoring/container_metrices"
)

func Host() {

	http.HandleFunc("/", Helloweb)
	fmt.Println("Starting the web server")

	server := http.Server{
		Addr:    ":8089",
		Handler: nil,
	}

	server.ListenAndServe()

	fmt.Println("Connection Closed")

}

// start the web server

func Helloweb(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1 style=\"color: #ffffff; background-color: #337ab7; font-family: Arial, sans-serif; text-align: center; font-size: 36px; padding: 10px;\">Hello Web</h1>")
	fmt.Fprintf(w, "<a href=\"/dashboard\">Dashboard</a>")
	fmt.Fprintf(w, "<a href=\"/logs\">Logs</a>")
	fmt.Fprintf(w, "<a href=\"/nodes\">Nodes</a>")
	fmt.Fprintf(w, "<a href=\"/\">Dashboard</a>")

	// M.cpu = float64(cpuUsage)
	// M.cores = cpuCores
	// M.nodes = int64(len(nodes.Items))
	// M.totalmemory = float64(memory)
	// M.usedmemory = float64(memoryUsage)
	// M.disk = float64(diskUsage)
	// M.totaldisk = float64(disk)
	// Access the monitoring struct data

	fmt.Fprintf(w, "<p>CPU Usage: %.2f</p>", "ssxsxs")
	// fmt.Fprintf(w, "<p>CPU Cores: %d</p>", metrices.M.Cores)
	// fmt.Fprintf(w, "<p>Nodes: %d</p>", metrices.M.Nodes)
	// fmt.Fprintf(w, "<p>Total Memory: %.2f</p>", metrices.M.Totalmemory)
	// fmt.Fprintf(w, "<p>Used Memory: %.2f</p>", metrices.M.Usedmemory)
	// fmt.Fprintf(w, "<p>Disk Usage: %.2f</p>", metrices.M.Disk)
	// fmt.Fprintf(w, "<p>Total Disk: %.2f</p>", metrices.M.Totaldisk)

}
