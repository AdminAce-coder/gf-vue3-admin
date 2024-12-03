package terminal

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"

	"github.com/gorilla/websocket"
)

// WsMsg 是 WebSocket 消息的结构体
type WsMsg struct {
	Type string `json:"type"`
	Data string `json:"data"`
	Key  string `json:"key,omitempty"`
}

// LogicSshWsSession 是逻辑 SSH WebSocket 会话的结构体
type SshWsSession struct {
	wsConn *websocket.Conn
	sshcon *SshConn
	exitCh chan bool
}

func NewSshWsSession(sshcon *SshConn, wsConn *websocket.Conn) *SshWsSession {
	return &SshWsSession{sshcon: sshcon, wsConn: wsConn}
}

// 关闭WSeision
func (s *SshWsSession) Close() {
	s.sshcon.Close()
}

func (s *SshWsSession) Start(ctx context.Context) {
	glog.Info(ctx, "SshWsSession 启动了")
	s.exitCh = make(chan bool)
	go s.SendComboOutput(s.exitCh)

	for {
		_, data, err := s.wsConn.ReadMessage()
		if err != nil {
			glog.Error(ctx, "读取WebSocket消息失败:", err)
			s.exitCh <- true
			return
		}

		var wsMsg WsMsg
		if err := json.Unmarshal(data, &wsMsg); err != nil {
			glog.Error(ctx, "解析WebSocket消息失败:", err)
			continue
		}

		if err := s.handleWsMsg(ctx, wsMsg); err != nil {
			glog.Error(ctx, "处理WebSocket消息失败:", err)
			continue
		}
	}
}

// handleWsMsg 处理WebSocket消息
func (s *SshWsSession) handleWsMsg(ctx context.Context, wsMsg WsMsg) error {
	switch wsMsg.Type {
	case "message":
		// 添加回车换行符
		data := wsMsg.Data + "\r\n"
		if err := s.SendmgsToPipe([]byte(data)); err != nil {
			return fmt.Errorf("写入SSH失败: %v", err)
		}
	case "key":
		if wsMsg.Key == "ctrl+c" {
			if err := s.SendmgsToPipe([]byte{3}); err != nil {
				return fmt.Errorf("发送Ctrl+C失败: %v", err)
			}
		}
	}
	return nil
}

// SendmgsToPipe 发送消息到SSH管道
func (s *SshWsSession) SendmgsToPipe(msg []byte) error {
	_, err := s.sshcon.StdinPipe.Write(msg)
	return err
}

// SendComboOutput 发送组合输出
func (s *SshWsSession) SendComboOutput(exitCh chan bool) {
	ctx := gctx.New()
	wscon := s.wsConn
	ticker := time.NewTicker(time.Millisecond * 10) // 减少检查间隔到10ms
	defer ticker.Stop()

	for {
		select {
		case <-exitCh:
			return
		case <-ticker.C:
			input := s.sshcon.ComboOutput.Bytes()
			if len(input) > 0 {
				// 发送输出到WebSocket
				err := wscon.WriteJSON(WsMsg{
					Type: "cmd",
					Data: string(input),
				})
				if err != nil {
					glog.Error(ctx, "发送WebSocket消息失败:", err)
					return
				}
				// 立即清空缓冲区
				s.sshcon.ComboOutput.Reset()
			}
		}
	}
}
