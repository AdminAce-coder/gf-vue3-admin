package s3

import (
	"fmt"
	"gf-vue3-admin/cmd"
	"github.com/spf13/cobra"
)

// lgCmd represents the lg command
var S3cmd = &cobra.Command{
	Use:   "S3",
	Short: "A brief description of your command",
	Long:  `AWS S3 服务相关命令`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("lg called")
	},
}

func init() {
	cmd.RootCmd.AddCommand(S3cmd)
}
