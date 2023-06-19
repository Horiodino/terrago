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

	network.IncomingTraffic()
	// network.Outbound_Traffic()

	// alerts.PodFailure()

	// localhost.Host()

}
