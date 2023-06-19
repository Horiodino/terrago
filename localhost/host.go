package localhost

import (
	"fmt"
	"net/http"
	"time"

	metrices "github.com/Horiodino/terrago/monitoring"
)

func Host() {
	// Start a goroutine to continuously update the metrics data
	go updateMetrics()

	http.HandleFunc("/", Helloweb)
	fmt.Println("Starting the web server")

	server := http.Server{
		Addr:    ":8089",
		Handler: nil,
	}

	server.ListenAndServe()

	fmt.Println("Connection Closed")
}

func updateMetrics() {
	for {
		metrices.Getinfo()

		// Sleep for a specific duration to control the update frequency
		time.Sleep(time.Second * 1)
	}
}

func Helloweb(w http.ResponseWriter, r *http.Request) {
	// Display the metrics data in the web server
	// fmt.Fprintf(w, "<h1 style=\"color: #ffffff; background-color: #337ab7; font-family: Arial, sans-serif; text-align: center; font-size: 36px; padding: 10px;\">Hello Web</h1>")
	// fmt.Fprintf(w, "<a href=\"/dashboard\">Dashboard</a>")
	// fmt.Fprintf(w, "<a href=\"/logs\">Logs</a>")
	// fmt.Fprintf(w, "<a href=\"/nodes\">Nodes</a>")
	// fmt.Fprintf(w, "<a href=\"/\">Dashboard</a>")

	// // Access the metrics data from the metrices package
	// fmt.Fprintf(w, "<p>CPU Usage: %.2f</p>", metrices.M.Cpu)
	// fmt.Fprintf(w, "<p>CPU Cores: %d</p>", metrices.M.Cores)
	// fmt.Fprintf(w, "<p>Nodes: %d</p>", metrices.M.Nodes)
	// fmt.Fprintf(w, "<p>Total Memory: %.2f</p>", metrices.M.Totalmemory)
	// fmt.Fprintf(w, "<p>Used Memory: %.2f</p>", metrices.M.Usedmemory)
	// fmt.Fprintf(w, "<p>Disk Usage: %.2f</p>", metrices.M.Disk)
	// fmt.Fprintf(w, "<p>Total Disk: %.2f</p>", metrices.M.Totaldisk)

}
