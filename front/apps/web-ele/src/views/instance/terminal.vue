<template>
  <div class="terminal-container">
    <div class="tab-buttons">
      <el-button-group>
        <el-button
          v-for="tab in tabs"
          :key="tab.name"
          :type="activeTab === tab.name ? 'primary' : 'default'"
          @click="activeTab = tab.name">
          {{ tab.label }}
        </el-button>
      </el-button-group>
    </div>

    <div class="content-area">
      <div class="content-window">
        <template v-if="activeTab === 'host'">
          <div class="host-header">
          </div>
          <host-table />
        </template>

        <div v-if="activeTab === 'quick-command'" class="quick-command-content">
          快速命令内容
        </div>
        <div v-if="activeTab === 'terminal'" class="terminal-content">
          <Terminal />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import HostTable from './t_instance.vue'
import Terminal from './t_terminal.vue'

const activeTab = ref('host')
const tabs = [
  { label: '主机', name: 'host' },
  { label: '快速命令', name: 'quick-command' },
  { label: '终端', name: 'terminal' }
]
</script>

<style scoped>
.terminal-container {
  height: 100%;
  display: flex;
  flex-direction: column;
  background-color: #ffffff;
  font-size: 16px;
}

.tab-buttons {
  padding: 20px;
  background-color: #ffffff;
  border-bottom: 1px solid #ebeef5;
  display: flex;
  justify-content: flex-start;
}

.content-area {
  flex: 1;
  background-color: #f5f7fa;
  padding: 20px;
  overflow: hidden;
}

.content-window {
  background-color: #ffffff;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  height: calc(100vh - 180px);
  padding: 24px;
}

:deep(.el-button) {
  padding: 12px 24px;
  font-weight: 500;
  font-size: 14px;
  transition: all 0.3s;
}

:deep(.el-button--primary) {
  background-color: #409eff;
  border-color: #409eff;
  color: #ffffff;
}

:deep(.el-button--default) {
  background-color: #ffffff;
  border-color: #dcdfe6;
  color: #606266;
}

:deep(.el-button--default:hover) {
  color: #409eff;
  border-color: #409eff;
  background-color: #ecf5ff;
}

:deep(.el-button-group .el-button) {
  border-radius: 0;
}

:deep(.el-button-group .el-button:first-child) {
  border-radius: 4px 0 0 4px;
}

:deep(.el-button-group .el-button:last-child) {
  border-radius: 0 4px 4px 0;
}

.terminal-content,
.host-content,
.quick-command-content {
  height: 100%;
  font-size: 16px;
}

.host-header {
  margin-bottom: 16px;
  display: flex;
  justify-content: flex-end;
}
</style>
