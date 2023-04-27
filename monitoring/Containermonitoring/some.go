package main

import (
	"context"
	"fmt"
	"github.com/google/cadvisor/client/v2"
	"github.com/google/cadvisor/info/v1"
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

func getContainerNames() {
	// create a new client
	client, err := v2.NewClient("http://localhost:8080")
	if err != nil {
		// handle error
	}
	// now get the container names
	containerNames, err := client.AllDockerContainers(context.Background(), &v1.ContainerInfoRequest{})
	if err != nil {
		// handle error
	}

	for i, containerName := range containerNames {
		//append the container names to the slice
		containerNames = append(containerNames, containerName.Name)
	}
}

func getContainerMetrics() {

	client, err := v2.NewClient("http://localhost:8080")
	if err != nil {
		// handle error
	}
	//now get the cointainer stats

	containerStats, err := client.ContainerInfo(context.Background(), "/", &v1.ContainerInfoRequest{})
	if err != nil {
		// handle error
	}

	for _, containerStat := range containerStats {

		fmt.Printf("Container: %s\n", containerStat.Name)
		fmt.Printf("   ID: %s\n", containerStat.Id)
		fmt.Printf("   Image: %s\n", containerStat.Spec.Image)
		fmt.Printf("   Labels: %v\n", containerStat.Spec.Labels)
		fmt.Printf("   State: %v\n", containerStat.Spec.Labels)
		fmt.Printf("   Spec: %v\n", containerStat.Spec)
		fmt.Printf("   Stats: %v\n", containerStat.Stats)
		fmt.Printf("   Stats: %v\n", containerStat.Stats[0].Cpu.Usage.Total)
		fmt.Printf("   Stats: %v\n", containerStat.Stats[0].Memory.Usage)
		fmt.Printf("   Stats: %v\n", containerStat.Stats[0].DiskIo.IoServiceBytes)
		fmt.Printf("   Stats: %v\n", containerStat.Stats[0].Network.TxBytes)
		fmt.Printf("   Stats: %v\n", containerStat.Stats[0].Network.RxBytes)
		fmt.Printf("   Stats: %v\n", containerStat.Stats[0].Network.Interfaces)
		fmt.Printf("   Stats: %v\n", containerStat.Stats[0].Filesystem)
		fmt.Printf("   Stats: %v\n", containerStat.Stats[0].Accelerators)
		fmt.Printf("   Stats: %v\n", containerStat.Stats[0].Timestamp)
		fmt.Printf("   Stats: %v\n", containerStat.Stats[0].Errors)
		fmt.Printf("   Stats: %v\n", containerStat.Stats[0].Perf)
		fmt.Printf("   Stats: %v\n", containerStat.Stats[0].Cpu.Usage.PerCpu)
		fmt.Printf("   Stats: %v\n", containerStat.Stats[0].Cpu.Usage.User)
		fmt.Printf("   Stats: %v\n", containerStat.Stats[0].Cpu.Usage.System)
		fmt.Printf("   Stats: %v\n", containerStat.Stats[0].Cpu.Usage.PerCpu)
		fmt.Printf("   Stats: %v\n", containerStat.Stats[0].Cpu.Usage.PerCpu)
		fmt.Printf("   Stats: %v\n", containerStat.Stats[0].Cpu.Usage.PerCpu)
		fmt.Printf("   Stats: %v\n", containerStat.Stats[0].Cpu.Usage.PerCpu)
		fmt.Printf("   Stats: %v\n", containerStat.Stats[0].Cpu.Usage.PerCpu)
		fmt.Printf("   Stats: %v\n", containerStat.Stats[0].Cpu.Usage.PerCpu)
		fmt.Printf("   Stats: %v\n", containerStat.Stats[0].Cpu.Usage.PerCpu)
		fmt.Printf("   Stats: %v\n", containerStat.Stats[0].Cpu.Usage.PerCpu)
	}
	// getting the subcontainers for a particular container
	// what is subcontainers?
	// suppose you have a container running, and inside that container you have 3 more containers running
	// so the container running is the parent container and the 3 containers running inside it are the subcontainers
	// so this function will return the subcontainers for a particular container

	//  accoridng to the documentation, the first argument is the container name
	// we will fetch the container name from the func ALLDockerContainers
	//  we will use the name from slice of containerStats

	// now user can enter the container name
	// by seeing the container name, we can get the subcontainers for that particular container

	//user will enter the container name

	var containerName string
	fmt.Println("Enter the container name: ")
	fmt.Scanln(&containerName)

	subcontainers, err := client.SubcontainersInfo(context.Background(), containerName, &v1.ContainerInfoRequest{})
	if err != nil {
		// custom error handling
	}

	for _, subcontainer := range subcontainers {

		fmt.Printf("Subcontainer: %s\n", subcontainer.Name)
		fmt.Printf("   ID: %s\n", subcontainer.Id)
		fmt.Printf("   Image: %s\n", subcontainer.Spec.Image)
		fmt.Printf("   Labels: %v\n", subcontainer.Spec.Labels)
		fmt.Printf("   State: %v\n", subcontainer.Spec.Labels)
		fmt.Printf("   Spec: %v\n", subcontainer.Spec)
		fmt.Printf("   Stats: %v\n", subcontainer.Stats)
		fmt.Printf("   Stats: %v\n", subcontainer.Stats[0].Cpu.Usage.Total)
		fmt.Printf("   Stats: %v\n", subcontainer.Stats[0].Memory.Usage)
		fmt.Printf("   Stats: %v\n", subcontainer.Stats[0].DiskIo.IoServiceBytes)
		fmt.Printf("   Stats: %v\n", subcontainer.Stats[0].Network.TxBytes)
		fmt.Printf("   Stats: %v\n", subcontainer.Stats[0].Network.RxBytes)
		fmt.Printf("   Stats: %v\n", subcontainer.Stats[0].Network.Interfaces)
		fmt.Printf("   Stats: %v\n", subcontainer.Stats[0].Filesystem)
		fmt.Printf("   Stats: %v\n", subcontainer.Stats[0].Accelerators)
		fmt.Printf("   Stats: %v\n", subcontainer.Stats[0].Timestamp)
		fmt.Printf("   Stats: %v\n", subcontainer.Stats[0].Errors)
		fmt.Printf("   Stats: %v\n", subcontainer.Stats[0].Perf)
		fmt.Printf("   Stats: %v\n", subcontainer.Stats[0].Cpu.Usage.PerCpu)
		fmt.Printf("   Stats: %v\n", subcontainer.Stats[0].Cpu.Usage.User)
		fmt.Printf("   Stats: %v\n", subcontainer.Stats[0].Cpu.Usage.System)
		fmt.Printf("   Stats: %v\n", subcontainer.Stats[0].Cpu.Usage.PerCpu)
		fmt.Printf("   Stats: %v\n", subcontainer.Stats[0].Cpu.Usage.PerCpu)
		fmt.Printf("   Stats: %v\n", subcontainer.Stats[0].Cpu.Usage.PerCpu)
		fmt.Printf("   Stats: %v\n", subcontainer.Stats[0].Cpu.Usage.PerCpu)
		fmt.Printf("   Stats: %v\n", subcontainer.Stats[0].Cpu.Usage.PerCpu)

	}

}
