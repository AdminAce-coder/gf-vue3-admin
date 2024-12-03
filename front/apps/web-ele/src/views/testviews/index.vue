<template>
  <div class="terminal-container">
    <!-- 添加连接表单 -->
    <el-form :model="sshForm" class="ssh-form">
      <el-form-item label="主机地址">
        <el-input v-model="sshForm.host" placeholder="例如：192.168.1.100"></el-input>
      </el-form-item>
      <el-form-item label="端口">
        <el-input v-model="sshForm.port" placeholder="例如：22"></el-input>
      </el-form-item>
      <el-form-item label="用户名">
        <el-input v-model="sshForm.username" placeholder="例如：root"></el-input>
      </el-form-item>
      <el-form-item label="密码">
        <el-input v-model="sshForm.password" type="password"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="connectSSH">连接服务器</el-button>
      </el-form-item>
    </el-form>

    <!-- 终端容器 -->
    <div v-show="showTerminal" class="terminal-wrapper">
      <div ref="terminalRef" class="terminal"></div>
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
const showTerminal = ref(false)
let terminal = null
let fitAddon = null
let ws = null

const sshForm = ref({
  host: '',
  port: '22',
  username: '',
  password: ''
})

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
      // 不在这里写入换行，让服务器的回显来处理
      if (commandBuffer.trim()) {
        ws.send(JSON.stringify({
          type: 'message',
          data: commandBuffer + '\n'
        }))
      }
      commandBuffer = ''
    } else if (ev.keyCode === 8) { // Backspace
      if (commandBuffer.length > 0) {
        commandBuffer = commandBuffer.substring(0, commandBuffer.length - 1)
        // 发送退格命令到服务器
        ws.send(JSON.stringify({
          type: 'message',
          data: '\b'
        }))
      }
    } else if (printable) {
      commandBuffer += key
      // 直接发送字符到服务器，让服务器处理回显
      ws.send(JSON.stringify({
        type: 'message',
        data: key
      }))
    }

    // 处理 Ctrl+C
    if (ev.ctrlKey && (ev.key === 'c' || ev.key === 'C')) {
      ws.send(JSON.stringify({
        type: 'key',
        key: 'ctrl+c'
      }))
      commandBuffer = ''
    }
  })

  // 禁用默认的 onData 处理
  terminal.onData(() => {})
}

// 修改 WebSocket 消息处理
const connectWebSocket = () => {
  const wsUrl = 'ws://1.92.75.225:9443/ws'
  ws = new WebSocket(wsUrl)

  ws.onopen = () => {
    console.log('WebSocket连接成功')
    isConnected.value = true
    terminal.write('\r\nWebSocket连接成功\r\n')
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

// 连接SSH
const connectSSH = () => {
  if (!sshForm.value.host || !sshForm.value.username || !sshForm.value.password) {
    ElMessage.error('请填写完整的连接信息')
    return
  }

  showTerminal.value = true

  setTimeout(() => {
    if (!terminal) {
      initTerminal()
    }
    connectWebSocket()
  })
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
</script>

<style scoped>
.terminal-container {
  display: flex;
  flex-direction: column;
  height: 100%;
}

.ssh-form {
  margin-bottom: 20px;
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
</style>
