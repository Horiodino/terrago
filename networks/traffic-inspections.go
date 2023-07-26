// this code will we containerized and run on the all the nodes of the cluster
// and will be able to access the node network interface card and get the network info

// first we will make a connection from the node in which the code is running to the other node
// in which our main code is running and then we will get the network info from the node network interface card

package networks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"net"

	metrices "github.com/Horiodino/terrago/cluster-metrices"
)

func Nodeip() {

	for _, node := range metrices.NodeInfoList {
		fmt.Println(node.IP)
		fmt.Println(node.Name)
	}
}

func AcepptRequest(name, address string) {
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

var DATAVAR string

func handleconnection(conn net.Conn, name string) {
	// Read data from the connection
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(err)
		conn.Close()
		return
	}

	data := string(buffer[:n])

	// sort the data then write it in dataToWrite map

	dataToWrite := map[string]interface{}{
		"Recived_Packets": data,
		// "Recived_Data":    Total_Data_Rcvd ,
	}

	err = writeDataIntoJsonFile(dataToWrite)
	if err != nil {
		// Handle the error if needed
		panic(err)
	}
	conn.Close()
}

func writeDataIntoJsonFile(data map[string]interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Write the data to the file with the given filePath
	err = ioutil.WriteFile("incoming.json", jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
