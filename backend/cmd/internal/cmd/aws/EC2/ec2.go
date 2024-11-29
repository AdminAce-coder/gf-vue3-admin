package EC2

import (
	"context"
	"encoding/base64"
	"fmt"
	"gf-vue3-admin/cmd"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/gogf/gf/v2/os/gfile"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/spf13/cobra"
)

type Ec2 struct {
	ImageId      string
	InstanceType string
}

// lgCmd represents the lg command
var Ec2cmd = &cobra.Command{
	Use:   "k8s",
	Short: "A brief description of your command",
	Long:  `EC2相关操作`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("lg called")
	},
}

func init() {
	cmd.RootCmd.AddCommand(Ec2cmd)
}

func CreatrEc2(ctx context.Context, ec2Client *ec2.Client) error {
	// 读取文件
	userData := gfile.GetContents("./start.sh")
	// 转为BASE64
	userData = base64.StdEncoding.EncodeToString([]byte(userData))
	fmt.Printf("userData:%s", userData)
	_, err := ec2Client.RunInstances(context.Background(), &ec2.RunInstancesInput{
		ImageId:      aws.String("ami-04a81a99f5ec58529"),
		InstanceType: types.InstanceTypeT2Micro,
		MinCount:     aws.Int32(1),
		MaxCount:     aws.Int32(1),
		UserData:     aws.String(userData),
	})
	if err != nil {
		return err
		fmt.Println("创建ec2实例失败", err)
	}
	return nil
	// ... handle error ...
}
