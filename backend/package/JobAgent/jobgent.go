package JobAgent

import (
	"context"
	"fmt"

	"github.com/AdminAce-coder/JobAgent/jobagentclient"
	"github.com/AdminAce-coder/JobAgent/pb/jobAgent"
	"github.com/zeromicro/go-zero/zrpc"
)

func main() {
	conf := zrpc.RpcClientConf{
		Target: "1.92.75.225:8080",
	}

	client := zrpc.MustNewClient(conf)
	jobagentclient := jobagentclient.NewJobAgent(client)

	// 使用客户端
	response, err := jobagentclient.DoJob(context.Background(), &jobAgent.Request{
		// 设置参数
		Command: "free -m",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("响应:%s", response.Result)

}
