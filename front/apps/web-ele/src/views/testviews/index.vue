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
    cols: 100,
    rows: 24
  })

  fitAddon = new FitAddon()
  terminal.loadAddon(fitAddon)

  terminal.open(terminalRef.value)
  fitAddon.fit()

  terminal.write('\r\n$ ')

  let commandBuffer = ''
  let isProcessingCommand = false

  terminal.onData(data => {
    if (!isConnected.value) {
      terminal.write('\r\nWebSocket未连接，请等待连接成功...\r\n$ ')
      return
    }

    if (data === '\r') {
      isProcessingCommand = true
      const command = commandBuffer.trim()
      if (command) {
        console.log('Sending command:', command)
        ws.send(JSON.stringify({
          type: 'message',
          data: command
        }))
      }
      commandBuffer = ''
      terminal.write('\r\n')
      isProcessingCommand = false
    } else if (data === '\u007f') { 
      if (commandBuffer.length > 0) {
        commandBuffer = commandBuffer.slice(0, -1)
        terminal.write('\b \b')
      }
    } else if (!isProcessingCommand && data >= ' ') { 
      commandBuffer += data
      terminal.write(data)
    }
  })
}

// 连接WebSocket
const connectWebSocket = () => {
  const wsUrl = 'ws://1.92.75.225:9443/ws'
  ws = new WebSocket(wsUrl)

  ws.onopen = () => {
    console.log('WebSocket连接成功')
    isConnected.value = true
    terminal.write('WebSocket连接成功...\r\n$ ')
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
      switch(data.type) {
        case 'test':
          terminal.write(`${data.data}`)
          if (!data.data.endsWith('\n')) {
            terminal.write('\r\n')
          }
          terminal.write('$ ')
          break
        case 'cmd':
          const output = data.data.toString()
          const cleanOutput = output.replace(/\n+$/, '')
          terminal.write(cleanOutput)
          if (!cleanOutput.endsWith('\n')) {
            terminal.write('\r\n')
          }
          terminal.write('$ ')
          break
        default:
          terminal.write(data.data)
          if (!data.data.endsWith('\n')) {
            terminal.write('\r\n')
          }
          terminal.write('$ ')
      }
    } catch (e) {
      console.error('解析消息错误:', e)
      terminal.write('\r\n消息解析错误\r\n$ ')
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

    if (ws) {
      ws.close()
      ws = null
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
