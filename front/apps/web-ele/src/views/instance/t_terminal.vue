<template>
  <div class="terminal-container">
    <!-- 终端容器 -->
    <div class="terminal-wrapper">
      <div v-if="!isConnected" class="no-connection-tip">
        暂无终端连接
      </div>
      <div ref="terminalRef" class="terminal" v-else></div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import { Terminal } from 'xterm'
import { FitAddon } from 'xterm-addon-fit'
import 'xterm/css/xterm.css'
import { ElMessage } from 'element-plus'

const terminalRef = ref(null)
let terminal = null
let fitAddon = null
let ws = null

// 添加连接状态变量
const isConnected = ref(false)

// 初始化终端
const initTerminal = () => {
  terminal = new Terminal({
    cursorBlink: true,
    fontSize: 14,
    theme: {
      background: '#1e1e1e'
    },
    convertEol: true,
    cols: 120,
    rows: 30,
    cursorStyle: 'block',
    fontFamily: 'Consolas, "Courier New", monospace',
    rendererType: 'canvas',
    disableStdin: false,
    allowProposedApi: true,
    windowsMode: false
  })

  fitAddon = new FitAddon()
  terminal.loadAddon(fitAddon)

  terminal.open(terminalRef.value)
  fitAddon.fit()

  let commandBuffer = ''

  // 修改终端输入处理
  terminal.onKey(({ key, domEvent }) => {
    if (!isConnected.value || !ws || ws.readyState !== WebSocket.OPEN) {
      return
    }

    const ev = domEvent
    const printable = !ev.altKey && !ev.ctrlKey && !ev.metaKey

    if (ev.keyCode === 13) { // Enter
      if (commandBuffer.trim()) {
        ws.send(JSON.stringify({
          type: 'message',
          data: commandBuffer
        }))
        commandBuffer = ''
      }
      terminal.write('\r\n')
    } else if (ev.keyCode === 8) { // Backspace
      if (commandBuffer.length > 0) {
        commandBuffer = commandBuffer.substring(0, commandBuffer.length - 1)
        terminal.write('\b \b')
      }
    } else if (printable) {
      commandBuffer += key
      terminal.write(key)
    }

    // 处理 Ctrl+C
    if (ev.ctrlKey && (ev.key === 'c' || ev.key === 'C')) {
      ws.send(JSON.stringify({
        type: 'key',
        key: 'ctrl+c'
      }))
      commandBuffer = ''
      terminal.write('^C\r\n')
    }
  })

  // 禁用默认的 onData 处理
  terminal.onData(() => {})
}

// 修改 WebSocket 连接函数，移除 SSH 表单相关逻辑
const connectWebSocket = (host) => {
  const wsUrl = `ws://${host}:9443/ws`
  ws = new WebSocket(wsUrl)

  ws.onopen = () => {
    console.log('WebSocket连接成功')
    isConnected.value = true
    terminal.write('\r\nWebSocket连接成功\r\n')
    ElMessage.success('终端连接成功')
  }

  ws.onclose = () => {
    console.log('WebSocket连接关闭')
    isConnected.value = false
    terminal.write('\r\nWebSocket连接已断开，正在尝试重新连接...\r\n')
    setTimeout(connectWebSocket, 3000)
  }

  ws.onerror = (error) => {
    console.error('WebSocket错误:', error)
    isConnected.value = false
  }

  ws.onmessage = (event) => {
    try {
      const data = JSON.parse(event.data)
      if (data.type === 'cmd' || data.type === 'test') {
        // 直接写入服务器返回的数据，包括回显
        terminal.write(data.data)
      }
    } catch (error) {
      console.error('处理WebSocket消息时出错:', error)
    }
  }
}

// 导出连接方法供外部调用
const connect = (host) => {
  if (!terminal) {
    initTerminal()
  }
  connectWebSocket(host)
}

// 导出断开连接方法
const disconnect = () => {
  if (ws) {
    ws.close()
  }
  isConnected.value = false
}

// 清理资源
onUnmounted(() => {
  if (terminal) {
    terminal.dispose()
  }
  if (ws) {
    ws.close()
  }
})

// 导出方法供外部使用
defineExpose({
  connect,
  disconnect
})
</script>

<style scoped>
.terminal-container {
  display: flex;
  flex-direction: column;
  height: 100%;
}

.terminal-wrapper {
  padding: 10px;
  background-color: #1e1e1e;
  border-radius: 4px;
  height: 400px;
}

.terminal {
  height: 100%;
  width: 100%;
}

.no-connection-tip {
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #888;
  font-size: 14px;
}
</style>
