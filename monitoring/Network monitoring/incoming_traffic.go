//  In addition to CPU and memory utilization, it's important to monitor network traffic and latency within the Kubernetes cluster

// getting the incomimg traffic

package Networkmonitoring

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	// how we can get the total incoming traffic
	// expose the total number of incoming traffic as an API endpoint. we can use  web framework of
	// choice to create a new HTTP endpoint that returns the total number of incoming bytes as a JSON response.

	// get the endpoints of the pods then get get then create a new HTTP endpoint that
	//  returns the total number of incoming bytes as a JSON response.

	// we already have endpoints of the pods and all the cluster
	// import the getEndpoints() function from the endpoints.go file

	slice, totalnumofendpoints := getEndpoints()
	fmt.Println(totalnumofendpoints)

}

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
