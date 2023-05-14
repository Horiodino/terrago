//  In addition to CPU and memory utilization, it's important to monitor network traffic and latency within the Kubernetes cluster

// getting the incomimg traffic

package Networkmonitoring

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	// "C:\Users\Holiodin\webgo\monitoring\Containermonitoring\metrices.go"
)

// how we can get the total incoming traffic
// simply get the total number of bytes received on all interfaces by the pod

// for ipconfig
type Interface struct {
	Name       string `json:"name"`
	Flags      string `json:"flags"`
	MTU        int    `json:"mtu"`
	Inet       string `json:"inet,omitempty"`
	Netmask    string `json:"netmask,omitempty"`
	Inet6      string `json:"inet6,omitempty"`
	Prefixlen  int    `json:"prefixlen,omitempty"`
	Scopeid    int    `json:"scopeid,omitempty"`
	Ether      string `json:"ether,omitempty"`
	Txqueuelen int    `json:"txqueuelen,omitempty"`
	RX         struct {
		Packets int `json:"packets"`
		Bytes   int `json:"bytes"`
		Errors  int `json:"errors"`
		Dropped int `json:"dropped"`
		Overrun int `json:"overrun"`
		Frame   int `json:"frame"`
	} `json:"rx,omitempty"`
	TX struct {
		Packets  int `json:"packets"`
		Bytes    int `json:"bytes"`
		Errors   int `json:"errors"`
		Dropped  int `json:"dropped"`
		Overruns int `json:"overruns"`
		Carrier  int `json:"carrier"`
		Colls    int `json:"collisions"`
	} `json:"tx,omitempty"`
}

func totalincomingtraffic() {
	/*

		Iterate over each container in the list and get the corresponding list of interfaces in that container.
		Iterate over each interface in the list and get the corresponding list of incoming traffic in that interface.
		Add up the total incoming traffic for each interface to get the total incoming traffic for that container.
		Add up the total incoming traffic for each container to get the total incoming traffic for that pod.
		Add up the total incoming traffic for each pod to get the total incoming traffic for that endpoint.
		Add up the total incoming traffic for all the endpoints to get the total incoming traffic for the entire cluster.
	*/

	// --------------------------------------------------------------------
	//Get the list of all the endpoints in the cluster using the Kubernetes client API.
	//Iterate over each endpoint in the list and get the corresponding list of pods behind that endpoint
	slice, podinfo, totalnumofendpoints := getEndpoints() // forgot to add the getEndpoints function
	fmt.Println(totalnumofendpoints)

	for _, pod := range podinfo {
		fmt.Println(pod)

		// now get container info for each pod to get the network-interface
		// for looping over the containers in the pod
		for _, container := range pod.Containers {
			fmt.Println(container)

			// getting the list of interfaces in the container
			// note this is only for a single container
			interfacelist := getInterfaces(pod.Name, pod.Namespace, container.Name)
			fmt.Println(interfacelist) // this will print the info of the network interface of one container)

		}
	}

	// Wrong ideolgy

	// Iterate over each pod in the list and get the corresponding list of containers in that pod.
	// for i := 0; i < len(podinfo); i++ {
	// 	podname := podinfo[i].Name
	// 	podnamespace := podinfo[i].Namespace
	// 	fmt.Printf("Pod Name: %s\n", podname, podnamespace)
	// 	fmt.Println("---------------------------")

	// now get container info for each pod to get the network-interface

	// for lopping over the containers in the pod
	// for j := 0; j < len(podinfo[i].Containers); j++ {
	// 	containername := podinfo[i].Containers[j].Name
	// 	fmt.Println(containername)
	// 	fmt.Println("---------------------------")

	// getting the list of interfaces in the container
	// note this is only for a single container
	// interfacelist := getInterfaces(podname, podnamespace, containername)
	// fmt.Println(interfacelist) // this will print the list of interfaces and their names

	// for looping over the interfaces in the container
	// for k := 0; k < len(interfacelist); k++ {
	// 	interfacename := interfacelist[k].Name
	// 	fmt.Printf("Interface Name: %s\n", interfacename)
	// 	fmt.Println("---------------------------")

	// 	// getting the list of incoming traffic in the interface
	// 	incomingtrafficlist := getIncomingTraffic(podname, podnamespace, containername, interfacename)
	// 	fmt.Printf("Incoming Traffic List: %s\n", incomingtrafficlist)

	// 	// for looping over the incoming traffic in the interface
	// 	for m := 0; m < len(incomingtrafficlist); m++ {
	// 		incomingtraffic := incomingtrafficlist[m].Bytes
	// 		fmt.Printf("Incoming Traffic: %s\n", incomingtraffic)
	// 		fmt.Println("---------------------------")
	// 	}
	// }
	//}

	// }
}

// how we can get the inertfaces in the container
// we can get the list of interfaces in the container by using the following command
// kubectl exec -it <podname> -n <podnamespace> -c <containername> -- /bin/bash -c "cat /proc/net/dev | awk '{print $1}' | cut -d ':' -f1 | grep -v Inter-| grep -v face | grep -v lo"
// this command will return the list of interfaces in the container
// we can aslo run  the following command to get the list of interfaces in the container
// kubectl exec -it <podname> -n <podnamespace> -c <containername> -- ifconfig
// this command will return the list of interfaces in the container
func getInterfaces(podname string, podnamespace string, containername string) string {

	// running the command to get the list of interfaces in the container
	cmd := exec.Command("kubectl", "exec", "-it", podname, "-n", podnamespace, "-c", containername, "--", "ifconfig")
	// stdout, err := cmd.StdoutPipe()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// if err := cmd.Start(); err != nil {
	// 	fmt.Println(err)
	// }

	// output, err := ioutil.ReadAll(stdout)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// if err := cmd.Wait(); err != nil {
	// 	fmt.Println(err)
	// }

	// op := (string(output))

	// return op
}

// ------------------------------------------------------------------------------------------------------------------

// func outgoingtraffic() {

// }

// func networktraffic() {

// }

// func networklatency() {

// }

// func networkerrors() {

// }

// func networkrequests() {

// }

// func exposedports() {

// }

// getting the network throughput

// getting the network packets

// getting the network connections

// getting the network connections

// getting the exposed services

// getting the exposed endpoints

// getting the exposed routes
