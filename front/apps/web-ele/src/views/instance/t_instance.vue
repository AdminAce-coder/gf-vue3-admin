<template>
  <div class="host-table">
    <el-table :data="tableData" style="width: 100%" border>
      <el-table-column prop="name" label="主机名称" />
      <el-table-column prop="ip" label="IP地址" />
      <el-table-column prop="port" label="端口" width="100" />
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
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// 模拟数据
const tableData = ref([
  {
    name: '测试服务器1',
    ip: '192.168.1.100',
    port: 22,
    status: '在线'
  },
  {
    name: '测试服务器2',
    ip: '192.168.1.101',
    port: 22,
    status: '离线'
  }
])

const handleConnect = (row: any) => {
  console.log('连接主机:', row)
}

const handleDelete = (row: any) => {
  console.log('删除主机:', row)
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
</style>
