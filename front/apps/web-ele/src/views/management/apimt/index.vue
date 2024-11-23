<template>
    <div class="user-management">
      <div class="operation-bar">
        <el-button type="primary" @click="handleAdd">
          <el-icon><Plus /></el-icon>新增API
        </el-button>
      </div>

      <div class="table-container">
        <el-table :data="tableData" stripe style="width: 100%" border @selection-change="handleSelectionChange">
          <el-table-column type="selection" width="55" />
          <el-table-column label="ID" width="80">
            <template #default="scope">
              {{ scope.$index + 1 }}
            </template>
          </el-table-column>
          <el-table-column prop="path" label="API路径" width="200" />
          <el-table-column prop="method" label="请求方法" width="100" />
          <el-table-column prop="description" label="API简介" width="200" />
          <el-table-column label="分组" width="150">
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

      <el-drawer
        v-model="drawer.visible"
        :title="drawer.title"
        size="800px"
      >
        <el-form :model="formData" label-width="100px">
          <el-form-item label="API路径">
            <el-input v-model="formData.path" placeholder="请输入API路径" />
          </el-form-item>
          <el-form-item label="鉴权">
            <el-switch
              v-model="formData.needAuth"
              active-text="是"
              inactive-text="否"
            />
          </el-form-item>
          <el-form-item label="API版本">
            <el-input v-model="formData.version" placeholder="请输入API版本" />
          </el-form-item>
          <el-form-item label="请求方法">
            <el-select v-model="formData.method" placeholder="请选择请求方法">
              <el-option label="GET-查询" value="GET" />
              <el-option label="POST-创建" value="POST" />
              <el-option label="PUT-更新" value="PUT" />
              <el-option label="DELETE-删除" value="DELETE" />
            </el-select>
          </el-form-item>
          <el-form-item label="API分组">
            <el-select
              v-model="formData.tags"
              multiple
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
                    <el-input v-model="param.name" placeholder="字段名" />
                  </el-form-item>
                </el-col>
                <el-col :span="6">
                  <el-form-item :label="index === 0 ? '数据类型' : ''" label-width="80px">
                    <el-select 
                      v-model="param.type" 
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
    </div>
  </template>
  
  <script lang="ts" setup>
  import { ref, onMounted } from 'vue'
  import { getApiInfo } from '#/api/user/apiinfo'
  import type { ApiInfo } from '#/api/user/apiinfo'
  import { Plus, Delete } from '@element-plus/icons-vue'

  const tableData = ref<ApiInfo[]>([])

  const drawer = ref({
    visible: false,
    title: '新增API'
  })

  const formData = ref({
    path: '',
    method: '',
    needAuth: false,
    version: '',
    tags: [] as string[],
    description: '',
    parameters: [] as Parameter[]
  })

  const selectedRows = ref([])

  const existingTags = ref<string[]>([])

  const loadApiInfo = async () => {
    try {
      const response = await getApiInfo()
      console.log('API响应数据:', response)
      if (response.apiInfo) {
        tableData.value = response.apiInfo
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

  const handleAdd = () => {
    drawer.value.visible = true
    formData.value = {
      path: '',
      method: '',
      needAuth: false,
      tags: [],
      description: '',
      parameters: []
    }
  }

  const submitForm = () => {
    console.log('提交的表单数据:', formData.value)
    drawer.value.visible = false
  }
  
  const handleEdit = (row: any) => {
    console.log('编辑API:', row)
  }
  
  const handleDelete = (row: any) => {
    console.log('删除API:', row)
  }

  const addParameter = () => {
    formData.value.parameters.push({
      name: '',
      type: 'string',
      required: false,
      description: ''
    })
  }

  const removeParameter = (index: number) => {
    formData.value.parameters.splice(index, 1)
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
  </style>
  