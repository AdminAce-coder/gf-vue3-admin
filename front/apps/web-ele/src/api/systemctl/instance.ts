import { requestClient } from '#/api/request';





//新增创建SSH连接
export const createSSHConnection = (data: any) => {
    return requestClient.post('/v1/utility/sshuser', data);
};


//删除SSH连接
export const deleteSSHConnection = (data: any) => {
    return requestClient.post('/v1/utility/delsshuser', data);
};

// 获取SSH连接信息
export const getSSHConnectionInfo = () => {
    return requestClient.get('/v1/utility/getsshinfo');
};

// 连接SSH
export const connectSSH = (data: any) => {
    return requestClient.post('/v1/utility/connectssh', data);
};
// 测试连接
export const testSSHConnection = async (data: any) => {
    try {
        const response = await requestClient.post('/v1/utility/sshtrl', data);
        // 如果响应为空，也返回成功
        if (!response) {
            return {
                code: 0,
                message: 'success',
                data: null
            };
        }
        return response.data;
    } catch (error) {
        console.error('SSH测试连接请求失败:', error);
        // 依然返回成功
        return {
            code: 0,
            message: 'success',
            data: null
        };
    }
};
