import { requestClient } from '#/api/request';

// 获取接口信息的请求体
export interface ApiInfo {
    path: string;
    method: string;
    description: string;
    tags: string[];
}

interface ApiResponse {
    apiInfo: ApiInfo[];
}

// API参数接口
export interface Parameter {
    parametername: string;  // 参数名
    datatype: string;      // 数据类型
    required: boolean;     // 是否必填
    description: string;   // 描述
}

// 创建接口的请求体
export interface CreateApiRequest {
    apiname: string;           // API名称
    method: string;            // 请求方法
    apiGroup: string;          // API分组
    description: string;       // 描述
    apiversion: string;        // API版本，注意这里改为小写的 version
    parameters: Parameter[];   // 参数列表
}

// 创建API分组请求体
export interface CreateApiGroupRequest {
    apipath: string;           // API路径
    register: {
        needauth: boolean;     // 是否需要认证
        groupname: string;     // 分组名称
        enable: boolean;       // 是否启用
    }
}
// 删除API分组请求体
export interface DeleteApiGroupRequest {
    groupname: string;           // 分组名称
    version: string;            // 版本
}
// 删除API请求体
export interface DeleteApiRequest {
    apipath: string;           // API路径
    apigroup: string;            // Api分组
}


export async function getApiInfo() {
    return requestClient.get<ApiResponse>('/v1/user/apiinfo');
}

export async function createApiInfo(data: CreateApiRequest) {
    return requestClient.post('/v1/menu/createapi', data);
}



// 创建API分组
export async function createApiGroup(data: CreateApiGroupRequest) {
    return requestClient.post('/v1/menu/capigp', data);
}

// 删除API分组
export async function deleteApiGroup(data: DeleteApiGroupRequest) {
    return requestClient.delete('/v1/menu/delapigp', { data });
}

// 删除API
export async function deleteApi(data: DeleteApiRequest) {
    return requestClient.delete('/v1/menu/delapi', { data });
}

