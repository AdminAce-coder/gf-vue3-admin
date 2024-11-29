package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync/atomic"

	"github.com/olahol/melody"
)

var idCounter atomic.Int64 // 用于生成唯一的ID

func StartWebsocket() {
	m := melody.New()

	// 配置消息大小限制
	m.Config.MaxMessageSize = 1024 * 1024 * 10 // 10MB
	m.Config.MessageBufferSize = 1024 * 1024   // 1MB

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		// 添加CORS头信息
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// 处理 OPTIONS 预检请求
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		m.HandleRequest(w, r)
	})

	m.HandleConnect(func(s *melody.Session) {
		id := idCounter.Add(1) //这是

		s.Set("id", id)

		s.Write([]byte(fmt.Sprintf("iam %d", id)))
	})

	m.HandleDisconnect(func(s *melody.Session) {
		if id, ok := s.Get("id"); ok {
			m.BroadcastOthers([]byte(fmt.Sprintf("dis %d", id)), s)
		}
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		fmt.Printf("收到消息: %s\n", string(msg))

		var message struct {
			Type string      `json:"type"`
			Data interface{} `json:"data"`
		}

		if err := json.Unmarshal(msg, &message); err != nil {
			fmt.Printf("解析消息失败: %v\n", err)
			s.Write([]byte(fmt.Sprintf("错误: 无效的消息格式")))
			return
		}

		switch message.Type {
		case "connect":
			fmt.Printf("收到连接请求: %+v\n", message.Data)
			// 处理连接请求
		case "input":
			fmt.Printf("收到输入: %+v\n", message.Data)
			// 处理输入
		default:
			fmt.Printf("未知消息类型: %s\n", message.Type)
		}

		if _, ok := s.Get("id"); ok {
			m.BroadcastOthers(msg, s)
		}
	})

	fmt.Printf("WebSocket服务启动在 :6000 端口...\n")
	if err := http.ListenAndServe(":6000", nil); err != nil {
		fmt.Printf("服务启动失败: %v\n", err)
	}
}
