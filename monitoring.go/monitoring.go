package monitoring

//in this file we will define the monitoring part of the cluster and all its components
// and we will give it to the webserver to display it

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// defining the monitoring struct
// in this struct we will define all the monitoring components
type Monitoring struct {
	Name   string
	Type   string
	Status string
}

// as we are using the Azure SDK for Go, we also need to fetcht the info regarding the cluster
// its components and the monitoring components so we will import the cluster.go file
