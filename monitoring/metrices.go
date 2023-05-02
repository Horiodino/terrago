package monitoring

// here we will get the info regarding the cluster and its components

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/metrics/pkg/client/clientset/versioned"

	// using mongoDB

	"log"
	"strconv"
	"time"
)

----------------------------------------------------------------------------------------------------------------------------
// as we will have multiple nodes we will define the nodes as an array
// also we will define the cpu usage for the nodes as well memory usage for the nodes
type monitoring struct {
	clusterName string
	cpu         float64
	cores       int64
	nodes       int64
	totalmemory float64
	usedmemory  float64
	disk        float64
	totaldisk   float64
	billing     float64
}

// this monotoring struct will get the info for alerting rules
func (m *monitoring) getinfo() {

	// cpu usage for the cluster
	// kubernetes client
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatal(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	// get the metrics for the cluster
	metricsClient, err := versioned.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	metrics, err := metricsClient.MetricsV1beta1().NodeMetricses().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	var cpuUsage int64
	for _, node := range metrics.Items {
		cpuUsage += node.Usage.Cpu().MilliValue()
	}
	nodes, err := clientset.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	var cpuCores int64
	for _, node := range nodes.Items {
		cpuCores += node.Status.Capacity.Cpu().Value()
	}
	var memoryUsage int64
	for _, node := range metrics.Items {
		memoryUsage += node.Usage.Memory().Value()
	}
	var memory int64
	for _, node := range nodes.Items {
		memory += node.Status.Capacity.Memory().Value()
	}
	var diskUsage int64
	for _, node := range metrics.Items {
		diskUsage += node.Usage.StorageEphemeral().Value()
	}
	var disk int64
	for _, node := range nodes.Items {
		disk += node.Status.Capacity.StorageEphemeral().Value()
	}

	// now we have all the info regarding the cpu usage, memory usage, disk usage, etc of the entire cluster
	// now append the info to the struct
	m.cpu = float64(cpuUsage)
	m.cores = cpuCores
	m.nodes = int64(len(nodes.Items))
	m.totalmemory = float64(memory)
	m.usedmemory = float64(memoryUsage)
	m.disk = float64(diskUsage)
	m.totaldisk = float64(disk)

}
// now we will save the info that we saved in the struct to the database
func (m *monitoring) savedata() {
}
----------------------------------------------------------------------------------------------------------------------------

// dont be confused ,   this struct is for the nodes info not for the entire cluster
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
----------------------------------------------------------------------------------------------------------------------------

// ressources struct for entire cluster
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

// as we are using the kubernetes API to get the info regarding the cluster and its components
// to use mongodb we will create a struct which will have the same fields as the struct which we created above
// we will use this struct to insert the data into the mongodb
// not surly if this is the best way  but working to saving in json format
type clusterInfoMongo struct {
	clusterName            string
	cpu                    float64
	cores                  int64
	nodes                  []string
	memory                 float64
	disk                   float64
	pods                   []string
	services               []string
	ingresses              []string
	deployments            []string
	statefulsets           []string
	daemonsets             []string
	confimap               []string
	secret                 []string
	namespaces             []string
	persistentvolumeclaims []string
	persistentvolumes      []string
}

// this function will insert the data into the mongodb
func insertDataMongo(resourcesList []resources) error {

	// create a new mongo client
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	// create a new context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// connect to the mongo client
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// create a new database
	database := client.Database("clusterInfo")

	// create a new collection
	collection := database.Collection("clusterInfo")

	// now we will insert the data into the collection
	for _, v := range resourcesList {

		// create a new struct
		clusterInfoMongo := clusterInfoMongo{
			clusterName:            v.clusterName,
			cpu:                    v.cpu,
			cores:                  v.cores,
			nodes:                  v.nodes,
			memory:                 v.memory,
			disk:                   v.disk,
			pods:                   v.pods,
			services:               v.services,
			ingresses:              v.ingresses,
			deployments:            v.deployments,
			statefulsets:           v.statefulsets,
			daemonsets:             v.daemonsets,
			confimap:               v.confimap,
			secret:                 v.secret,
			namespaces:             v.namespaces,
			persistentvolumeclaims: v.persistentvolumeclaims,
			persistentvolumes:      v.persistentvolumes,
		}

		// insert the data into the collection
		_, err := collection.InsertOne(ctx, clusterInfoMongo)
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}
----------------------------------------------------------------------------------------------------------------------------

// before that lets get the namespace name from the resources struct
// namepace struct
type namespace struct {
	namespaceName            string
	nspods                   []string
	nsservices               []string
	nsingresses              []string
	nsdeployments            []string
	nsstatefulsets           []string
	nsdaemonsets             []string
	nsconfimap               []string
	nssecret                 []string
	nspersistentvolumeclaims []string
	nspersistentvolumes      []string
}

// now we will get the info regarding the namespces
func namespacesInfo(r *resources, ns []namespace) ([]namespace, error) {
	namespaces := r.namespaces
	for i := 0; i < len(namespaces); i++ {
		config, err := rest.InClusterConfig()
		if err != nil {
			log.Fatal(err)
		}
		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			log.Fatal(err)
		}
		po, err := clientset.CoreV1().Pods(namespaces[i]).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}
		svc, err := clientset.CoreV1().Services(namespaces[i]).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}
		ig, err := clientset.NetworkingV1().Ingresses(namespaces[i]).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}
		depl, err := clientset.AppsV1().Deployments(namespaces[i]).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}
		ss, err := clientset.AppsV1().StatefulSets(namespaces[i]).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}
		ds, err := clientset.AppsV1().DaemonSets(namespaces[i]).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}
		cm, err := clientset.CoreV1().ConfigMaps(namespaces[i]).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}
		sec, err := clientset.CoreV1().Secrets(namespaces[i]).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}
		pvc, err := clientset.CoreV1().PersistentVolumeClaims(namespaces[i]).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}
		pv, err := clientset.CoreV1().PersistentVolumes().List(context.Background(), metav1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}

		// Create a new namespace struct
		newNS := namespace{namespaceName: namespaces[i]}

		// Append relevant data to the struct fields
		for _, pod := range po.Items {
			newNS.nspods = append(newNS.nspods, pod.Name)
		}
		for _, service := range svc.Items {
			newNS.nsservices = append(newNS.nsservices, service.Name)
		}
		for _, ingress := range ig.Items {
			newNS.nsingresses = append(newNS.nsingresses, ingress.Name)
		}
		for _, deployment := range depl.Items {
			newNS.nsdeployments = append(newNS.nsdeployments, deployment.Name)
		}
		for _, statefulset := range ss.Items {
			newNS.nsstatefulsets = append(newNS.nsstatefulsets, statefulset.Name)
		}
		for _, daemonset := range ds.Items {
			newNS.nsdaemonsets = append(newNS.nsdaemonsets, daemonset.Name)
		}
		for _, confimap := range cm.Items {
			newNS.nsconfimap = append(newNS.nsconfimap, confimap.Name)
		}
		for _, secret := range sec.Items {
			newNS.nssecret = append(newNS.nssecret, secret.Name)
		}
		for _, persistentvolumeclaim := range pvc.Items {
			newNS.nspersistentvolumeclaims = append(newNS.nspersistentvolumeclaims, persistentvolumeclaim.Name)
		}
		for _, persistentvolume := range pv.Items {
			newNS.nspersistentvolumes = append(newNS.nspersistentvolumes, persistentvolume.Name)
		}

		// Append the new namespace struct to the namespace slice
		ns = append(ns, newNS)
	}
	return ns, nil
}
----------------------------------------------------------------------------------------------------------------------------
