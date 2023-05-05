package Containermonitoring

import (
	"context"
	"fmt"
	"log"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/metrics/pkg/apis/metrics/v1beta1"
	"k8s.io/metrics/pkg/client/clientset/versioned"
)

// defining a strunct to store the container metrics and all
type containerMetrics struct {
	// conatainer name of type var containerNames
	cname                 []string
	podName               []string
	nsname                []string
	node                  []string
	cpuUsage              []float64
	memoryUsage           []int
	diskIo                []int
	networkTx             []int
	networkRx             []int
	containerID           []string
	containerImage        []string
	containerStatus       []string
	containerCreationTime []time.Time
	containerStartTime    []time.Time
	containerLabels       []string
}

var AllContainer int32

func containermatricesinfo() {
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
	//this for loop will iterate over all the pods
	for _, pod := range pods.Items {
		// this for loop will iterate over all the containers in the pod
		for _, container := range pod.Spec.Containers {
			containerMetrics := containerMetrics{
				cname:                 []string{container.Name},
				podName:               []string{pod.Name},
				nsname:                []string{pod.Namespace},
				node:                  []string{pod.Spec.NodeName},
				cpuUsage:              []float64{0},
				memoryUsage:           []int{0},
				diskIo:                []int{0},
				networkTx:             []int{0},
				networkRx:             []int{0},
				containerID:           []string{container.ContainerID},
				containerImage:        []string{container.Image},
				containerStatus:       []string{string(pod.Status.Phase)},
				containerCreationTime: []time.Time{pod.CreationTimestamp.Time},
				containerStartTime:    []time.Time{pod.Status.StartTime.Time},
				containerLabels:       []string{pod.Labels[i]},
	}
}

func GetCpuUsage(containerName, podName, namespace string) (float64, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return 0, fmt.Errorf("failed to get cluster config: %v", err)
	}

	// create the Kubernetes API clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return 0, fmt.Errorf("failed to create Kubernetes API clientset: %v", err)
	}

	// create the Kubernetes Metrics API clientset
	metricsClientset, err := versioned.NewForConfig(config)
	if err != nil {
		return 0, fmt.Errorf("failed to create Kubernetes Metrics API clientset: %v", err)
	}

	// get the pod by name and namespace
	pod, err := clientset.CoreV1().Pods(namespace).Get(context.Background(), podName, v1.GetOptions{})
	if err != nil {
		return 0, fmt.Errorf("failed to get pod: %v", err)
	}

	// get the start time of the pod
	startTime := pod.GetCreationTimestamp().Time

	// get the container metrics for the pod
	containerMetrics, err := metricsClientset.MetricsV1beta1().PodMetricses(namespace).Get(context.Background(), podName, v1.GetOptions{})
	if err != nil {
		return 0, fmt.Errorf("failed to get container metrics: %v", err)
	}

	// find the container by name
	//  this *v1beta1.ContainerMetrics is a pointer to the container metrics
	// provided by the Kubernetes Metrics API
	var container *v1beta1.ContainerMetrics
	for _, c := range containerMetrics.Containers {
		if c.Name == containerName {
			container = &c
			break
		}
	}
	if container == nil {
		return 0, fmt.Errorf("container not found: %s", containerName)
	}
	cpuUsage := container.Usage.Cpu().MilliValue()
	elapsedTime := time.Since(startTime)
	cpuUsagePercent := float64(cpuUsage) / float64(elapsedTime.Nanoseconds()) * 100
	return cpuUsagePercent, nil
}
