package api

import (
	"fmt"
	"log"

	"k8s.io/metrics/pkg/client/clientset/versioned"

	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	one   = 1
	two   = 2
	three = 3
	four  = 4
	five  = 5
	six   = 6
	seven = 7
	eight = 8
	nine  = 9
)

var ClusterInfo *ClusterInfoStruct

func (CLIENT *K8sclient) PullClusterInfo() (ClusterInfoStruct, error) {
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
	return *ClusterInfo, nil
}

func (CLIENT *K8sclient) PullNodeInfo() (string, error) {

	return "NodeInfo", nil
}

func (CLIENT *K8sclient) PullPodInfo() (PodInfoStruct, error) {
	Pods, err := CLIENT.clientset.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	PodInfo := &PodInfoStruct{}
	for i := 0; i < len(Pods.Items); i++ {
		PodInfo.Name = append(PodInfo.Name, Pods.Items[i].Name)
		PodInfo.Phase = append(PodInfo.Phase, string(Pods.Items[i].Status.Phase))
		PodInfo.StartTime = append(PodInfo.StartTime, Pods.Items[i].Status.StartTime.Format("2006-01-02 15:04:05"))
		PodInfo.PodIP = append(PodInfo.PodIP, Pods.Items[i].Status.PodIP)
		PodInfo.HostIP = append(PodInfo.HostIP, Pods.Items[i].Status.HostIP)
		PodInfo.QOSClass = append(PodInfo.QOSClass, string(Pods.Items[i].Status.QOSClass))
		PodInfo.Message = append(PodInfo.Message, Pods.Items[i].Status.Message)
		PodInfo.Reason = append(PodInfo.Reason, Pods.Items[i].Status.Reason)
		// PodInfo.ContainerStatuses = append(PodInfo.ContainerStatuses, Pods.Items[i].Status.ContainerStatuses)
		// PodInfo.InitContainerStatuses = append(PodInfo.InitContainerStatuses, Pods.Items[i].Status.InitContainerStatuses)
		PodInfo.NominatedNodeName = append(PodInfo.NominatedNodeName, Pods.Items[i].Status.NominatedNodeName)
		// PodInfo.Conditions = append(PodInfo.Conditions, Pods.Items[i].Status.Conditions)
		PodInfo.EphemeralContainerStatuses = append(PodInfo.EphemeralContainerStatuses, Pods.Items[i].Status.NominatedNodeName)

	}

	return *PodInfo, nil
}

func (CLIENT *K8sclient) PullDeploymentInfo() (DeploymentInfoStruct, error) {

	Deployments, err := CLIENT.clientset.AppsV1().Deployments("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	Deploy := &DeploymentInfoStruct{}
	for i := 0; i < len(Deployments.Items); i++ {
		Deploy.Name = append(Deploy.Name, Deployments.Items[i].Name)
		Deploy.Namespace = append(Deploy.Namespace, Deployments.Items[i].Namespace)
		Deploy.CreationTimestamp = append(Deploy.CreationTimestamp, Deployments.Items[i].CreationTimestamp.Format("2006-01-02 15:04:05"))
		// #TODO ADD MORE INFO
	}

	return *Deploy, nil
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
