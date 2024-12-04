<template>
  <div class="host-table">
    <div class="table-header">
      <el-button type="primary" size="large" @click="showAddDialog">
        新增主机
      </el-button>
      <el-button type="primary" size="large" :icon="Refresh" @click="refreshTable">
        刷新
      </el-button>
      <TInstanceDrop ref="instanceDropRef" />
    </div>

    <el-table :data="sshInfoList" style="width: 100%" border>
      <el-table-column prop="hostname" label="主机名称" />
      <el-table-column prop="host" label="IP地址" />
      <el-table-column prop="port" label="端口" width="100" />
      <el-table-column prop="user" label="用户名" />
      <el-table-column prop="status" label="状态" width="100">
        <template #default="{ row }">
          <el-tag :type="row.status === '在线' ? 'success' : 'danger'">
            {{ row.status }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="200" fixed="right">
        <template #default="{ row }">
          <el-button type="primary" size="small" @click="handleConnect(row)">
            连接
          </el-button>
          <el-button type="danger" size="small" @click="handleDelete(row)">
            删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>
    <TTerminal ref="terminalRef" />
  </div>
</template>

<script setup lang="ts">
import { ref, nextTick } from 'vue'
import { getSSHConnectionInfo as fetchSSHInfo } from '#/api/systemctl/instance'
import { Refresh } from '@element-plus/icons-vue'
import TInstanceDrop from './t_instance_drop.vue'
// 测试连接
import { testSSHConnection } from '#/api/systemctl/instance'
// 连接SSH
import { connectSSH } from '#/api/systemctl/instance'
import { ElMessage } from 'element-plus'
import TTerminal from './t_terminal.vue'  // 导入终端组件
import { useRouter } from 'vue-router'  // 添加这行

const sshInfoList = ref([])
const instanceDropRef = ref(null)
const terminalRef = ref(null)
const router = useRouter()

// 获取SSH连接信息
const getSSHConnectionInfo = async () => {
  try {
    const res = await fetchSSHInfo()
    console.log('完整响应:', res)
    sshInfoList.value = res.ssh_info_list
  } catch (error) {
    console.error('获取SSH连接信息失败:', error)
  }
}

// 页面加载时获取数据
getSSHConnectionInfo()

const handleConnect = async (row: any) => {
  try {
    // 首先检查终端组件是否已挂载
    if (!terminalRef.value) {
      ElMessage.error('终端组件未准备就绪')
      return
    }

    // 第一步：测试SSH连接
    const testResult = await testSSHConnection({
      host: row.host
    })

    if (!testResult) {
      ElMessage.error('连接测试失败：服务器无响应')
      return
    }

    if (testResult.code === 0) {
      // 第二步：建立SSH连接
      const connectResult = await connectSSH({
        host: row.host,
      })

      if (!connectResult || connectResult.code === 0) {
        // 连接成功后跳转到终端页面
        router.push({
          name: 'terminal',  // 替换成您的终端页面路由名称
          params: {
            host: row.host
          }
        })
        ElMessage.success('SSH连接成功')
      } else {
        ElMessage.error(`SSH连接失败: ${connectResult?.message || '未知错误'}`)
      }
    } else {
      ElMessage.error(`SSH连接测试失败: ${testResult.message || '未知错误'}`)
    }

  } catch (error) {
    console.error('连接过程出错:', error)
    ElMessage.error('连接失败：' + (error.message || '未知错误'))
  }
}

const handleDelete = (row: any) => {
  console.log('删除主机:', row)
}

// 刷新表格数据
const refreshTable = () => {
  getSSHConnectionInfo()
}

const showAddDialog = () => {
  instanceDropRef.value?.showDialog()
}
</script>

<style scoped>
.host-table {
  height: 100%;
}

:deep(.el-table) {
  font-size: 14px;
}

:deep(.el-button--small) {
  padding: 8px 15px;
  margin: 0 5px;
}

.table-header {
  margin-bottom: 16px;
  text-align: right;
}
</style>
