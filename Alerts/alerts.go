package Alerts

import (
	"context"
	"log"
	"os"

	metrices "github.com/Horiodino/terrago/monitoring"
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
	// get

	Deployments, err := clientset.AppsV1().Deployments(namespace.Name).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, deployment := range Deployments.Items {
		if deployment.Status.Replicas != deployment.Status.ReadyReplicas {
			log.Println("Deployment", deployment.Name, "is not ready")
		}
	}

	metrices.NamespacesListDetailed()

}
