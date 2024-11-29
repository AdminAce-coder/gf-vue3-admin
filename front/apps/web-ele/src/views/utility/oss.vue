<template>
  <div class="storage-bucket">
    <el-card class="bucket-card">
      <template #header>
        <div class="card-header">
          <h2>存储桶管理</h2>
          <el-button type="primary" @click="createBucket">
            <el-icon><Plus /></el-icon>新建存储桶
          </el-button>
        </div>
      </template>

      <!-- 存储桶列表 -->
      <el-table :data="bucketList" stripe style="width: 100%">
        <el-table-column prop="name" label="存储桶名称" />
        <el-table-column prop="region" label="所在地域" />
        <el-table-column prop="access" label="访问权限" />
        <el-table-column prop="createTime" label="创建时间" />
        <el-table-column label="操作" width="250">
          <template #default="scope">
            <el-button-group>
              <el-button type="primary" @click="viewDetails(scope.row)">详情</el-button>
              <el-button type="warning" @click="editBucket(scope.row)">设置</el-button>
              <el-button type="danger" @click="deleteBucket(scope.row)">删除</el-button>
            </el-button-group>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { Plus } from '@element-plus/icons-vue'

// 存储桶列表数据
const bucketList = ref([
  {
    name: 'test-bucket',
    region: '华东-上海',
    access: '私有',
    createTime: '2024-03-20'
  }
])

// 新建存储桶
const createBucket = () => {
  // TODO: 实现新建存储桶逻辑
}

// 查看详情
const viewDetails = (row) => {
  // TODO: 实现查看详情逻辑
}

// 编辑存储桶
const editBucket = (row) => {
  // TODO: 实现编辑存储桶逻辑
}

// 删除存储桶
const deleteBucket = (row) => {
  // TODO: 实现删除存储桶逻辑
}
</script>

<style scoped>
.storage-bucket {
  padding: 20px;
}

.bucket-card {
  margin-bottom: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.card-header h2 {
  margin: 0;
}
</style>
