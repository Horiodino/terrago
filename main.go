package main

import (
	"fmt"

	alerts "github.com/Horiodino/terrago/Alerts"
	metrices "github.com/Horiodino/terrago/monitoring"
	container "github.com/Horiodino/terrago/monitoring/container_metrices"
)

func main() {
	metrices.Getinfo()
	metrices.Display()

	fmt.Println("=====================================")
	metrices.Cpu()
	metrices.ClusterInfo()
	fmt.Println("=====================================")

	metrices.NamespacesInfo()

	fmt.Println("=====================================")

	metrices.GetnamespaceInfoDetailed()

	container.Containermatricesinfo()
	fmt.Println("=====================================")
	alerts.DeploumentsFailure()
	alerts.Cpu_exceed()

}
