package Alerts

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type FailureStatus struct {
	Pods []PodsFailure
}

type PodsFailure struct {
	PodName string
	Created string
	Age     string
	Status  string
	IP      string
	CID     string
	Image   string
	State   string
	Restart string
	Ready   bool
	Labels  map[string]string
}

var FailureStatusSlice []FailureStatus

func DeploumentsFailure() {

	// image pull failure
	config, err := clientcmd.BuildConfigFromFlags("", os.Getenv("HOME")+"/.kube/config")
	if err != nil {
		log.Fatal(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	Deployments, err := clientset.AppsV1().Deployments("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, deployment := range Deployments.Items {

		created := deployment.CreationTimestamp
		current_time := time.Now()

		if deployment.Status.Replicas != deployment.Status.ReadyReplicas && current_time.Sub(created.Time).Minutes() > 5 {

			fmt.Println("Deployment is not ready")
		}

		// image pull failure alert TODO--------------------
	}
	pods, err := clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	// Print the names of all pods
	for _, pod := range pods.Items {

		create := pod.CreationTimestamp
		current_time := time.Now()

		age := current_time.Sub(create.Time).Minutes()

		pod_restart := pod.Status.ContainerStatuses[0].RestartCount

		var PodsFailureSlice []PodsFailure

		if age > 2 && pod.Status.Phase == "Running" && pod_restart > 3 {

			pod_ip := pod.Status.PodIP

			container_id := pod.Status.ContainerStatuses[0].ContainerID

			Image := pod.Spec.Containers[0].Image

			State := pod.Status.ContainerStatuses[0].State

			restart_count := pod_restart

			ready := pod.Status.ContainerStatuses[0].Ready
			labels := pod.Labels

			PodsFailure := PodsFailure{
				PodName: pod.Name,
				Created: create.String(),
				Age:     fmt.Sprintf("%f", age),
				Status:  string(pod.Status.Phase),
				IP:      pod_ip,
				CID:     container_id,
				Image:   Image,
				State:   string(State.Waiting.Reason),
				Restart: fmt.Sprintf("%d", restart_count),
				Ready:   ready,
				Labels:  labels,
			}

			PodsFailureSlice = append(PodsFailureSlice, PodsFailure)

		}

		if age > 2 && pod.Status.Phase != "Ready" && pod_restart > 3 && pod.Status.Phase != "Running" {
			for _, containerStatus := range pod.Status.ContainerStatuses {
				fmt.Println("Pod Status: " + string(containerStatus.State.Waiting.Reason))
			}
			pod_ip := pod.Status.PodIP

			container_id := pod.Status.ContainerStatuses[0].ContainerID

			Image := pod.Spec.Containers[0].Image

			State := pod.Status.ContainerStatuses[0].State

			restart_count := pod_restart

			ready := pod.Status.ContainerStatuses[0].Ready

			labels := pod.Labels

			PodsFailure := PodsFailure{
				PodName: pod.Name,
				Created: create.String(),
				Age:     fmt.Sprintf("%f", age),
				Status:  string(pod.Status.Phase),
				IP:      pod_ip,
				CID:     container_id,
				Image:   Image,
				State:   string(State.Waiting.Reason),
				Restart: fmt.Sprintf("%d", restart_count),
				Ready:   ready,
				Labels:  labels,
			}

			PodsFailureSlice = append(PodsFailureSlice, PodsFailure)

		}
		FailureStatus := FailureStatus{
			Pods: PodsFailureSlice,
		}

		FailureStatusSlice = append(FailureStatusSlice, FailureStatus)
	}
	for _, FailureStatus := range FailureStatusSlice {
		for _, PodsFailure := range FailureStatus.Pods {
			fmt.Println("┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
			fmt.Println("Pod Name: " + PodsFailure.PodName)
			fmt.Println("Created: " + PodsFailure.Created)
			fmt.Println("Age: " + PodsFailure.Age)
			fmt.Println("Status: " + PodsFailure.Status)
			fmt.Println("IP: " + PodsFailure.IP)
			fmt.Println("CID: " + PodsFailure.CID)
			fmt.Println("Image: " + PodsFailure.Image)
			fmt.Println("State: " + PodsFailure.State)
			fmt.Println("Restart: " + PodsFailure.Restart)
			fmt.Println("Ready: " + fmt.Sprintf("%t", PodsFailure.Ready))
			fmt.Println("Labels: " + fmt.Sprintf("%v", PodsFailure.Labels))
			fmt.Println("┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		}
	}

	SendFailureAlert()
}

type CpuStatus struct {
	Pods  []PodsCpu
	Nodes []NodesCpu
}

type PodsCpu struct {
	PodName string
	message string
}
type NodesCpu struct {
	NodeName string
}

var CpuStatusSlice []CpuStatus

func Cpu_exceed() {

	config, err := clientcmd.BuildConfigFromFlags("", os.Getenv("HOME")+"/.kube/config")
	if err != nil {
		log.Fatal(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	pods, err := clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	var PodsCpuSlice []PodsCpu
	for _, pod := range pods.Items {

		pod_cpu := pod.Spec.Containers[0].Resources.Requests.Cpu().MilliValue()
		pod_cpu_limit := pod.Spec.Containers[0].Resources.Limits.Cpu().MilliValue()

		if pod_cpu > pod_cpu_limit {

			PodsCpu := PodsCpu{
				PodName: pod.Name,
				message: "Pod CPU exceed",
			}

			PodsCpuSlice = append(PodsCpuSlice, PodsCpu)
		}

	}

	nodes, err := clientset.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	var NodesCpuSlice []NodesCpu

	for _, node := range nodes.Items {

		node_cpu := node.Status.Allocatable.Cpu().MilliValue()
		node_cpu_limit := node.Status.Capacity.Cpu().MilliValue()

		if node_cpu > node_cpu_limit {

			NodesCpu := NodesCpu{
				NodeName: node.Name,
			}

			NodesCpuSlice = append(NodesCpuSlice, NodesCpu)
		}
	}

	CpuStatus := CpuStatus{
		Pods:  PodsCpuSlice,
		Nodes: NodesCpuSlice,
	}

	CpuStatusSlice = append(CpuStatusSlice, CpuStatus)

	for _, CpuStatus := range CpuStatusSlice {
		for _, PodsCpu := range CpuStatus.Pods {
			fmt.Println("┏━━━━━━━━━━━")
			fmt.Println("Pod Name: " + PodsCpu.PodName)
			fmt.Println("Message: " + PodsCpu.message)
			fmt.Println("┗━━━━━━━━━━━")
		}
		for _, NodesCpu := range CpuStatus.Nodes {
			fmt.Println("┏━━━━━━━━━━━")
			fmt.Println("Node Name: " + NodesCpu.NodeName)
			fmt.Println("Message: Node CPU exceed")
			fmt.Println("┗━━━━━━━━━━━")
		}
	}

}

func Memory_exceed() {

	// memory exceed alert TODO--------------------
}

func SendFailureAlert() {

}
