package network

import (
	"fmt"

	"net"
)

// how its going to work ?
// we will use hostnetwork : true  so we will containerize the code and run it on the host network which will
// be node network. so we will be able to access the node network from the container.
// and we will make a requeste to another pod that will be acessing the node network-info and display it on the browser.

// for now lets code the network part for getting all the possible network info from the node network interface card

func AcepptRequests(name, address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleconnection(conn, name)
	}
}

// this handelconnection will handle the connection that is being made to the node
// it will recieve the data and then process it and then close the connection
// this is the function that will be called when the connection is made to the node
func handleConnection(conn net.Conn, name string) {
	// Read data from the connection
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(err)
		conn.Close()
		return
	}

	data := string(buffer[:n])
	fmt.Printf("[%s] Received data: %s\n", name, data)

	conn.Close()
}

// this func will be invoked if user want more details then it will send the request to the node to get more packet details
// the data string will be constant always
func GetMoreDetails(address, data string) {

	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	_, err = conn.Write([]byte(data))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Data sent to", address)

}
