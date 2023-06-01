package monitoring

// here we will get the info regarding the cluster and its components

import (
	"context"

	// Kubernetes API client libraries and packages
	".mongodb.org/mongo-driver/mongo"go
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

// aglobal struct for authorization client for kubernetes client so that DRY principle can be followed

type client struct {
	clientset *kubernetes.Clientset
}

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

var m *monitoring

func getinfo() {

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

type resources struct {
	pods                   string
	services               string
	ingresses              string
	deployments            string
	statefulsets           string
	daemonsets             string
	confimap               string
	secret                 string
	namespaces             string
	persistentvolumes      string
	persistentvolumeclaims string
}

// now we will decalre an array of the struct which we created above
var resourcesList []resources

// we will get the info regarding the cluster and its components
func clusterInfo() ([]resources, error) {

	config, err := clientcmd.BuildConfigFromFlags("", os.Getenv("HOME")+"/.kube/config")
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	cm, err := clientset.CoreV1().ConfigMaps("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	sec, err := clientset.CoreV1().Secrets("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	po, err := clientset.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	svc, err := clientset.CoreV1().Services("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	depl, err := clientset.AppsV1().Deployments("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	ss, err := clientset.AppsV1().StatefulSets("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	ds, err := clientset.AppsV1().DaemonSets("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	ig, err := clientset.NetworkingV1().Ingresses("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	pv, err := clientset.CoreV1().PersistentVolumes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)

	}

	pvc, err := clientset.CoreV1().PersistentVolumeClaims("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	ns, err := clientset.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	resourcesList = append(resourcesList, resources{

		pods:                   strconv.Itoa(len(po.Items)),
		services:               strconv.Itoa(len(svc.Items)),
		ingresses:              strconv.Itoa(len(ig.Items)),
		deployments:            strconv.Itoa(len(depl.Items)),
		statefulsets:           strconv.Itoa(len(ss.Items)),
		daemonsets:             strconv.Itoa(len(ds.Items)),
		confimap:               strconv.Itoa(len(cm.Items)),
		secret:                 strconv.Itoa(len(sec.Items)),
		namespaces:             strconv.Itoa(len(ns.Items)),
		persistentvolumes:      strconv.Itoa(len(pv.Items)),
		persistentvolumeclaims: strconv.Itoa(len(pvc.Items)),
	})

	for _, resource := range resourcesList {
		fmt.Println("Pods: ", resource.pods)
		fmt.Println("Services: ", resource.services)
		fmt.Println("Ingresses: ", resource.ingresses)
		fmt.Println("Deployments: ", resource.deployments)
		fmt.Println("Statefulsets: ", resource.statefulsets)
		fmt.Println("Daemonsets: ", resource.daemonsets)
		fmt.Println("Configmaps: ", resource.confimap)
		fmt.Println("Secrets: ", resource.secret)
		fmt.Println("Namespaces: ", resource.namespaces)
		fmt.Println("Persistent Volumes: ", resource.persistentvolumes)
		fmt.Println("Persistent Volume Claims: ", resource.persistentvolumeclaims)
	}

	return resourcesList, nil
}
//works till here
type namespacestruct struct {
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
	nsjobs                   []string
}

var namespacesList []namespacestruct

func namespacesInfo() ([]namespacestruct, error) {
	// we will get the info for a particular namespace like
	config, err := clientcmd.BuildConfigFromFlags("", os.Getenv("HOME")+"/.kube/config")
	if err != nil {
		log.Fatal(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	namespaces, err := clientset.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	i := 1

	for _, namespace := range namespaces.Items {

		nsname := namespace.Name
		//gett all the info regarding the namespcace
		pods, err := clientset.CoreV1().Pods(namespace.Name).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}
		////// PENDING GET ALL THE RESOURCES NAMES
		services, err := clientset.CoreV1().Services(namespace.Name).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}

		ingresses, err := clientset.NetworkingV1().Ingresses(namespace.Name).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}

		deployments, err := clientset.AppsV1().Deployments(namespace.Name).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}

		statefulsets, err := clientset.AppsV1().StatefulSets(namespace.Name).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}

		daemonsets, err := clientset.AppsV1().DaemonSets(namespace.Name).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}

		confimap, err := clientset.CoreV1().ConfigMaps(namespace.Name).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}

		secret, err := clientset.CoreV1().Secrets(namespace.Name).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}

		persistentvolumeclaims, err := clientset.CoreV1().PersistentVolumeClaims(namespace.Name).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}
		persistentvolumes, err := clientset.CoreV1().PersistentVolumes().List(context.Background(), metav1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("working", i)
		jobs, err := clientset.BatchV1().Jobs(namespace.Name).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}

		namespaceInfo := namespacestruct{
			namespaceName:            nsname,
			nspods:                   []string{strconv.Itoa(len(pods.Items))},
			nsservices:               []string{strconv.Itoa(len(services.Items))},
			nsingresses:              []string{strconv.Itoa(len(ingresses.Items))},
			nsdeployments:            []string{strconv.Itoa(len(deployments.Items))},
			nsstatefulsets:           []string{strconv.Itoa(len(statefulsets.Items))},
			nsdaemonsets:             []string{strconv.Itoa(len(daemonsets.Items))},
			nsconfimap:               []string{strconv.Itoa(len(confimap.Items))},
			nssecret:                 []string{strconv.Itoa(len(secret.Items))},
			nspersistentvolumeclaims: []string{strconv.Itoa(len(persistentvolumeclaims.Items))},
			nspersistentvolumes:      []string{strconv.Itoa(len(persistentvolumes.Items))},
			nsjobs:                   []string{strconv.Itoa(len(jobs.Items))},
		}

		namespacesList = append(namespacesList, namespaceInfo)
		for _, namespace := range namespacesList {
			fmt.Println("┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
			fmt.Println("┃ Namespace Name: ", namespace.namespaceName)
			fmt.Println("┃ Pods: ", namespace.nspods)
			fmt.Println("┃ Services: ", namespace.nsservices)
			fmt.Println("┃ Ingresses: ", namespace.nsingresses)
			fmt.Println("┃ Deployments: ", namespace.nsdeployments)
			fmt.Println("┃ Statefulsets: ", namespace.nsstatefulsets)
			fmt.Println("┃ Daemonsets: ", namespace.nsdaemonsets)
			fmt.Println("┃ Configmaps: ", namespace.nsconfimap)
			fmt.Println("┃ Secrets: ", namespace.nssecret)
			fmt.Println("┃ Persistent Volume Claims: ", namespace.nspersistentvolumeclaims)
			fmt.Println("┃ Persistent Volumes: ", namespace.nspersistentvolumes)
			fmt.Println("┃ Jobs: ", namespace.nsjobs)
			fmt.Println("┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
			fmt.Println("")
			fmt.Println("")

		}

	}

	return namespacesList, nil
}

