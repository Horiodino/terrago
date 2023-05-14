package monitoring

import (
	"context"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// getEndpoints will get the endpoints of the cluster using the kubernetes client
func getEndpoints() ([]string, []string, int) {

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

	var correspondingpods []string

	// new addition here we will also save the pods behind the endpoints and also return it !!

	for _, endpoint := range endpoints.Items {
		for _, subset := range endpoint.Subsets {
			for _, address := range subset.Addresses {
				// get the pod name from the endpoint address
				podName := address.TargetRef.Name
				// append the pod name to the corresponding pods slice
				correspondingpods = append(correspondingpods, podName)
			}
		}
	}
	// creating the endpoints slice
	var endpointsSlice []string

	// append the endpoints to the endpoints slice
	for _, endpoint := range endpoints.Items {
		endpointsSlice = append(endpointsSlice, endpoint.Name)
	}

	totalnumofendpoints := len(endpointsSlice)

	// print while returning the endpoints slice
	return endpointsSlice, correspondingpods, totalnumofendpoints
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
	// get the nunber of HTTP requests
	for _, endpointSlice := range endpointSlices.Items {
		numHTTPRequests += len(endpointSlice.Ports)
	}

	return numHTTPRequests
}

// Application specific requests
// here we are going to get the application specific requests from the endpoints
// so we can monitor the application specific requests
