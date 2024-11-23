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
    apiPath: string;           // API路径
    isauthentication: boolean; // 是否鉴权
    apiVersion: string;        // API版本
    method: string;            // 请求方法
    apiGroup: string;          // API分组
    description: string;       // 描述
    parameters: Parameter[];   // 参数列表
}

export async function getApiInfo() {
    return requestClient.get<ApiResponse>('/v1/user/apiinfo');
}

export async function createApiInfo(data: CreateApiRequest) {
    return requestClient.post('/v1/menu/createapi', data);
}
