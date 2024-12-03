package terminal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"

	"github.com/gorilla/websocket"
)

// WsMsg 是 WebSocket 消息的结构体
type WsMsg struct {
	Type string `json:"type"`
	Data string `json:"data,omitempty"` // WsMsgCmd
	//Cols      int    `json:"cols,omitempty"`      // WsMsgResize
	//Rows      int    `json:"rows,omitempty"`      // WsMsgResize
	//Timestamp int `json:"timestamp,omitempty"` // WsMsgHeartbeat
}

// LogicSshWsSession 是逻辑 SSH WebSocket 会话的结构体
type SshWsSession struct {
	//stdinPipe io.WriteCloser  // 标准输入管道
	//session *ssh.Session    // SSH 会话特
	wsConn *websocket.Conn // WebSocket 连接
	sshcon *SshConn
}

func NewSshWsSession(sshcon *SshConn, wsConn *websocket.Conn) *SshWsSession {
	return &SshWsSession{sshcon: sshcon, wsConn: wsConn}
}

// 关闭WSeision
func (s *SshWsSession) Close() {
	s.sshcon.Close()
}
func (s *SshWsSession) Start(quitchan chan bool) {
	go s.receiveWsMsg(quitchan)
	go s.SendComboOutput(quitchan)
}

// 接收WebSocket 消息

func (s *SshWsSession) receiveWsMsg(exitCh chan bool) {
	ctx := gctx.New()
	wscon := s.wsConn
	for {
		select {
		case <-exitCh:
			return
		default:
			_, data, err := wscon.ReadMessage()
			if err != nil {
				glog.Error(ctx, "读取WebSocket消息错误:", err)
				return
			}
			glog.Info(ctx, "正在解析消息")
			rmgs := WsMsg{}
			if err := json.Unmarshal(data, &rmgs); err != nil {
				glog.Error(ctx, "解析消息失败:", err)
				continue
			}

			// 过滤空命令
			command := strings.TrimSpace(rmgs.Data)
			if command == "" {
				continue
			}

			glog.Infof(ctx, "发送的命令是：%s", command)
			// 将字符串转换为字节切片，确保命令以换行符结束
			byteMsg := []byte(command + "\n")
			if err := s.SendmgsToPipe(byteMsg); err != nil {
				glog.Error(ctx, "发送命令到SSH管道失败:", err)
				continue
			}
		}
	}
}

// 发送 WebSocket 输入命令到 SSH 会话的标准输入管道
func (s *SshWsSession) SendmgsToPipe(cmdBytes []byte) error {
	fmt.Println("正在 输入命令到 SSH 会话的标准输入管道:", string(cmdBytes))
	_, err := s.sshcon.StdinPipe.Write(cmdBytes)
	return err
}

// 发送 WebSocket 消息
func (s *SshWsSession) SendComboOutput(exitCh chan bool) {
	ctx := gctx.New()
	wscon := s.wsConn
	// 创建一个定时器，减少检查间隔以提高响应速度
	ticker := time.NewTicker(time.Millisecond * time.Duration(30))
	defer ticker.Stop()

	var buffer bytes.Buffer
	for {
		select {
		case <-ticker.C:
			if s.sshcon.ComboOutput == nil {
				return
			}
			// 获取输入
			input := s.sshcon.ComboOutput.Bytes()
			if len(input) > 0 {
				// 将新数据追加到缓冲区
				buffer.Write(input)

				// 检查是否有完整的输出行
				if bytes.Contains(input, []byte{'\n'}) || len(buffer.Bytes()) > 1024 {
					output := buffer.String()
					wsdata, err := json.Marshal(WsMsg{
						Type: "cmd",
						Data: output,
					})
					if err != nil {
						glog.Error(ctx, "序列化消息失败:", err)
						continue
					}

					// 发送数据
					if err = wscon.WriteMessage(websocket.TextMessage, wsdata); err != nil {
						glog.Error(ctx, "发送数据失败:", err)
					}

					// 清空缓冲区
					buffer.Reset()
					s.sshcon.ComboOutput.Buffer.Reset()
				}
			}
		case <-exitCh:
			return
		}
	}
}
