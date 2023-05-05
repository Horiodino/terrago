package Containermonitoring

import "k8s.io/client-go/kubernetes"

// defining a strunct to store the container metrics and all
type containerMetrics struct {
	// conatainer name of type var containerNames
	cname       []string
	cpuUsage    []int
	memoryUsage []int
	diskIo      []int
	networkTx   []int
	networkRx   []int
}

// getting all the container names
// and storing them in a slice
// slice declaration
// var containerNames []string
// var containerN string

//here are the thing that we need to do
// Container name
// CPU usage
// Memory usage
// Filesystem usage
// Network usage
// Container status (e.g., running, exited, etc.)
// Container ID
// Container image
// Container creation and start time
// Container labels
// Container environment variables
// Container command and arguments

func getContaienMetrices() {
	// we will sue kubelet api to get the container metrics as the kubelet api is the best way to get the container metrics

	//kuebrnetes api to get the container metrics
	// craete the kuerbetes client
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// now get the container names using the kubernetes api

}
