package k8s

import (
	"fmt"
	"gf-vue3-admin/cmd"
	"github.com/spf13/cobra"
)

// lgCmd represents the lg command
var K8scmd = &cobra.Command{
	Use:   "k8s",
	Short: "A brief description of your command",
	Long:  `在云平台部署K8S`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("lg called")
	},
}

func init() {
	cmd.RootCmd.AddCommand(K8scmd)
}
