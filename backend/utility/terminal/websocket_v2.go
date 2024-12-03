package terminal

import (
	"fmt"
	"gf-vue3-admin/internal/model/utiliy"
	"net/http"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/gogf/gf/os/glog"
	"github.com/gorilla/websocket"
)

// 升级器配置：处理 HTTP 升级为 WebSocket 的握手
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 允许所有来源连接（生产环境建议增加安全检查）
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WebSocket 连接处理函数
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// 升级 HTTP 连接到 WebSocket
	ctx := gctx.New()
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		glog.Error(ctx, "WebSocket升级失败:", err)
		return
	}
	defer conn.Close()
	// 创建SSH连接
	ssh := Sshconfig{
		Userinfo: utiliy.SshUser,
	}
	sshClient, err := ssh.NewSshConfig(ctx)
	if err != nil {
		glog.Error(ctx, "SSH客户端配置失败:", err)
		return
	}
	sshConn, err := sshClient.NewSshConn(2048, 2048)
	if err != nil {
		glog.Error(ctx, "创建SSH连接失1败:", err)
		return
	}
	SshWsSession := NewSshWsSession(sshConn, conn)
	done := make(chan bool)
	SshWsSession.Start(done)

	fmt.Println("SshWsSession 启动了")

	// 简单的消息循环：读取消息并回显给客户端
	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Read Error:", err)
			break
		}
		fmt.Printf("Received: %s\n", string(msg))

		// 回显消息
		err = conn.WriteMessage(messageType, msg)
		if err != nil {
			fmt.Println("Write Error:", err)
			break
		}
	}
}

func StartwebScoket() {

	// 设置 WebSocket 路由
	http.HandleFunc("/ws", handleWebSocket)
	// 启动 HTTP 服务
	fmt.Println("WebSocket server started at :9443")
	err := http.ListenAndServe(":9443", nil)
	if err != nil {
		fmt.Println("Server Error:", err)
	}
}
