// here we will get the ednpoints to get the more info aboute the cluster and its components

package monitoring

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"log"
)

// here we will get the ednpoints to get the more info aboute the cluster and its components
// we will use the kubernetes client to get the endpoints
// what are the details we will get from the endpoints
// we will get the info regarding the nodes, pods, services, deployments, etc
// we will get the info regarding the cpu usage, memory usage, disk usage, etc
// we will get the info regarding the billing as well
// we will get the info regarding the network usage as well

func getEndpoints() {

	// get the endpoints of the cluster using the kubernetes client

	//create the kubernetes client
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

	// get the endpoints of the cluster

}

func getMetrics() {

}
