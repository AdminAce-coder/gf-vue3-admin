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

// 初始化终端
const initTerminal = () => {
  terminal = new Terminal({
    cursorBlink: true,
    fontSize: 14,
    theme: {
      background: '#1e1e1e'
    }
  })

  fitAddon = new FitAddon()
  terminal.loadAddon(fitAddon)

  terminal.open(terminalRef.value)
  fitAddon.fit()

  // 处理终端输入
  terminal.onData(data => {
    if (ws && ws.readyState === WebSocket.OPEN) {
      ws.send(JSON.stringify({ type: 'input', data }))
    }
  })
}

// 连接SSH
const connectSSH = () => {
  showTerminal.value = true

  // 延迟一帧等待 DOM 更新
  setTimeout(() => {
    if (!terminal) {
      initTerminal()
    }

    // 连接WebSocket
    ws = new WebSocket(`ws://your-backend-url/ssh`)

    ws.onopen = () => {
      // 发送连接信息
      ws.send(JSON.stringify({
        type: 'connect',
        data: {
          host: sshForm.value.host,
          port: parseInt(sshForm.value.port),
          username: sshForm.value.username,
          password: sshForm.value.password
        }
      }))
    }

    ws.onmessage = (event) => {
      const data = JSON.parse(event.data)
      if (data.type === 'terminal') {
        terminal.write(data.data)
      }
    }

    ws.onerror = (error) => {
      console.error('WebSocket错误:', error)
      terminal.write('\r\n连接错误！\r\n')
    }

    ws.onclose = () => {
      terminal.write('\r\n连接已关闭\r\n')
    }
  }, 0)
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
