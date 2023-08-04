package api

import (
	"context"
	"log"
	"os"
	"time"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type Pullinfo interface {
	PullClusterInfo(c *client) (string, error)
	PullNodeInfo(c *client) (string, error)
	PullPodInfo(c *client) (string, error)
	PullNamespaceInfo(c *client) (string, error)
	PullDeploymentInfo(c *client) (string, error)
	PullServiceInfo(c *client) (string, error)
	PullIngressInfo(c *client) (string, error)
	PullConfigMapInfo(c *client) (string, error)
	PullSecretInfo(c *client) (string, error)
	PullPersistentVolumeInfo(c *client) (string, error)
	PullPersistentVolumeClaimInfo(c *client) (string, error)
	PullStorageClassInfo(c *client) (string, error)
	PullRoleInfo(c *client) (string, error)
	PullRoleBindingInfo(c *client) (string, error)
	PullClusterRoleInfo(c *client) (string, error)
	PullClusterRoleBindingInfo(c *client) (string, error)
	PullServiceAccountInfo(c *client) (string, error)
	PullJobInfo(c *client) (string, error)
	PullCronJobInfo(c *client) (string, error)
	PullStatefulSetInfo(c *client) (string, error)
	PullDaemonSetInfo(c *client) (string, error)
	PullReplicaSetInfo(c *client) (string, error)
	PullReplicationControllerInfo(c *client) (string, error)
	PullHorizontalPodAutoscalerInfo(c *client) (string, error)
	PullPodDisruptionBudPullInfo(c *client) (string, error)
	PullNetworkPolicyInfo(c *client) (string, error)
	PullPodSecurityPolicyInfo(c *client) (string, error)
	PullLimitRangeInfo(c *client) (string, error)
	PullResourceQuotaInfo(c *client) (string, error)
	PullEndpointsInfo(c *client) (string, error)
	PullEventInfo(c *client) (string, error)
	PullComponentStatusInfo(c *client) (string, error)
	PullCustomResourceDefinitionInfo(c *client) (string, error)
	PullMutatingWebhookConfigurationInfo(c *client) (string, error)
	PullValidatingWebhookConfigurationInfo(c *client) (string, error)
	PullPodTemplateInfo(c *client) (string, error)
	PullPodSecurityPolicyTemplateInfo(c *client) (string, error)
	PullLeaseInfo(c *client) (string, error)
	PullPriorityClassInfo(c *client) (string, error)
	PullRuntimeClassInfo(c *client) (string, error)
	PullCertificateSigningRequestInfo(c *client) (string, error)
	PullTokenReviewInfo(c *client) (string, error)
	PullSelfSubjectAccessReviewInfo(c *client) (string, error)
	PullSelfSubjectRulesReviewInfo(c *client) (string, error)
	PullSubjectAccessReviewInfo(c *client) (string, error)
	PullHorizontalPodAutoscalerInfoV2Beta1(c *client) (string, error)
	PullHorizontalPodAutoscalerInfoV1(c *client) (string, error)
	PullPodDisruptionBudgetInfoV1Beta1(c *client) (string, error)
}

type AllInfo interface {
	//take time and data to get the info feom hte file
	SaveClusterInfo(c context.Context, Date, Time time.Time) (string, error)
	SaveNodeInfo(c context.Context, Date, Time time.Time) (string, error)
	SavePodInfo(c context.Context, Date, Time time.Time) (string, error)
	SaveNamespaceInfo(c context.Context, Date, Time time.Time) (string, error)
}

type SavedInfo struct {
	AllInfo
	context.Context
	ClusterInfo *ClusterInfoStruct
	NodeInfo    *NodeInfoStruct
}

type ClusterInfoStruct struct {
}
type NodeInfoStruct struct {
}

type client struct {
	clientset *kubernetes.Clientset
	Pullinfo
}

func NewClient() *client {
	config, err := clientcmd.BuildConfigFromFlags("", os.Getenv("HOME")+"/.kube/config")
	if err != nil {
		log.Fatal(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	return &client{
		clientset: clientset,
	}
}
