type Monitoring struct {
	ClusterName string
	Cpu         float64
	Cores       int64
	Nodes       int64
	Totalmemory float64
	Usedmemory  float64
	Disk        float64
	Totaldisk   float64
	Billing     float64
}

type NodeInfo struct {
	Name    []string
	Memory  []float64
	CPU     []float64
	Disk    []float64
	CpuTemp []float64
	IP      []string
}
type resources struct {
	pods                   string
	services               string
	ingresses              string
	deployments            string
	statefulsets           string
	daemonsets             string
	confimap               string
	secret                 string
	namespaces             string
	persistentvolumes      string
	persistentvolumeclaims string
}
type namespacestruct struct {
	namespaceName            string
	nspods                   []string
	nsservices               []string
	nsingresses              []string
	nsdeployments            []string
	nsstatefulsets           []string
	nsdaemonsets             []string
	nsconfimap               []string
	nssecret                 []string
	nspersistentvolumeclaims []string
	nspersistentvolumes      []string
	nsjobs                   []string
}

type namespaceInfoDetailed struct {
	namespaceName            string
	nspods                   []nspods
	nsservices               []nsservices
	nsingresses              []nsingresses
	nsdeployments            []nsdeployments
	nsstatefulsets           []nsstatefulsets
	nsdaemonsets             []nsdaemonsets
	nsconfimap               []nsconfimap
	nssecret                 []nssecret
	nspersistentvolumeclaims []nspersistentvolumeclaims
	// nspersistentvolumes      []string
	// nsjobs                   [][]string
}

// key value pairs pods for resoueces and limits
type nspods struct {
	podName  string
	poStatus string
	podCPU   string
	podMEM   string
	poImages []string
}

type nsservices struct {
	serviceName       string
	serviceType       string
	serviceTarget     string
	serviceTargetname string
}

// key value for ingresses
type nsingresses struct {
	ingressName       string
	ingressHost       string
	ingressPath       string
	ingressTarget     string
	ingressTargetPort string
	ingressesLBtype   string
}

type nsdeployments struct {
	deploymentName     string
	deploymentReplicas string
	deploymentStatus   string
	deploymentStrategy string
	deploymentImages   []string
}

type nsstatefulsets struct {
	statefulsetName     string
	statefulsetReplicas string
	statefulsetStatus   string
	statefulsetStrategy string
	statefulsetImages   []string
}

type nsdaemonsets struct {
	// what does deamonset have?
	daemonsetName              string
	daemonsetNamespace         string
	daemonsetLabels            map[string]string
	daemonsetMatchlables       map[string]string
	daemonsetContainers        []string
	daemonsetTerminationperiod string
	daemonsetVolumes           []string
	daemonsetVolumemount       []string
	daemonsetStatus            string
}

type nspersistentvolumeclaims struct {
	pvc_Name      string
	pvc_Labels    map[string]string
	pvc_tatus     string
	pvc_Volume    string
	pvc_Size      string
	pvc_AcessMode string
	// pvc_StorageClass string
}

type nsconfimap struct {
	confimapName string
}

type nssecret struct {
	secretName string
	// isusedBy   string
}

var NamespacesListDetailed []namespaceInfoDetailed