package api

import (
	"context"
	"time"
)

type SavedInfo struct {
	context.Context
	ClusterInfo *ClusterInfoStruct
	NodeInfo    *NodeInfoStruct
}

func SaveClusterInfo(c context.Context, Date, Time time.Time, Object string) (string, error) {
	return "ClusterInfo", nil
}
