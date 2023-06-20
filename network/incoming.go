package network

import (
	"fmt"
	"log"

	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

// how its going to work ?
// we will use hostnetwork : true  so we will containerize the code and run it on the host network which will
// be node network. so we will be able to access the node network from the container.
// and we will make a requeste to another pod that will be acessing the node network-info and display it on the browser.

// for now lets code the network part for getting all the possible network info from the node network interface card

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
		fmt.Println(Total_Incoming_Packets)
		fmt.Println(packet)
		Total_Data_Rcvd += packet.Metadata().CaptureInfo.Length
		fmt.Println("Total Data Rcvd:", Total_Data_Rcvd/1024, "KB")

	}
	fmt.Println("Total packets arrived:", Total_Incoming_Packets)
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
		fmt.Println(packet)
		Total_Data_Sent += packet.Metadata().CaptureInfo.Length
		fmt.Println("Total Data Sent:", Total_Data_Sent/1024, "KB")

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

	// now implemetn the logic for deep packet inspection
	// first get the packet and then get the data from the packet and then check the data for any malicious content

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
		fmt.Printf("Sender :: %s  :: Reciver  ", packet.TransportLayer().TransportFlow())
		fmt.Println()

		// TransportLayer.LayerContents() will give the data that is being sent

		fmt.Printf("Protocol :: %s \n", packet.TransportLayer().LayerType())

		fmt.Println("Data Content :: ", packet.TransportLayer().LayerContents())

		fmt.Println("Data Payload :: ", packet.TransportLayer().LayerPayload())

		fmt.Println("---------------------------------------------------------------")
		fmt.Println(packet.ApplicationLayer().LayerContents())
		// fmt.Println(packet.ApplicationLayer().LayerPayload())
		// fmt.Println(packet.ApplicationLayer().LayerType())
		// fmt.Println(packet.ApplicationLayer().Payload())
		// fmt.Println(packet.Data())
		// fmt.Println(packet.Metadata().CaptureInfo.Timestamp)
		// fmt.Println(packet.Metadata().CaptureInfo.Length)
		// fmt.Println(packet.Metadata().CaptureInfo.InterfaceIndex)
		// fmt.Printf("Error :: %s", packet.ErrorLayer())
		break
	}
}
