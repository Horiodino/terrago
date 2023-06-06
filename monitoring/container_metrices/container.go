package container_metrices

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/metrics/pkg/apis/metrics/v1beta1"
	"k8s.io/metrics/pkg/client/clientset/versioned"
)

type containerMetrics struct {
	// conatainer name of type var containerNames
	cname          string
	podName        string
	nsname         string
	node           string
	requestecpu    string
	limitscpu      string
	requestememory string
	limitsmemory   string
	diskIo         string
	networkTx      int
	networkRx      int
	containerImage string
	// volumemounts   []string
}

var ContainerInfo []containerMetrics

var AllContainer int32

func Containermatricesinfo() {
	config, err := clientcmd.BuildConfigFromFlags("", os.Getenv("HOME")+"/.kube/config")
	if err != nil {
		log.Fatal(err)
	}

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
			cname := container.Name
			podName := pod.Name
			nsname := pod.Namespace
			node := pod.Spec.NodeName
			cpu_request := container.Resources.Requests.Cpu()
			cpu_limits := container.Resources.Limits.Cpu()
			memory_request := container.Resources.Requests.Memory()
			memory_limits := container.Resources.Limits.Memory()

			containerimage := container.Image
			// vmounts := container.VolumeMounts

			fmt.Println("Container Name: ", cname)
			fmt.Println("Pod Name: ", podName)
			fmt.Println("Namespace Name: ", nsname)
			fmt.Println("Node Name: ", node)
			fmt.Println("Requested CPU: ", cpu_request)
			fmt.Println("Limits CPU: ", cpu_limits)
			fmt.Println("Requested Memory: ", memory_request)
			fmt.Println("Limits Memory: ", memory_limits)

			// fmt.Println("Disk IO: ", diskIo)

			ContainerMetrics := containerMetrics{
				cname:          cname,
				podName:        podName,
				nsname:         nsname,
				node:           node,
				requestecpu:    cpu_request.String(),
				limitscpu:      cpu_limits.String(),
				requestememory: memory_request.String(),
				limitsmemory:   memory_limits.String(),
				// diskIo:         diskIo,
				containerImage: containerimage,
				// volumemounts:   string(vmounts),
			}

			ContainerInfo = append(ContainerInfo, ContainerMetrics)
		}

	}
}

func GetCpuUsage(containerName string, podName string, namespace string) (float64, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return 0, fmt.Errorf("failed to get cluster config: %v", err)
	}

	// create the Kubernetes API clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return 0, fmt.Errorf("failed to create Kubernetes API clientset: %v", err)
	}
	metricsClientset, err := versioned.NewForConfig(config)
	if err != nil {
		return 0, fmt.Errorf("failed to create Kubernetes Metrics API clientset: %v", err)
	}

	// get the pod by name and namespace
	pod, err := clientset.CoreV1().Pods(namespace).Get(context.Background(), podName, v1.GetOptions{})
	if err != nil {
		return 0, fmt.Errorf("failed to get pod: %v", err)
	}
	// pod.GetCreationTimestamp() tells us when the pod was created
	startTime := pod.GetCreationTimestamp().Time

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
	// this elapsed time is the time since the pod was created
	elapsedTime := time.Since(startTime)
	cpuUsagePercent := float64(cpuUsage) / float64(elapsedTime.Nanoseconds()) * 100
	return cpuUsagePercent, nil
}
