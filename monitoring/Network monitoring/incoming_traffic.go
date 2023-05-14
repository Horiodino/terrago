//  In addition to CPU and memory utilization, it's important to monitor network traffic and latency within the Kubernetes cluster

// getting the incomimg traffic

package Networkmonitoring

import (
	"fmt"

	"k8s.io/client-go/kubernetes"
	// "C:\Users\Holiodin\webgo\monitoring\Containermonitoring\metrices.go"
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

		Iterate over each container in the list and get the corresponding list of interfaces in that container.
		Iterate over each interface in the list and get the corresponding list of incoming traffic in that interface.
		Add up the total incoming traffic for each interface to get the total incoming traffic for that container.
		Add up the total incoming traffic for each container to get the total incoming traffic for that pod.
		Add up the total incoming traffic for each pod to get the total incoming traffic for that endpoint.
		Add up the total incoming traffic for all the endpoints to get the total incoming traffic for the entire cluster.
	*/

	// --------------------------------------------------------------------
	//Get the list of all the endpoints in the cluster using the Kubernetes client API.
	//Iterate over each endpoint in the list and get the corresponding list of pods behind that endpoint
	slice, podinfo, totalnumofendpoints := getEndpoints()
	fmt.Println(totalnumofendpoints)

	// Iterate over each pod in the list and get the corresponding list of containers in that pod.
	for i := 0; i < len(podinfo); i++ {
		podname := podinfo[i].Name
		podnamespace := podinfo[i].Namespace
		fmt.Fprintf(w, "Pod Name: %s\n", podname, podnamespace)
		fmt.Println("---------------------------")

		// for lopping over the containers in the pod
		for j := 0; j < len(podinfo[i].Containers); j++ {
			containername := podinfo[i].Containers[j].Name
			fmt.Fprintf(w, "Container Name: %s\n", containername)
			fmt.Println("---------------------------")

			// getting the list 
	}

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
