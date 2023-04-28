/*
Copyright © 2023 holiodin <holiodin@gmaiil.com>
*/
package main

import (
	"github.com/Horiodino/terrago/monitoring/Containermonitoring"
)

// webgo-->monitoring
//
//	|            | |-->Containermonitoring
//	|		    |           |-->metrices.go
//	|			|----metrices.go
//	|----main.go
func main() {

	// Containermonitoring.getContainerNames()

	metrices.cpu()
	metrices.clusterInfo()

}
