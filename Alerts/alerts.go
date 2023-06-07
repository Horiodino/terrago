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
}

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

		// image pull failure alert
	}
	pods, err := clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	// Print the names of all pods
	for _, pod := range pods.Items {
		fmt.Println("┏━━━━━━━━━━━━━ " + pod.Name + " ━━━━━━━━━━━━━━━━━━━━━━━┓")

		// Print the name of the pod and the status of each container within the pod
		// for _, containerStatus := range pod.Status.ContainerStatuses {
		// 	fmt.Println("Pod Status: " + string(containerStatus.State.Waiting.Reason))
		// }

		create := pod.CreationTimestamp
		current_time := time.Now()

		age := current_time.Sub(create.Time).Minutes()

		if age > 5 && pod.Status.Phase != "Running" {

			for _, containerStatus := range pod.Status.ContainerStatuses {
				fmt.Println("Pod Status: " + string(containerStatus.State.Waiting.Reason))
			}
			fmt.Println("Pod is not running")

		}

		pod_restart := pod.Status.ContainerStatuses[0].RestartCount

		if age > 5 && pod.Status.Phase == "Running" && pod_restart > 3 {
			for _, containerStatus := range pod.Status.ContainerStatuses {
				fmt.Println("Pod Status: " + string(containerStatus.State.Waiting.Reason))
			}

			fmt.Println("Pod is restarting")

			// get pod ip
			pod_ip := pod.Status.PodIP
			fmt.Println(pod_ip)

		}

		fmt.Println("┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
	}

}
