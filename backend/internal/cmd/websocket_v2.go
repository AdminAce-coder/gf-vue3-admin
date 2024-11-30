package cmd

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
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
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("WebSocket Upgrade Error:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Client Connected")

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
	fmt.Println("WebSocket server started at :8443")
	err := http.ListenAndServe(":8443", nil)
	if err != nil {
		fmt.Println("Server Error:", err)
	}
}
