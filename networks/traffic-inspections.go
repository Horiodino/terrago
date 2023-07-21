// this code will we containerized and run on the all the nodes of the cluster
// and will be able to access the node network interface card and get the network info

// first we will make a connection from the node in which the code is running to the other node
// in which our main code is running and then we will get the network info from the node network interface card

package networks

import (
	"fmt"
	"log"

	"net"

	metrices "github.com/Horiodino/terrago/monitoring"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

type NetworkInfo struct {
	Total_Incoming_Packets int
	Total_Outgoing_Packets int
	Total_Data_Rcvd        int
	Total_Data_Sent        int
}

var NetworkInfoList []NetworkInfo

type PacketInfo struct {
	Packet        string	'json:"packet"'
	SourceIP      string	'json:"sourceip"'
	DestinationIP string 'json:"destinationip"'
	SourcePort    string	'json:"sourceport"'
	Destination   string	'json:"destinationport"'
	Protocol      string	'json:"protocol"'
	Data          string	'json:"data"'
	Data_Payload  string	'json:"data_payload"'
}

var PacketInfoList []PacketInfo

func Nic_info() {
	Nicinfo, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(Nicinfo)

}

var (
	Total_Data_Rcvd = 0
	Total_Data_Sent = 0
)

// per node
func IncomingTraffic() {

	iface := "wlo1"
	handle, err := pcap.OpenLive(iface, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	filter := "inbound"
	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatal(err)
	}
	Total_Incoming_Packets := 0
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packetSource.Packets() {
		// Process each captured packet here
		Total_Incoming_Packets++
		Total_Data_Rcvd += packet.Metadata().CaptureInfo.Length

		NetworkInfo := NetworkInfo{
			Total_Incoming_Packets: Total_Incoming_Packets,
			Total_Data_Rcvd:        Total_Data_Rcvd / 1024,
		}

		NetworkInfoList = append(NetworkInfoList, NetworkInfo)
	}
}

func Outbound_Traffic() {

	iface := "wlo1"
	handle, err := pcap.OpenLive(iface, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}

	filter := "outbound"
	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatal(err)
	}

	Total_Outgoing_Packets := 0

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packetSource.Packets() {
		// Process each captured packet here
		Total_Outgoing_Packets++
		fmt.Println(Total_Outgoing_Packets)
		Total_Data_Sent += packet.Metadata().CaptureInfo.Length
		fmt.Println("Total Data Sent:", Total_Data_Sent/1024, "KB")

		NetworkInfo := NetworkInfo{
			Total_Outgoing_Packets: Total_Outgoing_Packets,
			Total_Data_Sent:        Total_Data_Sent / 1024,
		}

		NetworkInfoList = append(NetworkInfoList, NetworkInfo)
	}

	fmt.Println("Total packets arrived:", Total_Outgoing_Packets)

}

func DeepPacketInspection() {
	// here we will get all the info regardin the packet like the source and destination ip and port and the protocol used
	// and also the data that is being sent and recieved.

	iface := "wlo1"
	handle, err := pcap.OpenLive(iface, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}

	filter := "inbound"

	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatal(err)
	}

	// get the packet sorce info who is sending the packet and who is recieving the packet

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packetSource.Packets() {
		// fmt.Println(packet)
		// this packet.NetworkLayer().NetworkFlow() will give the source and destination ip and port
		fmt.Printf("Sender :: %s  :: Reciver", packet.NetworkLayer().NetworkFlow())
		fmt.Println()
		// TransportLayer.TransportFlow() will give the source and destination port
		fmt.Printf("Port Address Sender :: %s  :: Reciver  ", packet.TransportLayer().TransportFlow())
		fmt.Println()

		// TransportLayer.LayerContents() will give the data that is being sent

		fmt.Printf("Protocol :: %s \n", packet.TransportLayer().LayerType())

		fmt.Println("Data Content :: ", packet.TransportLayer().LayerContents())

		fmt.Println("Data Payload :: ", packet.TransportLayer().LayerPayload())

		fmt.Println("---------------------------------------------------------------")

		// fmt.Println(packet.NetworkLayer().LayerContents())
		// fmt.Println(packet.Layers())

		// fmt.Println(packet.ApplicationLayer().LayerContents())
		// fmt.Println(packet.ApplicationLayer().LayerPayload())
		// fmt.Println(packet.ApplicationLayer().LayerType())
		// fmt.Println(packet.ApplicationLayer().Payload())
		// fmt.Println(packet.Data())
		fmt.Println(packet.Metadata().CaptureInfo.Timestamp)

		// this captureinfo.length will give the length of the packet in bytes
		fmt.Println("Packet Length ", packet.Metadata().CaptureInfo.Length, "bytes")
		// this captureinfo.interfaceindex will give the interface index of the network card that is being used to send the packet
		fmt.Println("Interface Index ", packet.Metadata().CaptureInfo.InterfaceIndex)
		fmt.Printf("Error :: %s", packet.ErrorLayer())

		fmt.Println("---------------------------------------------------------------")

		// now we will se the network layer info
		fmt.Println(packet.NetworkLayer().LayerContents())
		fmt.Println(packet.NetworkLayer().LayerPayload())
		fmt.Println(packet.NetworkLayer().LayerType())
		fmt.Println(packet.NetworkLayer().NetworkFlow())
		// fmt.Println(packet.NetworkLayer().LayerType().Contains())
		// fmt.Println(packet.NetworkLayer().LayerType().Decode())

		break
	}
}

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

	if data == "packetinfo" {
		DeepPacketInspection()
	}

	fmt.Printf("[%s] Received data: %s\n", name, data)

	conn.Close()
}

func SendNicInfo(address, data string) {

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