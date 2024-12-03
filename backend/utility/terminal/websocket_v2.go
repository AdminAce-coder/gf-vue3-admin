package terminal

import (
	"context"
	"fmt"
	"gf-vue3-admin/internal/model/utiliy"
	"net/http"

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
	ctx := r.Context()
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		glog.Error(ctx, "升级到WebSocket失败:", err)
		return
	}
	defer ws.Close()

	// 获取查询参数
	host := r.URL.Query().Get("host")
	port := r.URL.Query().Get("port")
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")
	cols := r.URL.Query().Get("cols")
	if cols == "" {
		cols = "80"
	}
	rows := r.URL.Query().Get("rows")
	if rows == "" {
		rows = "24"
	}

	if host == "" || port == "" || username == "" || password == "" {
		glog.Error(ctx, "缺少必要的SSH连接参数")
		ws.Close()
		return
	}

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
		glog.Error(ctx, "创建SSH连接失败:", err)
		return
	}

	// 创建WebSocket会话
	SshWsSession := NewSshWsSession(sshConn, ws)
	defer SshWsSession.Close()

	// 启动会话
	var cancelCtx func()
	ctx, cancelCtx = context.WithCancel(ctx)
	defer cancelCtx()
	SshWsSession.Start(ctx)

	fmt.Println("SshWsSession 启动了")

	// 简单的消息循环：读取消息并回显给客户端
	for {
		messageType, msg, err := ws.ReadMessage()
		if err != nil {
			fmt.Println("Read Error:", err)
			break
		}
		fmt.Printf("Received: %s\n", string(msg))

		// 回显消息
		err = ws.WriteMessage(messageType, msg)
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
