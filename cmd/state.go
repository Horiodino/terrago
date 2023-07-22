package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fatih/color"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func metriceserverstate() {

	config, err := clientcmd.BuildConfigFromFlags("", os.Getenv("HOME")+"/.kube/config")
	if err != nil {
		log.Fatal(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	var duration time.Duration = 30
	var i int
	for ; i < int(duration); i++ {
		status, err := clientset.AppsV1().Deployments("kube-system").Get(context.Background(), "metrics-server", metav1.GetOptions{})
		if err != nil {
			fmt.Println(err)
		}

		time.Sleep(1 * time.Second)
		if status.Status.ReadyReplicas >= 1 {
			fmt.Println()
			color.Blue("Metrics server is up and running")
			break
		}

	}
}
