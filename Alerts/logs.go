package Alerts

import (
	"fmt"
	"log"
	"os"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func Getlogs(resource_type string, resource_name string, namespace string, container string) {

	// use switch case to get logs of different resources

	switch resource_type {
	case "pod":
		Getpodlogs(resource_name, namespace, container)
	case "daemonset":
		Getdaemonsetlogs(resource_name, namespace, container)
	case "job":
		Getjoblogs(resource_name, namespace, container)
	}

}

func Getpodlogs(resource_name string, namespace string, container string) (string, error) {

	config, err := clientcmd.BuildConfigFromFlags("", os.Getenv("HOME")+"/.kube/config")
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	req := clientset.CoreV1().Pods(namespace).GetLogs(resource_name, &v1.PodLogOptions{Container: container})

	fmt.Println("Pod logs +++++++++++++++++++++++++++++++ ")
	fmt.Println(req)

	return "nil", nil

}
func Getdaemonsetlogs(resource_name string, namespace string, container string) {

	config, err := clientcmd.BuildConfigFromFlags("", os.Getenv("HOME")+"/.kube/config")
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	req := clientset.CoreV1().Pods(namespace).GetLogs(resource_name, &v1.PodLogOptions{Container: container})

	fmt.Println(req)

}

func Getjoblogs(resource_name string, namespace string, container string) {

	config, err := clientcmd.BuildConfigFromFlags("", os.Getenv("HOME")+"/.kube/config")
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	req := clientset.CoreV1().Pods(namespace).GetLogs(resource_name, &v1.PodLogOptions{Container: container})

	fmt.Println(req)
}
