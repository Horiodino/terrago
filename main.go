package main

import (
	"net/http"
	"text/template"
)

func main() {

	// alerts.DeploumentsFailure()
	// metrices.Getinfo()

	// metrices.Display()
	// network.Nic_info()

	// network.IncomingTraffic()
	// network.Outbound_Traffic()
	// network.DeepPacketInspection()

	/*

		network.GetEndpoints()

		// alerts.PodFailure()

		// localhost.Host()

		go network.AcepptRequest("Node 1", "localhost:8010")
		go network.AcepptRequest("Node 2", "localhost:8011")

		// Send data to a specific node
		network.SendNicInfo("localhost:8011", "hiii i am node 1")

	*/

	// sendData("localhost:8001", "")

	// Call GetContainerMetrics to retrieve container metrics
	// container.GetContainerMetrics()

	// }

	type Person struct {
		Name    string
		Age     int
		Country string
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		person := Person{Name: "John Doe", Age: 25, Country: "USA"}

		tmpl := template.Must(template.ParseFiles("template.html"))

		err := tmpl.Execute(w, person)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.ListenAndServe(":8001", nil)
}
