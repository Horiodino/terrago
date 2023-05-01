// here we will define the visual monitoring like how your app is performing and shoing like a diagram
// what are the thing are  connected to each other and how they are connected
// its just like we will show a cluster diagram in the monitoring page and we will show the deplotments
// and what thing are connected to each other
//for an example we will show the diagram like this
// deployment1 -> service1 -> pod1 -> node1
// deployment2 -> service2 -> pod2 -> node2
//   |
//   |----pvc1 -> pv -> database -> node3

//how we will show the diagram like this
// we will use the kubernetes client to get the info regarding the cluster and its components
// and we already have the info regarding the cluster and its components
// we will start with the nodes and we will get the info regarding the nodes
// make sure them with the help of labels and annotations and we will show the diagram

// get the info regarding the nodes
// we already have the info regarding the nodes

func nodeVisuals() {
	// we will get the info regarding the nodes
	// we already have the info regarding the nodes in the monitoring.go file / nodeinfo slice

	// first we have deployments and we will get the info regarding the deployments
	// we will get the info regarding the deployments from the monitoring.go file resources struct /deployments slice

	// when need the info regarding the deployments
	// the info is like we neede the deployment name, namespace, replicas, labels, annotations, etc

	// get the info regarding the deployments
	// using the kubernetes client

	// create the kubernetes client
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	// get the deployments of the cluster
	deployments, err := clientset.AppsV1().Deployments("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

}