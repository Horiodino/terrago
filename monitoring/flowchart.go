// here we will define the visual monitoring like how your app is performing and shoing like a diagram
// what are the thing are  connected to each other and how they are connected
// its just like we will show a cluster diagram in the monitoring page and we will show the deplotments
// and what thing are connected to each other
//for an example we will show the diagram like this
// deployment1 -> service1 -> pod1 -> node1
// deployment2 -> service2 -> pod2 -> node2
//   |
//   |----pvc1 -> pv -> database -> node3

//how we will show the diagram like this
// we will use the kubernetes client to get the info regarding the cluster and its components
// and we already have the info regarding the cluster and its components
// we will start with the nodes and we will get the info regarding the nodes
// make sure them with the help of labels and annotations and we will show the diagram

// get the info regarding the nodes
// we already have the info regarding the nodes
package monitoring

import (
	"context"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// slice for deployment info

var deploymentInfoSlice []deploymentInfo

type deploymentInfo struct {
	deploymentName string
	namespace      string
	replicas       int32
	labels         map[string]string
	annotations    map[string]string
}

func nodeVisuals() {
	// we will get the info regarding the nodes
	// we already have the info regarding the nodes in the monitoring.go file / nodeinfo slice

	// first we have deployments and we will get the info regarding the deployments
	// we will get the info regarding the deployments from the monitoring.go file resources struct /deployments slice

	// when need the info regarding the deployments
	// the info is like we neede the deployment name, namespace, replicas, labels, annotations, etc

	// get the info regarding the deployments
	// using the kubernetes client

	// create the kubernetes client
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	// get the deployments of the cluster
	deployments, err := clientset.AppsV1().Deployments("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	// now we will get inside the deployments and we will get the info regarding the deployments
	// we will get the info regarding the deployments and we will store them in a slice

	// now get the info regarding the deployments and store them in a slice
	for _, deployment := range deployments.Items {
		deplname := deployment.Name
		deplns := deployment.Namespace
		deplreplica := deployment.Spec.Replicas
		depllable := deployment.Labels
		deplannotation := deployment.Annotations

		// now we will store the info regarding the deployments in a slice
		// we will store the info in the deploymentInfo struct

		// now  append the info regarding the deployments in the deploymentInfo struct
		deploymentInfo := deploymentInfo{
			deploymentName: deplname,
			namespace:      deplns,
			replicas:       *deplreplica,
			labels:         depllable,
			annotations:    deplannotation,
		}

		// now we will append the info regarding the deployments in the deploymentInfo slice
		deploymentInfoSlice = append(deploymentInfoSlice, deploymentInfo)

	}
}

var serviceInfoSlice []serviceInfo

type serviceInfo struct {
	serviceName string
	namespace   string
	labels      map[string]string
	annotations map[string]string
	targetport  float32
}

func serviceVisuals() {

	// now we will get the info regarding the services
	// kuebrnetes client

	// create the kubernetes client
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	// get the services of the cluster
	services, err := clientset.CoreV1().Services("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	// now we will get inside the services and we will get the info regarding the services
	// we will get the info regarding the services and we will store them in a slice

	// now get the info regarding the services and store them in a slice
	for _, service := range services.Items {
		servicename := service.Name
		servicens := service.Namespace
		servicelable := service.Labels
		serviceannotation := service.Annotations
		serviceport := service.Spec.Ports

		// now we will store the info regarding the services in a slice
		// we will store the info in the serviceInfo struct

		// now  append the info regarding the services in the serviceInfo struct
		serviceInfo := serviceInfo{
			serviceName: servicename,
			namespace:   servicens,
			labels:      servicelable,
			annotations: serviceannotation,
			targetport:  serviceport[0].TargetPort.Float32(),
		}

		// now we will append the info regarding the services in the serviceInfo slice
		serviceInfoSlice = append(serviceInfoSlice, serviceInfo)
	}
}

var podInfoSlice []podInfo

type podInfo struct {
	podName      string
	namespace    string
	labels       map[string]string
	annotations  map[string]string
	replias      int32
	containers   []string
	containerImg []string
}

func podVisuals() {

	// kuebrnetes client
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	// get the pods of the cluster
	pods, err := clientset.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	// now we will get inside the pods and we will get the info regarding the pods
	// we will get the info regarding the pods and we will store them in a slice

	// now get the info regarding the pods and store them in a slice
	for _, pod := range pods.Items {

		podname := pod.Name
		podns := pod.Namespace
		podlable := pod.Labels
		podannotation := pod.Annotations
		// podcontainer := pod.Spec.Containers

		// now we will store the info regarding the pods in a slice
		// we will store the info in the podInfo struct

		// now  append the info regarding the pods in the podInfo struct
		podInfo := podInfo{
			podName:     podname,
			namespace:   podns,
			labels:      podlable,
			annotations: podannotation,
		}
		// now we will append the info regarding the pods in the podInfo slice
		podInfoSlice = append(podInfoSlice, podInfo)
	}
}
