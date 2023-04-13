package main

import (
	"fmt"
	"os"
	"time"
)

// this struct will define the user for terrago

type User struct {
	Username string
	Password string
}

// this structure will store the information about the resources for the user

type ClusterHistory struct {
	ClusterName     string
	ClusterID       string
	ClusterType     string
	ClusterSize     string
	Cluster_Created os.time
}

// store the VM History
type VMHistory struct {
	VMName     string
	VMID       string
	VMSize     string
	VM_Created os.time
}

// store the RG History
type RGHistory struct {
	RGName     string
	RGLOCATION string
	RG_Created os.time
}
