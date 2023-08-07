package metrices

import (
	"github.com/Horiodino/terrago/api"
)

func GET() {
	K8sclient := api.NewClient()

	K8sclient.PullClusterInfo()

}
