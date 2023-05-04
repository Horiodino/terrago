//  In addition to CPU and memory utilization, it's important to monitor network traffic and latency within the Kubernetes cluster

// getting the incomimg traffic

package Networkmonitoring

// how we can get the total incoming traffic
// simply get the total number of bytes received on all interfaces by the pod

func incomingtraffic() {
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// get the pods
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	// suing network interface to get the incoming traffic

}

func outgoingtraffic() {

}

func networktraffic() {

}

func networklatency() {

}

func networkerrors() {

}

func networkrequests() {

}

func exposedports() {

}

// getting the network throughput

// getting the network packets

// getting the network connections

// getting the network connections

// getting the exposed services

// getting the exposed endpoints

// getting the exposed routes
