package lightsail

import (
	"fmt"
	"gf-vue3-admin/cmd"
	"github.com/spf13/cobra"
)

// lgCmd represents the lg command
var LgCmd = &cobra.Command{
	Use:   "lg",
	Short: "A brief description of your command",
	Long:  `AWS Lightsail 服务相关命令`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("lg called")
	},
}

func init() {
	cmd.RootCmd.AddCommand(LgCmd)
}
