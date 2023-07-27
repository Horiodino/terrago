package cmd

import (
	"fmt"
	"net/http"
	"os/exec"

	metrices "github.com/Horiodino/terrago/cluster-metrices"
	"github.com/Horiodino/terrago/web"
	"github.com/spf13/cobra"
)

var WebCmd = &cobra.Command{
	Use:   "dashboard",
	Short: "Print cluster information",
	Run: func(cmd *cobra.Command, args []string) {
		web.StartWebServer(metrices.Monitoring{})
	},
}
