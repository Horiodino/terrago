package Containermonitoring

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// defining a strunct to store the container metrics and all
type containerMetrics struct {
	// conatainer name of type var containerNames
	cname       []string
	podName     []string
	nsname      []string
	node        []string
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

var AllContainerMetrics int32

func getContaienMetrices() {

	// first of all we will egt the container names

}

type containerNames struct {
	containerName string
	podName       string
	nsName        string
	nodeName      string
}

var containerinfo []containerNames

func getContainerNames() {
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatal(err)
	}

	// create the clientset
	// here its taking the config which we created above
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	pods, err := clientset.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, pod := range pods.Items {
		for _, container := range pod.Spec.Containers {
			// now we will get the container name and store it in a slice of struct
			containerinfo = append(containerinfo, containerNames{
				containerName: container.Name,
				podName:       pod.Name,
				nsName:        pod.Namespace,
				nodeName:      pod.Spec.NodeName,
			})
		}
	}
}

// now save the reterived data in a struct and also save in the database
// we will use the struct to save the data in the database

func savecontainerData() {
	// now we will save the data in the database
	// mongodb client
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	// creating a context also craete db if not exist
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
