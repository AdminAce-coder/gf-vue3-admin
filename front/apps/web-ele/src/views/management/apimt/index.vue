<template>
    <div class="user-management">
      <div class="table-container">
        <el-table :data="tableData" stripe style="width: 100%" border>
          <el-table-column label="ID" width="80">
            <template #default="scope">
              {{ scope.$index + 1 }}
            </template>
          </el-table-column>
          <el-table-column prop="path" label="API路径" width="200" />
          <el-table-column prop="method" label="请求方法" width="100" />
          <el-table-column prop="description" label="API简介" width="200" />
          <el-table-column label="标签" width="150">
            <template #default="{ row }">
              <el-tag v-for="tag in row.tags" :key="tag" size="small" style="margin-right: 4px">
                {{ tag }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column label="操作" width="200" fixed="right">
            <template #default="scope">
              <el-button type="primary" size="small" @click="handleEdit(scope.row)">
                <el-icon><Edit /></el-icon>编辑
              </el-button>
              <el-button type="danger" size="small" @click="handleDelete(scope.row)">
                <el-icon><Delete /></el-icon>删除
              </el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </div>
  </template>
  
  <script lang="ts" setup>
  import { ref, onMounted } from 'vue'
  import { getApiInfo } from '#/api/user/apiinfo'
  import type { ApiInfo } from '#/api/user/apiinfo'

  const tableData = ref<ApiInfo[]>([])

  const loadApiInfo = async () => {
    try {
      const response = await getApiInfo()
      console.log('API响应数据:', response)
      if (response.apiInfo) {
        tableData.value = response.apiInfo
      } else {
        console.error('返回数据格式不正确:', response)
      }
    } catch (error) {
      console.error('加载API信息失败:', error)
    }
  }
  // 页面加载时，获取API信息
  onMounted(() => {
    loadApiInfo()
  })

  const handleAdd = () => {
    // 处理新增用户的逻辑
  }
  
  const handleEdit = (row: any) => {
    // 处理编辑逻辑
    console.log('编辑用户:', row)
  }
  
  const handleDelete = (row: any) => {
    // 处理删除逻辑
    console.log('删除用户:', row)
  }
  </script>
  
  <style lang="scss" scoped>
  .user-management {
    padding: 20px;
    background-color: #fff;
    border-radius: 4px;
    
    .operation-bar {
      margin-bottom: 16px;
    }
    
    .table-container {
      background-color: #fff;
      border-radius: 4px;
    }
  }
  </style>
  