package network

// this code will we containerized and run on the all the nodes of the cluster
// and will be able to access the node network interface card and get the network info

// first we will make a connection from the node in which the code is running to the other node
// in which our main code is running and then we will get the network info from the node network interface card

type NetworkInfo struct {
	Total_Incoming_Packets int
	Total_Data_Rcvd        int
	Total_Data_Sent        int
	Node_Name              string
}

var NetworkInfoList []NetworkInfo

func RecivedInfo() {

	// to recive the info from the other node we will use the grpc
	// and we will use the grpc to send the info to the other node

}

func MakeConnection() {

	// this will make a connection to the other node
	// and will send the info to the other node

}
