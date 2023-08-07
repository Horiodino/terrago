package api

import (
	"fmt"
	"log"

	"k8s.io/metrics/pkg/client/clientset/versioned"

	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ClusterInfo *ClusterInfoStruct

func (CLIENT *K8sclient) PullClusterInfo() (string, error) {
	Cluster := &ClusterInfoStruct{}
	MetricresClient, err := versioned.NewForConfig(CLIENT.config)
	if err != nil {
		fmt.Println(err)
	}

	metrics, err := MetricresClient.MetricsV1beta1().NodeMetricses().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, node := range metrics.Items {
		Cluster.Cpu += float64((node.Usage.Cpu().MilliValue()))
	}
	nodes, err := CLIENT.clientset.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, node := range nodes.Items {
		Cluster.Cores += (node.Status.Capacity.Cpu().Value())
	}

	for _, node := range metrics.Items {
		Cluster.Usedmemory += float64((node.Usage.Memory().Value()) / 1000000)
	}

	for _, node := range nodes.Items {
		Cluster.Totalmemory += float64((node.Status.Capacity.Memory().Value()) / 1000000)
	}
	for _, node := range metrics.Items {
		Cluster.Disk += float64((node.Usage.StorageEphemeral().Value()) / 1000000)
	}

	for _, node := range nodes.Items {
		Cluster.Totaldisk += float64((node.Status.Capacity.StorageEphemeral().Value()) / 1000000)
	}
	Cluster.Nodes = int64(len(nodes.Items))
	ClusterInfo = &ClusterInfoStruct{
		ClusterName: Cluster.ClusterName,
		Cpu:         Cluster.Cpu,
		Cores:       Cluster.Cores,
		Nodes:       Cluster.Nodes,
		Totalmemory: Cluster.Totalmemory,
		Usedmemory:  Cluster.Usedmemory,
		Disk:        Cluster.Disk,
		Totaldisk:   Cluster.Totaldisk,
		Billing:     Cluster.Billing,
	}

	fmt.Println("ClusterName: ", Cluster.ClusterName)
	fmt.Println("Cpu: ", Cluster.Cpu)
	fmt.Println("Cores: ", Cluster.Cores)
	fmt.Println("Nodes: ", Cluster.Nodes)
	fmt.Println("Totalmemory: ", Cluster.Totalmemory)
	fmt.Println("Usedmemory: ", Cluster.Usedmemory)
	fmt.Println("Disk: ", Cluster.Disk)
	fmt.Println("Totaldisk: ", Cluster.Totaldisk)
	fmt.Println("Billing: ", Cluster.Billing)

	return "ClusterInfo", nil
}

func (CLIENT *K8sclient) PullNodeInfo() (string, error) {
	return "NodeInfo", nil
}

func (CLIENT *K8sclient) PullPodInfo() (string, error) {
	return "PodInfo", nil
}

func (CLIENT *K8sclient) PullDeploymentInfo() (string, error) {
	return "DeploymentInfo", nil
}

func (CLIENT *K8sclient) PullStatefulSetInfo() (string, error) {
	return "StatefulSetInfo", nil
}

func (CLIENT *K8sclient) PullDaemonSetInfo() (string, error) {
	return "DaemonSetInfo", nil
}

func (CLIENT *K8sclient) PullReplicaSetInfo() (string, error) {
	return "ReplicaSetInfo", nil
}

func (CLIENT *K8sclient) PullReplicationControllerInfo() (string, error) {
	return "ReplicationControllerInfo", nil
}

func (CLIENT *K8sclient) PullHorizontalPodAutoscalerInfo() (string, error) {
	return "HorizontalPodAutoscalerInfo", nil
}

func (CLIENT *K8sclient) PullPodDisruptionBudgetInfo() (string, error) {
	return "PodDisruptionBudgetInfo", nil
}

func (CLIENT *K8sclient) PullNetworkPolicyInfo() (string, error) {
	return "NetworkPolicyInfo", nil
}

func (CLIENT *K8sclient) PullPodSecurityPolicyInfo() (string, error) {
	return "PodSecurityPolicyInfo", nil
}

func (CLIENT *K8sclient) PullLimitRangeInfo() (string, error) {
	return "LimitRangeInfo", nil
}
