package Alerts

import (
	"context"
	"log"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func Getlogs(objtype, name string) {
	switch objtype {
	case "pod":
		podlogs(name)
	case "deployment":
		deploymentlogs(name)
	case "service":
		servicelogs(name)
	default:
		// Handle unknown objtype
	}
}

func podlogs(name string) {
	config, err := clientcmd.BuildConfigFromFlags("", os.Getenv("HOME")+"/.kube/config")
	if err != nil {
		log.Fatal(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	pod, err := clientset.CoreV1().Pods("default").Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		log.Fatal(err)
	}

}

func deploymentlogs(name string) {

	config, err := clientcmd.BuildConfigFromFlags("", os.Getenv("HOME")+"/.kube/config")
	if err != nil {
		log.Fatal(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	deployment, err := clientset.AppsV1().Deployments("default").Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		log.Fatal(err)
	}
}

func servicelogs(name string) {

	congif, err := clientcmd.BuildConfigFromFlags("", os.Getenv("HOME")+"/.kube/config")
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(congif)
	if err != nil {
		log.Fatal(err)
	}

	service, err := clientset.CoreV1().Services("default").Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		log.Fatal(err)
	}

}
