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
	// State   string
	Restart string
	Ready   bool
	Labels  map[string]string
}

var FailureStatusSlice []FailureStatus

func PodFailure() {

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

			// State := pod.Status.ContainerStatuses[0].State

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
				// State:   string(State.Waiting.Reason),
				Restart: fmt.Sprintf("%d", restart_count),
				Ready:   ready,
				Labels:  labels,
			}

			PodsFailureSlice = append(PodsFailureSlice, PodsFailure)

		}

		if age > 2 && pod.Status.Phase != "Ready" && pod_restart > 3 && pod.Status.Phase != "Running" {
			pod_ip := pod.Status.PodIP

			container_id := pod.Status.ContainerStatuses[0].ContainerID

			Image := pod.Spec.Containers[0].Image

			// State := pod.Status.ContainerStatuses[0].State

			restart_count := pod_restart

			ready := pod.Status.ContainerStatuses[0].Ready

			labels := pod.Labels

			// Getlogs("pod", pod.Name, "default", pod.Spec.Containers[0].Name)

			PodsFailure := PodsFailure{
				PodName: pod.Name,
				Created: create.String(),
				Age:     fmt.Sprintf("%f", age),
				Status:  string(pod.Status.Phase),
				IP:      pod_ip,
				CID:     container_id,
				Image:   Image,
				// State:   string(State.Waiting.Reason),
				Restart: fmt.Sprintf("%d", restart_count),
				Ready:   ready,
				Labels:  labels,
			}

			PodsFailureSlice = append(PodsFailureSlice, PodsFailure)

		}
		if age > 2 && pod.Status.Phase == "Pending" {

			pod_ip := pod.Status.PodIP

			container_id := pod.Status.ContainerStatuses[0].ContainerID

			Image := pod.Spec.Containers[0].Image

			// State := pod.Status.ContainerStatuses[0].State

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
				// State:   string(State.Waiting.Reason),
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
			fmt.Println("Restart: " + PodsFailure.Restart)
			fmt.Println("Ready: " + fmt.Sprintf("%t", PodsFailure.Ready))
			fmt.Println("Labels: " + fmt.Sprintf("%v", PodsFailure.Labels))
			fmt.Println("┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		}
	}
}

func 
