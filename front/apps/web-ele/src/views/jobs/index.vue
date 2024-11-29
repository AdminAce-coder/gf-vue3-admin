<template>
  <div class="ansible-container">
    <!-- 任务列表卡片 -->
    <el-card class="task-list">
      <template #header>
        <div class="card-header">
          <span>Ansible 任务列表</span>
          <el-button type="primary" @click="openTaskForm">新建任务</el-button>
        </div>
      </template>

      <el-table :data="taskList" style="width: 100%">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="name" label="任务名称" />
        <el-table-column prop="playbook" label="Playbook" />
        <el-table-column prop="status" label="状态">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row?.status)">{{ row?.status || '未知' }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="createTime" label="创建时间" />
        <el-table-column label="操作" width="200">
          <template #default="{ row }">
            <el-button link type="primary" @click="viewTaskDetail(row)">查看</el-button>
            <el-button link type="danger" @click="deleteTask(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <!-- 新建任务抽屉 -->
    <el-drawer
      v-model="drawerVisible"
      :title="drawerType === 'create' ? '新建任务' : '任务详情'"
      direction="rtl"
      size="500px"
    >
      <el-form
        ref="taskFormRef"
        :model="taskForm"
        :rules="taskRules"
        label-width="100px"
      >
        <el-form-item label="任务名称" prop="name">
          <el-input v-model="taskForm.name" placeholder="请输入任务名称" />
        </el-form-item>

        <el-form-item label="任务类型" prop="type">
          <el-select v-model="taskForm.type" placeholder="选择任务类型" @change="handleTypeChange">
            <el-option label="单次命令" value="command" />
            <el-option label="Playbook" value="playbook" />
          </el-select>
        </el-form-item>

        <el-form-item
          v-if="taskForm.type === 'command'"
          label="执行命令"
          prop="command"
        >
          <el-input
            v-model="taskForm.command"
            type="textarea"
            rows="2"
            placeholder="请输入要执行的命令"
          />
        </el-form-item>

        <el-form-item
          v-if="taskForm.type === 'playbook'"
          label="Playbook"
          prop="playbook"
        >
          <el-select v-model="taskForm.playbook" placeholder="选择 Playbook">
            <el-option
              v-for="item in playbookOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>

        <el-form-item label="主机组" prop="hostGroup">
          <el-select v-model="taskForm.hostGroup" multiple placeholder="选择主机组">
            <el-option
              v-for="item in hostGroupOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>

        <el-form-item label="变量" prop="vars">
          <el-input
            v-model="taskForm.vars"
            type="textarea"
            rows="4"
            placeholder="请输入 YAML 格式的变量"
          />
        </el-form-item>
      </el-form>

      <template #footer>
        <div style="flex: auto">
          <el-button @click="drawerVisible = false">取消</el-button>
          <el-button
            v-if="drawerType === 'create'"
            type="primary"
            @click="submitTask"
          >
            提交
          </el-button>
        </div>
      </template>
    </el-drawer>

    <!-- 在template最后添加执行进度弹窗 -->
    <el-drawer
      v-model="progressVisible"
      direction="btt"
      size="400px"
      :with-header="false"
      destroy-on-close
    >
      <div class="terminal-container">
        <div class="terminal-header">
          <h3>终端输出</h3>
          <div class="terminal-controls">
            <el-tag :type="executionStatus.type">{{ executionStatus.text }}</el-tag>
            <el-button
              type="danger"
              :disabled="executionStatus.text === '执行完成'"
              @click="stopExecution"
            >
              终止执行
            </el-button>
            <el-button @click="progressVisible = false">关闭</el-button>
          </div>
        </div>
        <div ref="terminalElement" class="terminal-content"></div>
      </div>
    </el-drawer>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, onUnmounted, nextTick } from 'vue'
import { Terminal } from 'xterm'
import { FitAddon } from 'xterm-addon-fit'
import { WebLinksAddon } from 'xterm-addon-web-links'
import 'xterm/css/xterm.css'
import { ElMessage, ElMessageBox } from 'element-plus'

// 任务列表数据
const taskList = ref([
  {
    id: 1,
    name: '系统更新',
    playbook: 'update.yml',
    status: 'success',
    createTime: '2024-03-20 10:00:00'
  },
  // ... 更多数据
])

// 抽屉表单相关
const drawerVisible = ref(false)
const drawerType = ref('create') // create 或 view
const taskFormRef = ref(null)

// 表单数据
const taskForm = reactive({
  name: '',
  type: '',
  command: '',
  playbook: '',
  hostGroup: [],
  vars: ''
})

