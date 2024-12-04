<!-- 新增主机弹窗 -->
<template>
  <el-dialog
    title="新增主机"
    v-model="dialogVisible"
    width="500px"
    @close="handleClose"
  >
    <el-form :model="form" :rules="rules" ref="formRef" label-width="100px">
      <el-form-item label="主机名称" prop="hostname">
        <el-input v-model="form.hostname" placeholder="请输入主机名称"></el-input>
      </el-form-item>
      <el-form-item label="主机地址" prop="addr">
        <el-input v-model="form.addr" placeholder="请输入主机IP地址"></el-input>
      </el-form-item>
      <el-form-item label="端口" prop="port">
        <el-input-number v-model="form.port" :min="1" :max="65535" placeholder="请输入端口号"></el-input-number>
      </el-form-item>
      <el-form-item label="用户名" prop="user">
        <el-input v-model="form.user" placeholder="请输入用户名"></el-input>
      </el-form-item>
      <el-form-item label="密码" prop="password">
        <el-input v-model="form.password" type="password" placeholder="请输入密码"></el-input>
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="handleClose">取消</el-button>
        <el-button type="primary" @click="handleSubmit">确定</el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { createSSHConnection } from '#/api/systemctl/instance'
import { ElMessage } from 'element-plus'
const dialogVisible = ref(false)
const formRef = ref(null)

const form = reactive({
  hostname: '',
  addr: '',
  port: 22,
  user: '',
  password: ''
})

const rules = {
  hostname: [{ required: true, message: '请输入主机名称', trigger: 'blur' }],
  addr: [{ required: true, message: '请输入主机地址', trigger: 'blur' }],
  port: [{ required: true, message: '请输入端口号', trigger: 'blur' }],
  user: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
  password: [{ required: true, message: '请输入密码', trigger: 'blur' }]
}

const handleClose = () => {
  dialogVisible.value = false
  formRef.value?.resetFields()
}

const handleSubmit = async () => {
  if (!formRef.value) return
  await formRef.value.validate(async (valid) => {
    if (valid) {
      try {
        const res = await createSSHConnection(form)
        ElMessage.success('主机添加成功')
        handleClose()
      } catch (error) {
        console.error('连接失败：', error)
        ElMessage.error('主机添加失败：' + error.message)
      }
    }
  })
}

// 导出方法供父组件调用
defineExpose({
  showDialog: () => {
    dialogVisible.value = true
  }
})
</script>
