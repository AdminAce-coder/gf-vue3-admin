<template>
  <div class="k8s-container">
    <!-- 集群列表 -->
    <el-card class="cluster-list">
      <template #header>
        <div class="card-header">
          <span>集群列表</span>
          <el-button type="primary" @click="showCreateDialog">创建集群</el-button>
        </div>
      </template>

      <el-table :data="clusterList" stripe>
        <el-table-column prop="name" label="集群名称" />
        <el-table-column prop="version" label="K8S版本" />
        <el-table-column prop="nodeCount" label="节点数量" />
        <el-table-column prop="status" label="状态">
          <template #default="{ row }">
            <el-tag :type="row.status === 'Running' ? 'success' : 'warning'">
              {{ row.status }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="200">
          <template #default="{ row }">
            <el-button type="text" @click="handleDetail(row)">详情</el-button>
            <el-button type="text" @click="handleDelete(row)" class="delete-btn">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <!-- 创建集群对话框 -->
    <el-drawer
      v-model="drawerVisible"
      title="创建集群"
      size="600px"
      :destroy-on-close="true"
    >
      <el-form
        ref="formRef"
        :model="formData"
        :rules="rules"
        label-width="120px"
        style="padding: 0 20px"
      >
        <el-form-item label="集群名称" prop="name">
          <el-input v-model="formData.name" placeholder="请输入集群名称" />
        </el-form-item>
        <el-form-item label="K8S版本" prop="version">
          <el-select v-model="formData.version" placeholder="请选择K8S版本" style="width: 100%">
            <el-option label="v1.24" value="1.24" />
            <el-option label="v1.23" value="1.23" />
            <el-option label="v1.22" value="1.22" />
          </el-select>
        </el-form-item>
        <el-form-item label="节点配置" prop="nodeConfig">
          <div class="node-config">
            <el-input-number v-model="formData.nodeConfig.count" :min="1" :max="10" label="节点数量" />
            <el-select v-model="formData.nodeConfig.instanceType" placeholder="选择实例类型">
              <el-option label="2核4G" value="2c4g" />
              <el-option label="4核8G" value="4c8g" />
              <el-option label="8核16G" value="8c16g" />
            </el-select>
          </div>
        </el-form-item>
      </el-form>

      <template #footer>
        <div class="drawer-footer">
          <el-button @click="drawerVisible = false">取消</el-button>
          <el-button type="primary" @click="handleCreate">确定</el-button>
        </div>
      </template>
    </el-drawer>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'

// 集群列表数据
const clusterList = ref([])

// 表单数据
const drawerVisible = ref(false)
const formRef = ref(null)
const formData = reactive({
  name: '',
  version: '',
  nodeConfig: {
    count: 1,
    instanceType: ''
  }
})

// 表单验证规则
const rules = {
  name: [
    { required: true, message: '请输入集群名称', trigger: 'blur' },
    { min: 3, max: 20, message: '长度在 3 到 20 个字符', trigger: 'blur' }
  ],
  version: [
    { required: true, message: '请选择K8S版本', trigger: 'change' }
  ]
}

// 显示创建对话框
const showCreateDialog = () => {
  drawerVisible.value = true
}

// 创建集群
const handleCreate = async () => {
  if (!formRef.value) return

  await formRef.value.validate((valid) => {
    if (valid) {
      // TODO: 调用创建集群API
      ElMessage.success('创建集群成功')
      drawerVisible.value = false
    }
  })
}

// 查看集群详情
const handleDetail = (row) => {
  // TODO: 跳转到集群详情页面
  console.log('查看集群详情:', row)
}

// 删除集群
const handleDelete = (row) => {
  ElMessageBox.confirm(
    '确认删除该集群吗？',
    '警告',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }
  ).then(() => {
    // TODO: 调用删除集群API
    ElMessage.success('删除成功')
  })
}
</script>

<style scoped>
.k8s-container {
  padding: 20px;
}

.cluster-list {
  margin-bottom: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.delete-btn {
  color: #F56C6C;
}

.node-config {
  display: flex;
  gap: 10px;
}

.node-config .el-select {
  flex: 1;
}

.drawer-footer {
  display: flex;
  justify-content: flex-end;
  padding: 20px;
  border-top: 1px solid #dcdfe6;
}
</style>
