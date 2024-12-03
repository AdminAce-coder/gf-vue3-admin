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
    cursorStyle: 'block',    // 添加光标样式
    fontFamily: 'Consolas, "Courier New", monospace',  // 使用等宽字体
    rendererType: 'canvas'   // 使用canvas渲染
  })

  fitAddon = new FitAddon()
  terminal.loadAddon(fitAddon)

  terminal.open(terminalRef.value)
  fitAddon.fit()

  let commandBuffer = ''
  let isProcessingCommand = false
  let ctrlPressed = false

  // 添加键盘事件监听
  terminal.attachCustomKeyEventHandler((event) => {
    // 检测 Ctrl 键的按下和释放
    if (event.type === 'keydown' && event.key === 'Control') {
      ctrlPressed = true
      return true
    }
    if (event.type === 'keyup' && event.key === 'Control') {
      ctrlPressed = false
      return true
    }

    // 检测 Ctrl+C
    if (event.type === 'keydown' && ctrlPressed && (event.key === 'c' || event.key === 'C')) {
      if (ws && ws.readyState === WebSocket.OPEN) {
        ws.send(JSON.stringify({
          type: 'key',
          key: 'ctrl+c'
        }))
      }
      return false // 不阻止默认行为，让终端显示 ^C
    }

    return true
  })

  // 处理终端输入
  terminal.onData(data => {
    if (!isConnected.value) {
      terminal.write('\r\nWebSocket未连接，请等待连接成功...\r\n$ ')
      return
    }

    // 处理回车键
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
      isProcessingCommand = false
    } else if (data === '\u007f') { // 退格键
      if (commandBuffer.length > 0) {
        commandBuffer = commandBuffer.slice(0, -1)
        terminal.write('\b \b')
      }
    } else if (!isProcessingCommand && data >= ' ') { // 可打印字符
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
      switch(data.type) {
        case 'test':
          terminal.write(data.data)
          if (!data.data.endsWith('\n')) {
            terminal.write('\r\n')
          }
          break
        case 'cmd':
          // 处理命令输出，保持格式
          const output = data.data.toString()
          // 移除末尾的换行符，因为我们会自己添加
          const cleanOutput = output.replace(/\n+$/, '')
          
          // 如果输出不是以提示符结尾，则添加提示符
          if (!cleanOutput.endsWith('# ') && !cleanOutput.endsWith('$ ')) {
            terminal.write(cleanOutput)
            if (!cleanOutput.endsWith('\n')) {
              terminal.write('\r\n')
            }
          } else {
            // 如果已经有提示符，直接输出
            terminal.write(cleanOutput)
          }
          break
        default:
          terminal.write(data.data)
          if (!data.data.endsWith('\n')) {
            terminal.write('\r\n')
          }
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
