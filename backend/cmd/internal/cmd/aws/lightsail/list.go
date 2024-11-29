package lightsail

import (
	"context"
	"gf-vue3-admin/cmd/internal/CplatformClinet"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/spf13/cobra"
)

type InstaceInfo struct {
	Status string // 状态
	Name   string // 实例名

}

// startCmd represents the start command
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long:  `列出所有Lightsail实例.`,
	Run: func(cmd *cobra.Command, args []string) {
		// 获取区域
		region, _ := cmd.Flags().GetString("region")
		group, _ := cmd.Flags().GetString("gourp")
		//获取客户端
		cl := CplatformClinet.GetClient[*lightsail.Client](CplatformClinet.WithRegion(region), CplatformClinet.WithClientType("lightsail"))
		StartLightsail(cl, group)
	},
}

func init() {
	LgCmd.AddCommand(startCmd)
	ListCmd.Flags().StringP("region", "r", "", "区域如:us-east-1")
	//startCmd.Flags().StringP("gourp", "g", "", "请输入实例的分组如:gourpA")
}

func ListIstance(ctx context.Context, cl *lightsail.Client) error {
	// 获取所有实例列表
	output, err := cl.GetInstances(ctx, &lightsail.GetInstancesInput{})
	if err != nil {
		return err
	}
	for _, instance := range output.Instances {
		glog.Infof(ctx, "实例名%s 实例状态：%s", gconv.String(instance.Name), gconv.String(instance.State))
	}
	return nil

}
