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

// 添加一个变量来存储当前输入的内容
let currentLine = ''

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

  // 添加提示符
  terminal.write('\r\n$ ')

  // 处理终端输入
  terminal.onData(data => {
    if (ws && ws.readyState === WebSocket.OPEN) {
      // 显示用户输入
      terminal.write(data)

      // 如果是回车键，发送消息并添加新的提示符
      if (data === '\r') {
        ws.send(JSON.stringify({
          type: 'message',
          data: currentLine
        }))
        currentLine = '' // 清空当前行
        terminal.write('\n$ ')
      } else {
        // 收集用户输入
        currentLine += data
      }
    }
  })
}

// 连接SSH
const connectSSH = () => {
  if (!sshForm.value.host || !sshForm.value.username || !sshForm.value.password) {
    ElMessage.error('请填写完整的连接信息')
    return
  }

  showTerminal.value = true

  // 延迟一帧等待 DOM 更新
  setTimeout(() => {
    if (!terminal) {
      initTerminal()
    }

    // 关闭已存在的连接
    if (ws) {
      ws.close()
      ws = null
    }

    try {
      const wsUrl = `ws://1.92.75.225:9443/ws `
      ws = new WebSocket(wsUrl)

      ws.onopen = () => {
        console.log('WebSocket连接成功')
        terminal.write('WebSocket连接成功...\r\n')
        // 发送测试消息
        ws.send(JSON.stringify({
          type: 'test',
          data: 'Hello WebSocket Server!'
        }))
      }

      ws.onmessage = (event) => {
        try {
          const data = JSON.parse(event.data)
          // 根据消息类型处理
          switch(data.type) {
            case 'test':
              // 测试消息可以简单显示
              terminal.write(`\r\n${data.data}\r\n`)
              break
            case 'cmd':
              // 命令输出直接写入，不需要显示json格式
              terminal.write(data.data)
              break
            default:
              terminal.write(`\r\n${data.data}\r\n`)
          }
        } catch (e) {
          console.error('解析消息错误:', e)
          terminal.write('\r\n消息解析错误\r\n')
        }
      }

      ws.onerror = (error) => {
        console.error('WebSocket错误:', error)
        terminal.write('\r\n连接错误！\r\n')
        terminal.write('正在尝试重新连接...\r\n')

        // 3秒后尝试重新连接
        setTimeout(() => {
          connectSSH()
        }, 3000)
      }

      ws.onclose = () => {
        terminal.write('\r\n连接已关闭\r\n')
      }
    } catch (error) {
      console.error('创建WebSocket失败:', error)
      terminal.write('\r\n创建连接失败！\r\n')
      ElMessage.error('创建WebSocket连接失败')
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
