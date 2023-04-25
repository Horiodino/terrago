package monitoring

// here we will get the info regarding the cluster and its components

import (
	"context"
	"fmt"
	"time"

	"k8s.io/client-go/informers/core"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/metrics/pkg/client/clientset/versioned"
)

// as we will have multiple nodes we will define the nodes as an array
// also we will define the cpu usage for the nodes as well memory usage for the nodes
type monitoring struct {
	clusterName string
	cpu         float64
	cores        int64
	nodes       []string
	memory      float64
	disk        float64
	totaldisk   float64
	billing     float64
}

type NodeInfo struct {
    Name         []string
    Memory       []int
    CPU          []int
    CPUUses      []int
}


// we will get the info regarding the cluster and its components
// cpu usage for the nodes
func cpu() (float64, error) {
	// we will get the info regarding the cpu usage for the nodes
	// we will use kubernetes API for this

	// creat the kubernetes client
	config, err := rest.InClusterConfig()
	if err != nil {
		return 0, err
	}

	// create the clientset
	// here its taking the config which we created above
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return 0, err
	}

	// we will get the info regarding the cpu usage for the nodes
	// we will use the metrics API for this
	// create the metrics client

	metricsClient, err := versioned.NewForConfig(config)
	if err != nil {
		return 0, err
	}

	// get the metrics for the nodes

	// get the list of nodes
	nodes, err := clientset.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return 0, err
	}

	// get the metrics for the nodes
	for _, node := range nodes.Items {
		metrics, err := metricsClient.MetricsV1beta1().NodeMetricses().Get(context.Background(), node.Name, metav1.GetOptions{})
		if err != nil {
			return 0, err
		}

		// get the cpu usage for the nodes
		cpuUsage := metrics.Usage.Cpu().MilliValue()
		// appending the node info struct 

		
		

	}
