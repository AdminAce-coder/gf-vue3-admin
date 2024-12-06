<template>
  <div class="user-management">
    <div class="operation-bar">
      <el-dropdown @command="handleAddCommand">
        <el-button type="primary">
          <el-icon><Plus /></el-icon>
          新增
          <el-icon class="el-icon--right"><ArrowDown /></el-icon>
        </el-button>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item command="api">添加API</el-dropdown-item>
            <el-dropdown-item command="group">添加API分组</el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
      
      <el-button type="danger" @click="deleteGroupDrawer.visible = true" style="margin-left: 10px">
        <el-icon><Delete /></el-icon>
        删除分组
      </el-button>
    </div>

    <div class="table-container">
      <el-table :data="tableData.slice((currentPage-1)*pageSize, currentPage*pageSize)" stripe style="width: 100%" border @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" />
        <el-table-column label="ID" width="80">
          <template #default="scope">
            {{ scope.$index + 1 }}
          </template>
        </el-table-column>
        <el-table-column prop="path" label="API路径" width="200" />
        <el-table-column prop="method" label="请求方法" width="100">
          <template #default="{ row }">
            <el-tag 
              :type="getMethodTagType(row.method)"
            >
              {{ row.method.toUpperCase() }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="description" label="API简介" width="200" />
        <el-table-column 
          prop="apiGroup" 
          label="分组" 
          sortable 
          :sort-method="sortByGroup"
          show-overflow-tooltip
        >
          <template #default="{ row }">
            <el-tag v-for="tag in row.tags" :key="tag" class="mx-1">
              {{ tag }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="300" fixed="right">
          <template #default="scope">
            <el-button type="primary" size="small" @click="handleEdit(scope.row)">
              <el-icon><Edit /></el-icon>编辑
            </el-button>
            <el-button type="success" size="small" @click="handleTest(scope.row)">
              <el-icon><VideoPlay /></el-icon>测试
            </el-button>
            <el-button type="danger" size="small" @click="handleDelete(scope.row)">
              <el-icon><Delete /></el-icon>删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      
      <div class="pagination-container">
        <el-pagination
          v-model:current-page="currentPage"
          v-model:page-size="pageSize"
          :page-sizes="[10, 20, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>
    </div>

    <el-drawer
      v-model="drawer.visible"
      :title="drawer.title"
      size="800px"
    >
      <el-form :model="formData" label-width="100px">
        <el-form-item label="API名称" required>
          <el-input v-model="formData.apiname" placeholder="请输入API名称" />
        </el-form-item>
        <el-form-item label="API版本">
          <el-input v-model="formData.apiversion" placeholder="请输入API版本" />
        </el-form-item>
        <el-form-item label="请求方法" required>
          <el-select v-model="formData.method" placeholder="请选择请求方法">
            <el-option label="GET-查询" value="GET" />
            <el-option label="POST-创建" value="POST" />
            <el-option label="PUT-更新" value="PUT" />
            <el-option label="DELETE-删除" value="DELETE" />
          </el-select>
        </el-form-item>
        <el-form-item label="API分组">
          <el-select
            v-model="formData.apiGroup"
            filterable
            allow-create
            default-first-option
            :reserve-keyword="false"
            placeholder="请选择或输入API分组"
          >
            <el-option
              v-for="tag in existingTags"
              :key="tag"
              :label="tag"
              :value="tag"
            />
          </el-select>
        </el-form-item>
        <!-- 新增请求参数配置 -->
        <template v-if="formData.method === 'POST'">
          <el-divider content-position="left">请求参数配置</el-divider>
          <div v-for="(param, index) in formData.parameters" :key="index" class="param-item">
            <el-row :gutter="10">
              <el-col :span="6">
                <el-form-item :label="index === 0 ? '参数名称' : ''" label-width="100px">
                  <el-input v-model="param.parametername" placeholder="字段名" />
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item :label="index === 0 ? '数据类型' : ''" label-width="80px">
                  <el-select 
                    v-model="param.datatype" 
                    placeholder="类型"
                    style="width: 100%"
                  >
                    <el-option-group label="基础类型">
                      <el-option label="string" value="string" />
                      <el-option label="int" value="int" />
                      <el-option label="int64" value="int64" />
                      <el-option label="float64" value="float64" />
                      <el-option label="bool" value="bool" />
                    </el-option-group>
                    <el-option-group label="复合类型">
                      <el-option label="[]string" value="[]string" />
                      <el-option label="[]int" value="[]int" />
                      <el-option label="[]int64" value="[]int64" />
                      <el-option label="map[string]string" value="map[string]string" />
                      <el-option label="map[string]interface{}" value="map[string]interface{}" />
                      <el-option label="interface{}" value="interface{}" />
                    </el-option-group>
                    <el-option-group label="时间类型">
                      <el-option label="time.Time" value="time.Time" />
                    </el-option-group>
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="4">
                <el-form-item :label="index === 0 ? '必须' : ''" label-width="60px">
                  <el-select v-model="param.required" placeholder="是否必须">
                    <el-option label="是" :value="true" />
                    <el-option label="否" :value="false" />
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item :label="index === 0 ? '描述' : ''" label-width="60px">
                  <el-input v-model="param.description" placeholder="参数描述" />
                </el-form-item>
              </el-col>
              <el-col :span="2" class="param-actions">
                <el-button type="danger" circle @click="removeParameter(index)">
                  <el-icon><Delete /></el-icon>
                </el-button>
              </el-col>
            </el-row>
          </div>
          <el-button type="primary" plain @click="addParameter" style="margin-left: 100px">
            <el-icon><Plus /></el-icon>添加参数
          </el-button>
        </template>
        <el-form-item label="API简介">
          <el-input
            v-model="formData.description"
            type="textarea"
            placeholder="请输入API简介"
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submitForm">确认</el-button>
          <el-button @click="drawer.visible = false">取消</el-button>
        </el-form-item>
      </el-form>
    </el-drawer>

    <el-drawer
      v-model="groupDrawer.visible"
      :title="groupDrawer.title"
      size="800px"
    >
      <el-form :model="groupFormData" label-width="100px">
        <el-form-item label="API路径" required>
          <el-input v-model="groupFormData.apipath" placeholder="请输入API路径，例如: /api/v1/apitest" />
        </el-form-item>
        <el-form-item label="分组名称" required>
          <el-input v-model="groupFormData.register.groupname" placeholder="请输入分组名称" />
        </el-form-item>
        <el-form-item label="是否鉴权">
          <el-switch
            v-model="groupFormData.register.needauth"
            active-text="是"
            inactive-text="否"
          />
        </el-form-item>
        <el-form-item label="是否启用">
          <el-switch
            v-model="groupFormData.register.enable"
            active-text="是"
            inactive-text="否"
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submitGroupForm">确认</el-button>
          <el-button @click="groupDrawer.visible = false">取消</el-button>
        </el-form-item>
      </el-form>
    </el-drawer>

    <el-drawer
      v-model="deleteGroupDrawer.visible"
      :title="deleteGroupDrawer.title"
      size="500px"
    >
      <el-form :model="deleteGroupFormData" label-width="100px">
        <el-form-item label="分组名称" required>
          <el-select
            v-model="deleteGroupFormData.groupname"
            placeholder="请选择要删除的分组"
            style="width: 100%"
          >
            <el-option
              v-for="tag in existingTags"
              :key="tag"
              :label="tag"
              :value="tag"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="API版本" required>
          <el-input v-model="deleteGroupFormData.version" placeholder="请输入API版本，例如v1" />
        </el-form-item>
        <el-form-item>
          <el-button type="danger" @click="submitDeleteGroup">确认删除</el-button>
          <el-button @click="deleteGroupDrawer.visible = false">取消</el-button>
        </el-form-item>
      </el-form>
    </el-drawer>

    <el-drawer
      v-model="testDrawer.visible"
      :title="testDrawer.title"
      size="600px"
    >
      <el-form :model="testFormData" label-width="100px">
        <el-form-item label="请求方法">
          <el-tag :type="getMethodTagType(testFormData.method)">
            {{ testFormData.method.toUpperCase() }}
          </el-tag>
        </el-form-item>
        <el-form-item label="请求路径">
          <el-input v-model="testFormData.path" disabled />
        </el-form-item>
        <el-form-item label="使用鉴权">
          <el-switch
            v-model="testFormData.useAuth"
            active-text="是"
            inactive-text="否"
          />
        </el-form-item>
        <el-form-item label="请求参数" v-if="testFormData.method !== 'GET'">
          <el-input
            v-model="testFormData.requestBody"
            type="textarea"
            :rows="8"
            placeholder="请输入JSON格式的请求参数"
          />
        </el-form-item>
        <el-divider>响应结果</el-divider>
        <el-form-item>
          <pre class="response-box">{{ testFormData.response }}</pre>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="executeTest">执行测试</el-button>
          <el-button @click="testDrawer.visible = false">关闭</el-button>
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getApiInfo,createApiInfo, createApiGroup ,deleteApiGroup,deleteApi} from '#/api/systemctl/apiinfo'
import type { ApiInfo } from '#/api/systemctl/apiinfo'
import { Plus, Delete, ArrowDown, VideoPlay } from '@element-plus/icons-vue'

const tableData = ref<ApiInfo[]>([])

const drawer = ref({
  visible: false,
  title: '新增API'
})

const formData = ref({
  apiname: '',
  method: '',
  apiGroup: '',
  apiversion: '',
  description: '',
  parameters: [] as {
    parametername: string;
    datatype: string;
    required: boolean;
    description: string;
  }[]
})

const selectedRows = ref([])

const existingTags = ref<string[]>([])

const groupDrawer = ref({
  visible: false,
  title: '新增API分组'
})

const groupFormData = ref({
  apipath: '',
  register: {
    needauth: true,
    groupname: '',
    enable: true
  }
})

const currentPage = ref(1)
const pageSize = ref(10)
const total = ref(0)

const loadApiInfo = async () => {
  try {
    const response = await getApiInfo()
    console.log('API响应数据:', response)
    if (response.apiInfo) {
      tableData.value = response.apiInfo
      total.value = response.apiInfo.length
      const allTags = new Set<string>()
      response.apiInfo.forEach((api: ApiInfo) => {
        if (Array.isArray(api.tags)) {
          api.tags.forEach(tag => allTags.add(tag))
        }
      })
      existingTags.value = Array.from(allTags)
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

const handleSelectionChange = (selection: any[]) => {
  selectedRows.value = selection
}

// 新增API
const handleAddCommand = (command: string) => {
  if (command === 'api') {
    drawer.value.visible = true
    drawer.value.title = '新增API'
    formData.value = {
      apiname: '',
      method: '',
      apiGroup: '',
      apiversion: '',
      description: '',
      parameters: []
    }
  } else if (command === 'group') {
    groupDrawer.value.visible = true
    groupFormData.value = {
      apipath: '',
      register: {
        needauth: true,
        groupname: '',
        enable: true
      }
    }
  } else if (command === 'deleteGroup') {
    deleteGroupDrawer.value.visible = true
    deleteGroupFormData.value = {
      groupname: '',
      version: ''
    }
  }
}
//提交新增API

const submitForm = async () => {
  try {
    const submitData = {
      apiname: formData.value.apiname,
      method: formData.value.method.toLowerCase(),
      apiGroup: formData.value.apiGroup,
      description: formData.value.description,
      apiversion: formData.value.apiversion,
      parameters: formData.value.parameters
    }

    console.log('提交的数据:', submitData)
    await createApiInfo(submitData)
    ElMessage({
      message: '创建成功',
      type: 'success'
    })
    drawer.value.visible = false
    await loadApiInfo()
  } catch (error: any) {
    ElMessage({
      message: error.message || '创建API失败',
      type: 'error'
    })
  }
}

const handleEdit = (row: any) => {
  console.log('编辑API:', row)
}

const handleDelete = async (row: any) => {
  try {
    // 从行数据中提取API路径和分组
    const deleteData = {
      apipath: row.path,  // API路径
      apigroup: Array.isArray(row.tags) ? row.tags[0] : ''  // 取第一个tag作为apigroup
    }

    await ElMessageBox.confirm(
      `确定要删除该API吗？`,
      '警告',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      }
    )

    await deleteApi(deleteData)
    ElMessage({
      type: 'success',
      message: '删除成功',
    })
    await loadApiInfo()  // 重新加载表格数据
  } catch (error: any) {
    if (error !== 'cancel') {
      ElMessage({
        type: 'error',
        message: error.message || '删除失败',
      })
    }
  }
}

const addParameter = () => {
  formData.value.parameters.push({
    parametername: '',
    datatype: 'string',
    required: false,
    description: ''
  })
}

// 删除Api参数

const removeParameter = (index: number) => {
  formData.value.parameters.splice(index, 1)
}

// 添加排序方法
const sortByGroup = (a: ApiInfo, b: ApiInfo) => {
  // 如果tags是数组，则比较第一个tag
  const tagA = Array.isArray(a.tags) ? a.tags[0] || '' : ''
  const tagB = Array.isArray(b.tags) ? b.tags[0] || '' : ''
  return tagA.localeCompare(tagB)
}

const submitGroupForm = async () => {
  try {
    await createApiGroup({
      apipath: groupFormData.value.apipath,
      register: {
        needauth: groupFormData.value.register.needauth,
        groupname: groupFormData.value.register.groupname,
        enable: groupFormData.value.register.enable
      }
    })
    ElMessage({
      message: '分组创建成功',
      type: 'success'
    })
    groupDrawer.value.visible = false
    await loadApiInfo()
  } catch (error: any) {
    ElMessage({
      message: error.message || '创建分组失败',
      type: 'error'
    })
  }
}

const handleSizeChange = (val: number) => {
  pageSize.value = val
  currentPage.value = 1
}

const handleCurrentChange = (val: number) => {
  currentPage.value = val
}

// 添加请求方法标签类型处理函数
const getMethodTagType = (method: string) => {
  const types: Record<string, string> = {
    get: 'success',
    post: 'primary',
    put: 'warning',
    delete: 'danger'
  }
  return types[method.toLowerCase()] || ''
}

// 添加请求方法选项数组
const methodOptions = [
  { label: 'GET-查询', value: 'GET' },
  { label: 'POST-创建', value: 'POST' },
  { label: 'PUT-更新', value: 'PUT' },
  { label: 'DELETE-删除', value: 'DELETE' }
]

// 添加删除分组的响应式数据
const deleteGroupDrawer = ref({
  visible: false,
  title: '删除API分组'
})

const deleteGroupFormData = ref({
  groupname: '',
  version: ''
})

// 添加删除分组的提交方法
const submitDeleteGroup = async () => {
  try {
    await ElMessageBox.confirm(
      `确定要删除分组 "${deleteGroupFormData.value.groupname}" 吗？此操作不可恢复！`,
      '警告',
      {
        confirmButtonText: '确定删除',
        cancelButtonText: '取消',
        type: 'warning',
      }
    )

    await deleteApiGroup(deleteGroupFormData.value)
    ElMessage({
      type: 'success',
      message: '分组删除成功',
    })
    deleteGroupDrawer.value.visible = false
    await loadApiInfo() // 重新加载数据
  } catch (error: any) {
    if (error !== 'cancel') {
      ElMessage({
        type: 'error',
        message: error.message || '删除分组失败',
      })
    }
  }
}

const testDrawer = ref({
  visible: false,
  title: 'API接口测试'
})

const testFormData = ref({
  method: '',
  path: '',
  requestBody: '',
  response: '',
  useAuth: true // 默认开启鉴权
})

const getAccessToken = () => {
  try {
    const tokenData = localStorage.getItem('vben-web-ele-5.4.7-dev-core-access')
    if (tokenData) {
      const parsed = JSON.parse(tokenData)
      return parsed.accessToken
    }
    return null
  } catch (error) {
    console.error('获取token失败:', error)
    return null
  }
}

const handleTest = (row: any) => {
  testDrawer.value.visible = true
  testFormData.value = {
    method: row.method,
    path: row.path,
    requestBody: '',
    response: '',
    useAuth: true
  }
}

const executeTest = async () => {
  try {
    const url = testFormData.value.path
    const options: any = {
      method: testFormData.value.method,
      headers: {
        'Content-Type': 'application/json'
      }
    }

    // 处理鉴权
    if (testFormData.value.useAuth) {
      const token = getAccessToken()
      if (token) {
        options.headers['Authorization'] = `Bearer ${token}`
      } else {
        ElMessage.warning('未找到访问令牌')
        return
      }
    }

    // 处理请求体
    if (testFormData.value.method !== 'GET' && testFormData.value.requestBody) {
      try {
        options.body = JSON.stringify(JSON.parse(testFormData.value.requestBody))
      } catch (e) {
        ElMessage.error('请求参数JSON格式不正确')
        return
      }
    }

    const response = await fetch(url, options)
    const responseText = await response.text() // 先获取文本响应

    try {
      // 尝试解析为JSON
      const data = JSON.parse(responseText)
      testFormData.value.response = JSON.stringify(data, null, 2)
    } catch (e) {
      // 如果不是JSON，直接显示文本
      testFormData.value.response = responseText
    }

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }
  } catch (error: any) {
    testFormData.value.response = `测试失败: ${error.message}`
    ElMessage.error('接口测试失败')
  }
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

.param-item {
  margin-bottom: 10px;
  
  .param-actions {
    display: flex;
    align-items: center;
    height: 100%;
    padding-top: 5px;
  }

  :deep(.el-select) {
    width: 100%;
  }
}

.param-item:first-child:last-child .param-actions {
  visibility: hidden;
}

.pagination-container {
  margin-top: 20px;
  display: flex;
  justify-content: flex-end;
}

.el-tag {
  text-transform: uppercase;
}

.response-box {
  padding: 10px;
  background-color: #f5f7fa;
  border-radius: 4px;
  max-height: 300px;
  overflow-y: auto;
  white-space: pre-wrap;
  word-wrap: break-word;
}
</style>
