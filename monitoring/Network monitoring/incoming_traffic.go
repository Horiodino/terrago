//  In addition to CPU and memory utilization, it's important to monitor network traffic and latency within the Kubernetes cluster

// getting the incomimg traffic

package Networkmonitoring

import (
	"fmt"

	"k8s.io/client-go/kubernetes"
)

// how we can get the total incoming traffic
// simply get the total number of bytes received on all interfaces by the pod

func incomingtraffic() {
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// using network interface to get the incoming traffic
	// this for loop will iterate over all the pods and get the incoming traffic
	totalincomingtraffic := totalincomingtraffic()
	fmt.Println(totalincomingtraffic)
}
func totalincomingtraffic() {
	/*
		Iterate over each endpoint in the list and get the corresponding list of pods behind that endpoint.
		Iterate over each pod in the list and get the corresponding list of containers in that pod.
		Iterate over each container in the list and get the corresponding list of interfaces in that container.
		Iterate over each interface in the list and get the corresponding list of incoming traffic in that interface.
		Add up the total incoming traffic for each interface to get the total incoming traffic for that container.
		Add up the total incoming traffic for each container to get the total incoming traffic for that pod.
		Add up the total incoming traffic for each pod to get the total incoming traffic for that endpoint.
		Add up the total incoming traffic for all the endpoints to get the total incoming traffic for the entire cluster.
	*/

	// --------------------------------------------------------------------
	//Get the list of all the endpoints in the cluster using the Kubernetes client API.
	slice, totalnumofendpoints := getEndpoints()
	fmt.Println(totalnumofendpoints)

	//Iterate over each endpoint in the list and get the corresponding list of pods behind that endpoint
	for _, endpoint := range slice {
		//Iterate over each pod in the list and get the corresponding list of containers in that pod.
		podlist := getPods(endpoint)
	}
}

func getEndpoints() ([]string, int) {
	//Get the list of all the endpoints in the cluster using the Kubernetes client API.
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	//Get the liat of pods corresponding to the endpoint

	return nil, 0
}

// ------------------------------------------------------------------------------------------------------------------

func outgoingtraffic() {

}

func networktraffic() {

}

func networklatency() {

}

func networkerrors() {

}

func networkrequests() {

}

func exposedports() {

}

// getting the network throughput

// getting the network packets

// getting the network connections

// getting the network connections

// getting the exposed services

// getting the exposed endpoints

// getting the exposed routes
