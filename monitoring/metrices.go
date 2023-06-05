package monitoring

// here we will get the info regarding the cluster and its components

import (
	"context"
	"fmt"
	"os"

	// Kubernetes API client libraries and packages
	// ".mongodb.org/mongo-driver/mongo"go

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/metrics/pkg/client/clientset/versioned"

	// using mongoDB

	"log"
	"strconv"
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

func Getinfo() {
	m = &monitoring{}

	// cpu usage for the cluster
	// kubernetes client
	config, err := clientcmd.BuildConfigFromFlags("", os.Getenv("HOME")+"/.kube/config")
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

func Display() {
	// display the stored info of Getinfo function that is stored in the struct
	fmt.Println("CPU Usage: ", m.cpu)
	fmt.Println("CPU Cores: ", m.cores)
	fmt.Println("Nodes: ", m.nodes)
	fmt.Println("Total Memory: ", m.totalmemory)
	fmt.Println("Used Memory: ", m.usedmemory)
	fmt.Println("Disk Usage: ", m.disk)
	fmt.Println("Total Disk: ", m.totaldisk)

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

// works till here
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

type namespaceInfoDetailed struct {
	namespaceName            string
	nspods                   []nspods
	nsservices               []nsservices
	nsingresses              []nsingresses
	nsdeployments            []nsdeployments
	nsstatefulsets           []nsstatefulsets
	nsdaemonsets             []nsdaemonsets
	nsconfimap               []nsconfimap
	nssecret                 []nssecret
	nspersistentvolumeclaims []nspersistentvolumeclaims
	// nspersistentvolumes      []string
	// nsjobs                   [][]string
}

// key value pairs pods for resoueces and limits
type nspods struct {
	podName  string
	poStatus string
	podCPU   string
	podMEM   string
	poImages []string
}

type nsservices struct {
	serviceName       string
	serviceType       string
	serviceTarget     string
	serviceTargetname string
}

// key value for ingresses
type nsingresses struct {
	ingressName       string
	ingressHost       string
	ingressPath       string
	ingressTarget     string
	ingressTargetPort string
	ingressesLBtype   string
}

type nsdeployments struct {
	deploymentName     string
	deploymentReplicas string
	deploymentStatus   string
	deploymentStrategy string
	deploymentImages   []string
}

type nsstatefulsets struct {
	statefulsetName     string
	statefulsetReplicas string
	statefulsetStatus   string
	statefulsetStrategy string
	statefulsetImages   []string
}

type nsdaemonsets struct {
	// what does deamonset have?
	daemonsetName              string
	daemonsetNamespace         string
	daemonsetLabels            map[string]string
	daemonsetMatchlables       map[string]string
	daemonsetContainers        []string
	daemonsetTerminationperiod string
	daemonsetVolumes           []string
	daemonsetVolumemount       []string
	daemonsetStatus            string
}

type nspersistentvolumeclaims struct {
	pvc_Name      string
	pvc_Labels    map[string]string
	pvc_tatus     string
	pvc_Volume    string
	pvc_Size      string
	pvc_AcessMode string
	// pvc_StorageClass string
}

type nsconfimap struct {
	confimapName string
}

type nssecret struct {
	secretName string
	// isusedBy   string
}

var namespacesListDetailed []namespaceInfoDetailed

func getnamespaceInfoDetailed() {
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

	for _, namespace := range namespaces.Items {
		nsname := namespace.Name
		pods, err := clientset.CoreV1().Pods(namespace.Name).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}
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

		var podList []nspods
		var serviceList []nsservices
		var ingressList []nsingresses
		var deploymentList []nsdeployments
		var statefulsetList []nsstatefulsets
		var daemonsetList []nsdaemonsets
		var confimapList []nsconfimap
		var secretList []nssecret
		var persistentvolumeclaimList []nspersistentvolumeclaims

		for _, pod := range pods.Items {
			podName := pod.Name

			poStatus := pod.Status

			podCPU := pod.Spec.Containers[0].Resources.Requests.Cpu().String()

			podMEM := pod.Spec.Containers[0].Resources.Requests.Memory().String()

			var podImages []string
			for _, container := range pod.Spec.Containers {
				podImages = append(podImages, container.Image)
			}
			podInfo := nspods{
				podName:  podName,
				poStatus: poStatus.String(),
				podCPU:   podCPU,
				podMEM:   podMEM,
				poImages: podImages,
			}

			podList = append(podList, podInfo)
		}

		for _, service := range services.Items {
			svcName := service.Name
			svcType := service.Spec.Type
			svcTarget := service.Spec.Ports[0].TargetPort.String()
			svcTargetname := service.Spec.Ports[0].Name

			serviceInfo := nsservices{
				serviceName:       svcName,
				serviceType:       string(svcType),
				serviceTarget:     svcTarget,
				serviceTargetname: svcTargetname,
			}

			serviceList = append(serviceList, serviceInfo)

		}

		for _, ingresses := range ingresses.Items {
			ingname := ingresses.Name
			inghost := ingresses.Spec.Rules[0].Host
			ingpath := ingresses.Spec.Rules[0].HTTP.Paths[0].Path
			ingtarget := ingresses.Spec.Rules[0].HTTP.Paths[0].Backend.Service.Name
			ingtargetport := ingresses.Spec.Rules[0].HTTP.Paths[0].Backend.Service.Port.Number
			inglbtype := ingresses.Spec.Rules[0].IngressRuleValue.HTTP.Paths[0].Backend.Service.Port.Name

			ingressInfo := nsingresses{
				ingressName:       ingname,
				ingressHost:       inghost,
				ingressPath:       ingpath,
				ingressTarget:     ingtarget,
				ingressTargetPort: string(ingtargetport),
				ingressesLBtype:   inglbtype,
			}
			ingressList = append(ingressList, ingressInfo)
		}

		for _, deployment := range deployments.Items {
			depname := deployment.Name
			depreplicas := deployment.Spec.Replicas
			depstatus := deployment.Status
			depstrategy := deployment.Spec.Strategy
			var depimages []string
			for _, container := range deployment.Spec.Template.Spec.Containers {
				depimages = append(depimages, container.Image)
			}

			deploymentInfo := nsdeployments{
				deploymentName:     depname,
				deploymentReplicas: string(*depreplicas),
				deploymentStatus:   depstatus.String(),
				deploymentStrategy: depstrategy.String(),
				deploymentImages:   depimages,
			}

			deploymentList = append(deploymentList, deploymentInfo)
		}

		for _, statefulset := range statefulsets.Items {
			stsname := statefulset.Name
			stsreplicas := statefulset.Spec.Replicas
			stsstatus := statefulset.Status
			stsstrategy := statefulset.Spec
			var stsimages []string
			for _, container := range statefulset.Spec.Template.Spec.Containers {
				stsimages = append(stsimages, container.Image)
			}

			statefulsetInfo := nsstatefulsets{
				statefulsetName:     stsname,
				statefulsetReplicas: string(*stsreplicas),
				statefulsetStatus:   stsstatus.String(),
				statefulsetStrategy: stsstrategy.String(),
				statefulsetImages:   stsimages,
			}

			statefulsetList = append(statefulsetList, statefulsetInfo)
		}

		for _, daemonset := range daemonsets.Items {
			daeName := daemonset.Name
			daeNamespace := daemonset.Namespace
			daeLabels := daemonset.Labels
			daeMatchlables := daemonset.Spec.Selector.MatchLabels
			var daeContainers []string
			for _, container := range daemonset.Spec.Template.Spec.Containers {
				daeContainers = append(daeContainers, container.Image)
			}
			daeTerminationperiod := daemonset.Spec.Template.Spec.TerminationGracePeriodSeconds
			var daeVolumes []string
			for _, volume := range daemonset.Spec.Template.Spec.Volumes {
				daeVolumes = append(daeVolumes, volume.Name)
			}
			var daeVolumemount []string
			for _, volumemount := range daemonset.Spec.Template.Spec.Containers[0].VolumeMounts {
				daeVolumemount = append(daeVolumemount, volumemount.Name)
			}
			daeStatus := daemonset.Status

			daemonsetInfo := nsdaemonsets{
				daemonsetName:              daeName,
				daemonsetNamespace:         daeNamespace,
				daemonsetLabels:            daeLabels,
				daemonsetMatchlables:       daeMatchlables,
				daemonsetContainers:        daeContainers,
				daemonsetTerminationperiod: strconv.Itoa(int(*daeTerminationperiod)),
				daemonsetVolumes:           daeVolumes,
				daemonsetVolumemount:       daeVolumemount,
				daemonsetStatus:            daeStatus.String(),
			}

			daemonsetList = append(daemonsetList, daemonsetInfo)
		}
		for _, confimap := range confimap.Items {
			confimapInfo := nsconfimap{
				confimapName: confimap.Name,
			}
			confimapList = append(confimapList, confimapInfo)

		}
		for _, secret := range secret.Items {
			secretInfo := nssecret{
				secretName: secret.Name,
				// isusedBy:   ,
			}
			secretList = append(secretList, secretInfo)
		}

		for _, pvc := range persistentvolumeclaims.Items {
			pvcname := pvc.Name
			pvclabels := pvc.Labels
			pvcstatus := pvc.Status
			pvcvolume := pvc.Spec.VolumeName
			pvcsize := pvc.Spec.Resources.Requests.Storage().String()
			pvcaccessmode := pvc.Spec.AccessModes[0]
			// pvcstorageclass := pvc.Spec.StorageClassName

			pvcInfo := nspersistentvolumeclaims{
				pvc_Name:      pvcname,
				pvc_Labels:    pvclabels,
				pvc_tatus:     pvcstatus.String(),
				pvc_Volume:    pvcvolume,
				pvc_Size:      pvcsize,
				pvc_AcessMode: string(pvcaccessmode),
				// pvc_StorageClass: pvcstorageclass,
			}

			persistentvolumeclaimList = append(persistentvolumeclaimList, pvcInfo)

		}

		namespaceInfoDetailed := namespaceInfoDetailed{
			namespaceName:            nsname,
			nspods:                   podList,
			nsservices:               serviceList,
			nsingresses:              ingressList,
			nsdeployments:            deploymentList,
			nsstatefulsets:           statefulsetList,
			nsdaemonsets:             daemonsetList,
			nsconfimap:               confimapList,
			nssecret:                 secretList,
			nspersistentvolumeclaims: persistentvolumeclaimList,
		}

		namespacesListDetailed = append(namespacesListDetailed, namespaceInfoDetailed)

	}

	for _, info := range namespacesListDetailed {
		fmt.Println("┏━━" + info.namespaceName + "━━━")
		// fmt.Println("┃")
		// for _, pod := range info.nspods {
		// 	fmt.Println("┃ Pod Name: ", pod.podName)
		// 	fmt.Println("┃ Pod Status: ", pod.poStatus)
		// 	fmt.Println("┃ Pod CPU: ", pod.podCPU)
		// 	fmt.Println("┃ Pod Memory: ", pod.podMEM)
		// 	for _, image := range pod.poImages {
		// 		fmt.Println(" Pod Image: ", image)
		// 	}
		// 	fmt.Println("┃")
		// 	fmt.Println("┃")
		// }
		// for _, service := range info.nsservices {
		// 	fmt.Println("┃ Service Name: ", service.serviceName)
		// 	fmt.Println("┃ Service Type: ", service.serviceType)
		// 	fmt.Println("┃ Service Target: ", service.serviceTarget)
		// 	fmt.Println("┃ Service Target Port: ", service.serviceTargetname)
		// 	fmt.Println("┃")
		// 	fmt.Println("┃")
		// }
		// for _, ingress := range info.nsingresses {
		// 	fmt.Println("┃ Ingress Name: ", ingress.ingressName)
		// 	fmt.Println("┃ Ingress Host: ", ingress.ingressHost)
		// 	fmt.Println("┃ Ingress Path: ", ingress.ingressPath)
		// 	fmt.Println("┃ Ingress Target: ", ingress.ingressTarget)
		// 	fmt.Println("┃ Ingress Target Port: ", ingress.ingressTargetPort)
		// 	fmt.Println("┃ Ingress LB Type: ", ingress.ingressesLBtype)
		// 	fmt.Println("┃")
		// 	fmt.Println("┃")
		// }

		// for _, deployment := range info.nsdeployments {
		// 	fmt.Println("┃ Deployment Name: ", deployment.deploymentName)
		// 	fmt.Println("┃ Deployment Replicas: ", deployment.deploymentReplicas)
		// 	fmt.Println("┃ Deployment Status: ", deployment.deploymentStatus)
		// 	fmt.Println("┃ Deployment Strategy: ", deployment.deploymentStrategy)
		// 	for _, image := range deployment.deploymentImages {
		// 		fmt.Println(" Deployment Image: ", image)
		// 	}
		// 	fmt.Println("┃")
		// 	fmt.Println("┃")
		// }
		// for _, statefulset := range info.nsstatefulsets {
		// 	fmt.Println("┃ Statefulset Name: ", statefulset.statefulsetName)
		// 	fmt.Println("┃ Statefulset Replicas: ", statefulset.statefulsetReplicas)
		// 	fmt.Println("┃ Statefulset Status: ", statefulset.statefulsetStatus)
		// 	fmt.Println("┃ Statefulset Strategy: ", statefulset.statefulsetStrategy)
		// 	for _, image := range statefulset.statefulsetImages {
		// 		fmt.Println(" Statefulset Image: ", image)
		// 	}

		// 	fmt.Println("┃")
		// 	fmt.Println("┃")
		// }

		// for _, daemonset := range info.nsdaemonsets {
		// 	fmt.Println("┃ Daemonset Name: ", daemonset.daemonsetName)
		// 	fmt.Println("┃ Daemonset Namespace: ", daemonset.daemonsetNamespace)
		// 	fmt.Println("┃ Daemonset Labels: ", daemonset.daemonsetLabels)
		// 	fmt.Println("┃ Daemonset Matchlables: ", daemonset.daemonsetMatchlables)
		// 	for _, container := range daemonset.daemonsetContainers {
		// 		fmt.Println(" Daemonset Container: ", container)

		// 	}
		// 	fmt.Println("┃ Daemonset Terminationperiod: ", daemonset.daemonsetTerminationperiod)
		// 	for _, volume := range daemonset.daemonsetVolumes {
		// 		fmt.Println(" Daemonset Volume: ", volume)
		// 	}
		// 	for _, volumemount := range daemonset.daemonsetVolumemount {
		// 		fmt.Println(" Daemonset Volumemount: ", volumemount)
		// 	}
		// 	fmt.Println("┃ Daemonset Status: ", daemonset.daemonsetStatus)
		// 	fmt.Println("┃")
		// 	fmt.Println("┃")
		// }

		// for _, confimap := range info.nsconfimap {
		// 	fmt.Println("┃ Configmap Name: ", confimap.confimapName)
		// 	fmt.Println("┃")
		// 	fmt.Println("┃")
		// }

		// for _, secret := range info.nssecret {
		// 	fmt.Println("┃ Secret Name: ", secret.secretName)
		// 	fmt.Println("┃")
		// 	fmt.Println("┃")
		// }

		for _, pvc := range info.nspersistentvolumeclaims {
			fmt.Println("┃ PVC Name: ", pvc.pvc_Name)
			fmt.Println("┃ PVC Labels: ", pvc.pvc_Labels)
			fmt.Println("┃ PVC Status: ", pvc.pvc_tatus)
			fmt.Println("┃ PVC Volume: ", pvc.pvc_Volume)
			fmt.Println("┃ PVC Size: ", pvc.pvc_Size)
			fmt.Println("┃ PVC Access Mode: ", pvc.pvc_AcessMode)
			fmt.Println("┃")
			fmt.Println("┃")
		}

		fmt.Println("┗━━━━━━ ")
	}
}
