package Containermonitoring

import (
	"fmt"

	"github.com/google/cadvisor/client"
	v1 "github.com/google/cadvisor/info/v1"
	"golang.org/x/net/context"
)

// defining a strunct to store the container metrics and all
type containerMetrics struct {
	// conatainer name of type var containerNames
	cpuUsage    []int
	memoryUsage []int
	diskIo      []int
	networkTx   []int
	networkRx   []int
}

// getting all the container names
// and storing them in a slice
// slice declaration
var containerNames []string
var containerN string

func getContainerNames() {
	client, err := client.NewClient("http://localhost:8080")
	if err != nil {
		fmt.Println("Error creating cAdvisor client:", err)
		return
	}

	containerName := containerN // Replace with the name of the container you want to get stats for
	query := v1.ContainerInfoRequest{NumStats: 1}
	info, err := client.ContainerInfo(context.Background(), containerName, &query) // <--- getting error here
	if err != nil {
		fmt.Println("Error getting container info:", err)
		return
	}

	fmt.Println("Container name:", info.ContainerReference.Name)
	for _, stat := range info.Stats {
		fmt.Println("CPU usage:", stat.Cpu.Usage.Total)
		fmt.Println("Memory usage:", stat.Memory.Usage)
	}
}
