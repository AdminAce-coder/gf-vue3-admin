package s3

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"

	//"gf-vue3-admin/cmd/internal/CplatformClinet"
	"github.com/gogf/gf/v2/os/glog"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/cobra"
)

type S3Info struct {
	S3Bucket string // 存储桶名
	Region   string
}

// startCmd represents the start command
var listCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long:  `启动lightsail.`,
	Run: func(cmd *cobra.Command, args []string) {
		// 获取区域
		//region, _ := cmd.Flags().GetString("region")
		//group, _ := cmd.Flags().GetString("gourp")
		//获取客户端
		//cl := CplatformClinet.GetClient[*s3.Client](CplatformClinet.WithRegion(region), CplatformClinet.WithClientType("lightsail"))
		//StartLightsail(cl, group)
	},
}

func init() {
	S3cmd.AddCommand(listCmd)
	listCmd.Flags().StringP("region", "r", "", "区域如:us-east-1")
	listCmd.Flags().StringP("gourp", "g", "", "请输入实例的分组如:gourpA")
}

// 列出存储桶
func ListS3Opject(ctx context.Context, cl *s3.Client) error {
	// 获取所有存储桶列表
	output, err := cl.ListBuckets(context.Background(), &s3.ListBucketsInput{})
	if err != nil {
		return err
	}
	for _, bk := range output.Buckets {
		glog.Infof(ctx, "bk.Name%s", gconv.String(bk.Name))

	}
	return nil
}
