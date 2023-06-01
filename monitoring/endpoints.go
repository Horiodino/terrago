package monitoring

import (
	"context"
	"encoding/json"
	"log"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// getEndpoints will get the endpoints of the cluster using the kubernetes API
func getEndpoints() {

	config, err := clientcmd.BuildConfigFromFlags("", os.Getenv("HOME")+"/.kube/config")
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	endpoints, err := clientset.CoreV1().Endpoints("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	jsonData, err := json.Marshal(endpoints)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("allednpoints.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.Write(jsonData)

}

func getHTTPRequests() {

}

func exposedServices() {

}
