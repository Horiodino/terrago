package clustermetrices

// here we will get the info regarding the cluster and its components

import (
	"context"
	"fmt"
	"os"

	// Kubernetes API client libraries and packages
	// ".mongodb.org/mongo-driver/mongo"go

	"github.com/fatih/color"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
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

type Monitoring struct {
	ClusterName string
	Cpu         float64
	Cores       int64
	Nodes       int64
	Totalmemory float64
	Usedmemory  float64
	Disk        float64
	Totaldisk   float64
	Billing     float64
}

var M *Monitoring

func Getinfo() {
	M = &Monitoring{}

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
		cpuUsage += (node.Usage.Cpu().MilliValue())
	}
	nodes, err := clientset.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	var cpuCores int64
	for _, node := range nodes.Items {
		cpuCores += (node.Status.Capacity.Cpu().Value())
	}
	var memoryUsage int64
	for _, node := range metrics.Items {
		memoryUsage += (node.Usage.Memory().Value()) / 1000000
	}
	var memory int64
	for _, node := range nodes.Items {
		memory += (node.Status.Capacity.Memory().Value()) / 1000000
	}
	var diskUsage int64
	for _, node := range metrics.Items {
		diskUsage += (node.Usage.StorageEphemeral().Value()) / 1000000
	}
	var disk int64
	for _, node := range nodes.Items {
		disk += (node.Status.Capacity.StorageEphemeral().Value()) / 1000000
	}

	// now we have all the info regar	"github.com/fatih/color"ding the cpu usage, memory usage, disk usage, etc of the entire cluster
	// now append the info to the struct
	M.Cpu = float64(cpuUsage)
	M.Cores = cpuCores
	M.Nodes = int64(len(nodes.Items))
	M.Totalmemory = float64(memory)
	M.Usedmemory = float64(memoryUsage)
	M.Disk = float64(diskUsage)
	M.Totaldisk = float64(disk)

}

func Display() {
	// display the stored info of Getinfo function that is stored in the struct
	color.Green("--------------------Cluster Info---------------------")
	fmt.Println("CPU Usage: ", M.Cpu)
	fmt.Println("CPU Cores: ", M.Cores)
	fmt.Println("Nodes: ", M.Nodes)
	fmt.Println("Total Memory: ", M.Totalmemory, "Mb")
	fmt.Println("Used Memory: ", M.Usedmemory, "Mb")
	fmt.Println("Disk Usage: ", M.Disk, "Mb")
	fmt.Println("Total Disk: ", M.Totaldisk, "Mb")
	color.Green("-----------------------------------------------------")

}

type NodeInfo struct {
	Name    []string
	Memory  []float64
	CPU     []float64
	Disk    []float64
	CpuTemp []float64
	IP      []string
}

var NodeInfoList []NodeInfo

// cpu usage for the nodes
func Cpu() ([]NodeInfo, error) {
	// we will get the info regarding the cpu usage for the nodes
	// we will use kubernetes API for this

	// creat the kubernetes client
	config, err := clientcmd.BuildConfigFromFlags("", os.Getenv("HOME")+"/.kube/config")
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
		cpuUsage := metrics.Usage.Cpu().MilliValue()
		cpuCores := node.Status.Capacity.Cpu().MilliValue()
		cpuUsagePercentage := float64(cpuUsage) / float64(cpuCores) * 100
		memoryUsage := metrics.Usage.Memory().Value()
		// get the memory capacity for the nodes
		memoryCapacity := node.Status.Capacity.Memory().Value()
		// get the memory usage percentage for the nodes
		memoryUsagePercentage := float64(memoryUsage) / float64(memoryCapacity) * 100
		diskUsage := metrics.Usage.StorageEphemeral().Value()
		diskCapacity := node.Status.Capacity.StorageEphemeral().Value()
		diskUsagePercentage := float64(diskUsage) / float64(diskCapacity) * 100

		Nodeip := node.Status.Addresses[0].Address

		nodeInfo := NodeInfo{
			Name:   []string{node.Name},
			CPU:    []float64{cpuUsagePercentage},
			Memory: []float64{memoryUsagePercentage},
			Disk:   []float64{diskUsagePercentage},
			IP:     []string{Nodeip},
		}

		NodeInfoList = append(NodeInfoList, nodeInfo)
	}

	return NodeInfoList, err

}

type Resources struct {
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

var ResourcesList []Resources

func ClusterInfo() ([]Resources, error) {

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

	ResourcesList = append(ResourcesList, Resources{

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

	return ResourcesList, nil
}
func ObjectDisplay() {
	for _, resource := range ResourcesList {
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

var NamespacesList []namespacestruct

func NamespacesInfo() ([]namespacestruct, error) {
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
		//gett all the info regarding the namespcace
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
		persistentvolumes, err := clientset.CoreV1().PersistentVolumes().List(context.Background(), metav1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}

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

		NamespacesList = append(NamespacesList, namespaceInfo)
	}

	return NamespacesList, nil
}
func Displaynsinfo() {
	for _, namespace := range NamespacesList {
		fmt.Printf("Namespace Name: %s\n", namespace.namespaceName)
		fmt.Printf("Pods: %s\n", namespace.nspods)
		fmt.Printf("Services: %s\n", namespace.nsservices)
		fmt.Printf("Ingresses: %s\n", namespace.nsingresses)
		fmt.Printf("Deployments: %s\n", namespace.nsdeployments)
		fmt.Printf("Statefulsets: %s\n", namespace.nsstatefulsets)
		fmt.Printf("Daemonsets: %s\n", namespace.nsdaemonsets)
		fmt.Printf("Configmaps: %s\n", namespace.nsconfimap)
		fmt.Printf("Secrets: %s\n", namespace.nssecret)
		fmt.Printf("Persistent Volume Claims: %s\n", namespace.nspersistentvolumeclaims)
		fmt.Printf("Persistent Volumes: %s\n", namespace.nspersistentvolumes)
		fmt.Printf("Jobs: %s\n", namespace.nsjobs)
	}
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

var NamespacesListDetailed []namespaceInfoDetailed

func GetnamespaceInfoDetailed() {
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

		NamespacesListDetailed = append(NamespacesListDetailed, namespaceInfoDetailed)
	}
}
