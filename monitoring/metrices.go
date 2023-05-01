package monitoring

// here we will get the info regarding the cluster and its components

import (
	"context"
	"log"
	"strconv"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/metrics/pkg/client/clientset/versioned"
)

// as we will have multiple nodes we will define the nodes as an array
// also we will define the cpu usage for the nodes as well memory usage for the nodes
type monitoring struct {
	clusterName string
	cpu         float64
	cores       int64
	nodes       []string
	memory      float64
	disk        float64
	totaldisk   float64
	billing     float64
}

type NodeInfo struct {
	Name   []string
	Memory []float64
	CPU    []float64
	Disk   []float64
}

var nodeInfoList []NodeInfo

// cpu usage for the nodes
func cpu() ([]NodeInfo, error) {
	// we will get the info regarding the cpu usage for the nodes
	// we will use kubernetes API for this

	// creat the kubernetes client
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

	// we will get the info regarding the cpu usage for the nodes
	// we will use the metrics API for this
	// create the metrics client

	metricsClient, err := versioned.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	// get the metrics for the nodes

	// get the list of nodes
	nodes, err := clientset.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	// get the metrics for the nodes
	for _, node := range nodes.Items {
		metrics, err := metricsClient.MetricsV1beta1().NodeMetricses().Get(context.Background(), node.Name, metav1.GetOptions{})
		if err != nil {
			log.Fatal(err)
		}

		// get the cpu usage for the nodes
		cpuUsage := metrics.Usage.Cpu().MilliValue()
		// get the cpu cores for the nodes
		cpuCores := node.Status.Capacity.Cpu().MilliValue()
		// get the cpu usage percentage for the nodes
		cpuUsagePercentage := float64(cpuUsage) / float64(cpuCores) * 100

		// get the memory usage for the nodes
		memoryUsage := metrics.Usage.Memory().Value()
		// get the memory capacity for the nodes
		memoryCapacity := node.Status.Capacity.Memory().Value()
		// get the memory usage percentage for the nodes
		memoryUsagePercentage := float64(memoryUsage) / float64(memoryCapacity) * 100

		// get the disk usage for the nodes
		diskUsage := metrics.Usage.StorageEphemeral().Value()
		// get the disk capacity for the nodes
		diskCapacity := node.Status.Capacity.StorageEphemeral().Value()
		// get the disk usage percentage for the nodes
		diskUsagePercentage := float64(diskUsage) / float64(diskCapacity) * 100

		// now append the data to the node struct

		nodeInfo := NodeInfo{
			Name:   []string{node.Name},
			CPU:    []float64{cpuUsagePercentage},
			Memory: []float64{memoryUsagePercentage},
			Disk:   []float64{diskUsagePercentage},
		}
		nodeInfoList = append(nodeInfoList, nodeInfo)
	}

	return nodeInfoList, nil

}

// ressources struct
type resources struct {
	clusterName            string
	cpu                    float64
	cores                  int64
	memory                 float64
	disk                   float64
	nodes                  []string
	pods                   []string
	services               []string
	ingresses              []string
	deployments            []string
	statefulsets           []string
	daemonsets             []string
	confimap               []string
	secret                 []string
	namespaces             []string
	persistentvolumes      []string
	persistentvolumeclaims []string
}

// now we will decalre an array of the struct which we created above
var resourcesList []resources

// we will get the info regarding the cluster and its components
func clusterInfo() ([]resources, error) {
	// we will get the info regarding the cluster and its components
	// we will use kubernetes API for this

	// creat the kubernetes client
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
	// get the list of nodes
	nod, err := clientset.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	// get the list of configmaps
	cm, err := clientset.CoreV1().ConfigMaps("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	// get the list of secrets
	sec, err := clientset.CoreV1().Secrets("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	// get the list of pods
	po, err := clientset.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	// get the list of services
	svc, err := clientset.CoreV1().Services("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	// get the list of deployments
	depl, err := clientset.AppsV1().Deployments("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	// get the list of statefulsets
	ss, err := clientset.AppsV1().StatefulSets("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	// get the list of daemonsets
	ds, err := clientset.AppsV1().DaemonSets("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	// get the list of ingresses
	ig, err := clientset.NetworkingV1().Ingresses("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	// get the list of persistentvolumes
	pv, err := clientset.CoreV1().PersistentVolumes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)

	}

	// get the list of persistentvolumeclaims
	pvc, err := clientset.CoreV1().PersistentVolumeClaims("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	// get the list of namespaces
	ns, err := clientset.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	// now we will append the data to the struct
	// we will get the info regarding the cluster and its components

	// get the cluster name
	clusterName := nod.Items[0].Labels["kubernetes.io/cluster-service"]

	// get the cpu usage for the cluster
	cpuUsage := nod.Items[0].Status.Allocatable.Cpu().MilliValue()
	// get the cpu cores for the cluster
	cpuCores := nod.Items[0].Status.Capacity.Cpu().MilliValue()
	// get the cpu usage percentage for the cluster
	cpuUsagePercentage := float64(cpuUsage) / float64(cpuCores) * 100

	// get the memory usage for the cluster
	memoryUsage := nod.Items[0].Status.Allocatable.Memory().Value()
	// get the memory capacity for the cluster
	memoryCapacity := nod.Items[0].Status.Capacity.Memory().Value()

	// now append the data to the node struct
	resourcesList = append(resourcesList, resources{
		clusterName:            clusterName,
		cpu:                    cpuUsagePercentage,
		cores:                  cpuCores,
		nodes:                  []string{strconv.Itoa(len(nod.Items))},
		memory:                 float64(memoryUsage) / float64(memoryCapacity) * 100,
		disk:                   float64(nod.Items[0].Status.Allocatable.StorageEphemeral().Value()) / float64(nod.Items[0].Status.Capacity.StorageEphemeral().Value()) * 100,
		pods:                   []string{strconv.Itoa(len(po.Items))},
		services:               []string{strconv.Itoa(len(svc.Items))},
		ingresses:              []string{strconv.Itoa(len(ig.Items))},
		deployments:            []string{strconv.Itoa(len(depl.Items))},
		statefulsets:           []string{strconv.Itoa(len(ss.Items))},
		daemonsets:             []string{strconv.Itoa(len(ds.Items))},
		confimap:               []string{strconv.Itoa(len(cm.Items))},
		secret:                 []string{strconv.Itoa(len(sec.Items))},
		namespaces:             []string{strconv.Itoa(len(ns.Items))},
		persistentvolumeclaims: []string{strconv.Itoa(len(pvc.Items))},
		persistentvolumes:      []string{strconv.Itoa(len(pv.Items))},
	})

	return resourcesList, nil
}
