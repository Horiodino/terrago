package api

import (
	"context"
	"time"
)

type AllInfo interface {
	//take time and data to get the info from the  file
	SaveClusterInfo(c context.Context, Date, Time time.Time) (string, error)
	SaveNodeInfo(c context.Context, Date, Time time.Time) (string, error)
	SavePodInfo(c context.Context, Date, Time time.Time) (string, error)
	SaveNamespaceInfo(c context.Context, Date, Time time.Time) (string, error)
}

type SavedInfo struct {
	AllInfo
	context.Context
	ClusterInfo *ClusterInfoStruct
	NodeInfo    *NodeInfoStruct
}
