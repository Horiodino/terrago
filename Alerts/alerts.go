package Alerts

import (
	"context"
	"log"
	"os"
	"time"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type FailureStatus struct {
	Deployment string
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
			// send alert
			log.Println("Deployment is not ready")
		}

		// image pull failure alert
	}

	Pods, err := clientset.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, pod := range Pods.Items {

		// get pod start time
		created := pod.CreationTimestamp
		current_time := time.Now()

		pod_age := current_time.Sub(created.Time).Minutes()

		// get pod restart count
		restart_count := pod.Status.ContainerStatuses[0].RestartCount

		pod_status := pod.Status.Phase

		if pod_status == "CrashLoopBackOff" && pod_age > 2 && restart_count > 2 {
			// send alert
			log.Println("Pod is in CrashLoopBackOff state")
		}

		if pod_status == "Failed" && pod_age > 2 && restart_count > 2 {
			// send alert
			log.Println("Pod is in Failed state")
		}

		if pod_status == "Pending" && pod_age > 2 && restart_count > 2 {
			// send alert
			log.Println("Pod is in Pending state")
		}

		// now get the logs of the pod

		podLogs, err := clientset.CoreV1().Pods(pod.Name).GetLogs("", &v1.PodLogOptions{}).Stream(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		log.Println(podLogs)
	}

}
