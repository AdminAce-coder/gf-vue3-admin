<template>
  <div class="user-management">
    <div class="table-container">
      <el-table :data="tableData" stripe style="width: 100%" border>
        <el-table-column prop="avatar" label="头像" width="80" />
        <el-table-column prop="id" label="ID" width="50" />
        <el-table-column prop="username" label="用户名" width="100" />
        <el-table-column prop="nickname" label="昵称" width="100" />
        <el-table-column prop="phone" label="手机号" width="120" />
        <el-table-column prop="email" label="邮箱" width="180" />
        <el-table-column prop="role" label="用户角色" width="120" />
        <el-table-column prop="status" label="启用" width="80" />
        <el-table-column label="操作" width="260" fixed="right">
          <template #default="scope">
            <el-button type="primary" size="small" @click="handleEdit(scope.row)">
              <el-icon><Edit /></el-icon>编辑
            </el-button>
            <el-button type="warning" size="small" @click="handlePassword(scope.row)">
              <el-icon><Key /></el-icon>修改密码
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

const tableData = [
  {
    avatar: 'https://example.com/avatar.jpg',
    id: '001',
    username: 'zhangsan',
    nickname: '张三',
    phone: '13800138000',
    email: 'zhangsan@example.com',
    role: '管理员',
    status: true
  }
]

const handleAdd = () => {
  // 处理新增用户的逻辑
}

const handleEdit = (row: any) => {
  // 处理编辑逻辑
  console.log('编辑用户:', row)
}

const handlePassword = (row: any) => {
  // 处理修改密码逻辑
  console.log('修改密码:', row)
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
