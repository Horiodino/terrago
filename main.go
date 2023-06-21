package main

import (
	// alerts "github.com/Horiodino/terrago/Alerts"
	// metrices "github.com/Horiodino/terrago/monitoring"
	// container "github.com/Horiodino/terrago/monitoring/container_metrices"
	// localhost "github.com/Horiodino/terrago/localhost"
	network "github.com/Horiodino/terrago/network"
)

func main() {

	// alerts.DeploumentsFailure()
	// metrices.Getinfo()

	// metrices.Display()
	// network.Nic_info()

	// network.IncomingTraffic()
	// network.Outbound_Traffic()
	// network.DeepPacketInspection()
	network.GetEndpoints()

	// alerts.PodFailure()

	// localhost.Host()

	go network.AcepptRequest("Node 1", "localhost:8010")
	go network.AcepptRequest("Node 2", "localhost:8011")

	// Send data to a specific node
	network.SendNicInfo("localhost:8011", "hiii i am node 1")
	// sendData("localhost:8001", "")

}
