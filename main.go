// package main

// import "github.com/Horiodino/terrago/web"

// // networks "github.com/Horiodino/terrago/networks"

// func main() {

// 	// 	metrices.Getinfo()

// 	// 	metrices.Display()
// 	// 	metrices.Cpu()
// 	// 	// network.Nic_info()

// 	// 	// network.IncomingTraffic()
// 	// 	// network.Outbound_Traffic()
// 	// 	// network.DeepPacketInspection()

// 	// 	/*

// 	// 		network.GetEndpoints()

// 	// 		// alerts.PodFailure()

// 	// 		// localhost.Host()

// go networks.AcepptRequest("Node 1", "localhost:8010")

// go networks.AcepptRequest("Node 2", "localhost:8011")
// fmt.Print("\rPacket Received:", networks.DATAVAR)
// select {}

// 	// Send data to a specific node
// networks.SendNicInfo("localhost:8021", "hiii i am node 1")

// 	// 	*/

// 	// 	// sendData("localhost:8001", "")

// 	// 	// Call GetContainerMetrics to retrieve container metrics
// 	// 	// container.GetContainerMetrics()

// 	// }

// }

// 	web.Startweb()

// }

package main

import (
	"github.com/Horiodino/terrago/cmd"
)

// "github.com/Horiodino/terrago/cmd"

func main() {

	cmd.Execute()
}