// 表单验证规则
const taskRules = {
  name: [{ required: true, message: '请输入任务名称', trigger: 'blur' }],
  type: [{ required: true, message: '请选择任务类型', trigger: 'change' }],
  command: [{ required: true, message: '请输入执行命令', trigger: 'blur' }],
  playbook: [{ required: true, message: '请选择 Playbook', trigger: 'change' }],
  hostGroup: [{ required: true, message: '请选择主机组', trigger: 'change' }]
}

// 选项数据
const playbookOptions = [
  { label: '系统更新', value: 'update.yml' },
  { label: '安装 Docker', value: 'docker.yml' }
]

const hostGroupOptions = [
  { label: 'web服务器', value: 'web' },
  { label: '数据库服务器', value: 'db' }
]

// 状态标签类型
const getStatusType = (status) => {
  if (!status) return 'info'

  const types = {
    success: 'success',
    running: 'warning',
    failed: 'danger',
    pending: 'info'
  }
  return types[status] || 'info'
}

// 打开新建任务表单
const openTaskForm = () => {
  drawerType.value = 'create'
  drawerVisible.value = true
}

// 查看任务详情
const viewTaskDetail = (task) => {
  drawerType.value = 'view'
  Object.assign(taskForm, task)
  drawerVisible.value = true
}

// 删除任务
const deleteTask = async (task) => {
  try {
    await ElMessageBox.confirm('确认删除该任务？', '提示', {
      type: 'warning'
    })
    // TODO: 调用删除 API
    ElMessage.success('删除成功')
  } catch (error) {
    // 用户取消删除
  }
}

// 进度窗口相关状态
const progressVisible = ref(false)
const terminalElement = ref(null)
let terminal = null
let fitAddon = null
let ws = null

// 初始化终端
const initTerminal = () => {
  if (!terminalElement.value) return

  terminal = new Terminal({
    cursorBlink: true,
    theme: {
      background: '#1e1e1e',
      foreground: '#ffffff'
    },
    fontSize: 14,
    fontFamily: 'Consolas,Liberation Mono,Menlo,Courier,monospace',
    rendererType: 'canvas',
    convertEol: true,
    scrollback: 800,
  })

  // 添加插件
  fitAddon = new FitAddon()
  terminal.loadAddon(fitAddon)
  terminal.loadAddon(new WebLinksAddon())

  // 挂载终端
  terminal.open(terminalElement.value)
  fitAddon.fit()

  // 监听窗口大小变化
  window.addEventListener('resize', () => {
    fitAddon?.fit()
  })
}

// 连接WebSocket
const connectWebSocket = () => {
  // 根据您的后端WebSocket地址修改
  ws = new WebSocket('ws://your-backend-url/ws/terminal')

  ws.onopen = () => {
    terminal.writeln('Connected to terminal...')
  }

  ws.onmessage = (event) => {
    terminal.write(event.data)
  }

  ws.onclose = () => {
    terminal.writeln('\r\nConnection closed.')
  }

  ws.onerror = (error) => {
    terminal.writeln(`\r\nError: ${error.message}`)
  }
}

// 修改提交任务函数
const submitTask = async () => {
  if (!taskFormRef.value) return

  try {
    await taskFormRef.value.validate()

    progressVisible.value = true
    executionStatus.text = '执行中'
    executionStatus.type = 'primary'

    // 初始化终端
    await nextTick()
    initTerminal()
    connectWebSocket()

    // 发送任务信息到后端
    ws?.send(JSON.stringify({
      type: 'task',
      data: {
        type: taskForm.type,
        command: taskForm.command,
        playbook: taskForm.playbook,
        hostGroup: taskForm.hostGroup
      }
    }))

    drawerVisible.value = false
  } catch (error) {
    console.error('表单验证失败:', error)
  }
}

// 终止执行
const stopExecution = () => {
  ws?.send(JSON.stringify({ type: 'stop' }))
  executionStatus.text = '已终止'
  executionStatus.type = 'danger'
}

// 组件卸载时清理
onUnmounted(() => {
  ws?.close()
  terminal?.dispose()
  window.removeEventListener('resize', fitAddon?.fit)
})

// 处理类型变更
const handleTypeChange = () => {
  // 清空相关字段
  taskForm.command = ''
  taskForm.playbook = ''
}
</script>

<style scoped>
.ansible-container {
  padding: 20px;
}

.task-list {
  margin-bottom: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.terminal-container {
  height: 100%;
  display: flex;
  flex-direction: column;
  padding: 20px;
}

.terminal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.terminal-controls {
  display: flex;
  gap: 10px;
  align-items: center;
}

.terminal-content {
  flex: 1;
  background: #1e1e1e;
  border-radius: 4px;
  padding: 4px;
  overflow: hidden;
}

:deep(.xterm) {
  height: 100%;
  padding: 8px;
}
</style>
