/*
Copyright © 2023 holiodin <holiodin@gmaiil.com>
*/
package main

import (
	"github.com/Horiodino/terrago/monitoring"
)

func main() {

	// getting the metrices using the metrices.go file
	monitoring.cpu()
	monitoring.clusterInfo()
	monitoring.insertDataMongo()
	monitoring.namespacesInfo()

}
