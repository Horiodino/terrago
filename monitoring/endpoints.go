package monitoring

import (
	"context"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// creating the endpoints slice
var endpointsSlice []string

// getEndpoints will get the endpoints of the cluster using the kubernetes client
func getEndpoints() {

	// create the kubernetes client
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	// get the endpoints of the cluster
	endpoints, err := clientset.CoreV1().Endpoints("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	// append the endpoints to the endpoints slice
	for _, endpoint := range endpoints.Items {
		endpointsSlice = append(endpointsSlice, endpoint.Name)
	}

}

// getHTTPRequests will get the number of HTTP requests using the endpoints
func getHTTPRequests() int {

	// create the Kubernetes client
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	// use the EndpointSlice API to get the number of HTTP requests
	endpointSlices, err := clientset.DiscoveryV1beta1().EndpointSlices("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	var numHTTPRequests int

	// iterate over each EndpointSlice and count the number of HTTP requests
	for _, endpointSlice := range endpointSlices.Items {
		for _, endpoint := range endpointSlice.Endpoints {
			for _, subset := range endpoint.Subsets {
				for _, address := range subset.Addresses {
					for _, port := range subset.Ports {
						if port.Protocol == "TCP" && port.Port == 80 {
							numHTTPRequests += len(address.TargetRef)
						}
					}
				}
			}
		}
	}

	return numHTTPRequests
}
