import { requestClient } from '#/api/request';

export interface ApiInfo {
    path: string;
    method: string;
    description: string;
    tags: string[];
}

interface ApiResponse {
    apiInfo: ApiInfo[];
}

export async function getApiInfo() {
    return requestClient.get<ApiResponse>('/user/apiinfo');
}

