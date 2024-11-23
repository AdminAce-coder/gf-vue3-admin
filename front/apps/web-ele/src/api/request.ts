/**
 * 该文件可自行根据业务逻辑进行调整
 */
import type { HttpResponse } from '@vben/request';

import { useAppConfig } from '@vben/hooks';
import { preferences } from '@vben/preferences';
import {
  authenticateResponseInterceptor,
  errorMessageResponseInterceptor,
  RequestClient,
} from '@vben/request';
import { useAccessStore } from '@vben/stores';

import { ElMessage } from 'element-plus';

import { useAuthStore } from '#/store';

import { refreshTokenApi } from './core';

const { apiURL } = useAppConfig(import.meta.env, import.meta.env.PROD);

function createRequestClient(baseURL: string) {
  const client = new RequestClient({
    baseURL,
  });

  /**
   * 重新认证逻辑
   */
  async function doReAuthenticate() {
    console.log('[Auth] Starting re-authentication process');
    console.warn('Access token or refresh token is invalid or expired.');
    const accessStore = useAccessStore();
    const authStore = useAuthStore();
    accessStore.setAccessToken(null);
    if (
      preferences.app.loginExpiredMode === 'modal' &&
      accessStore.isAccessChecked
    ) {
      console.log('[Auth] Showing login expired modal');
      accessStore.setLoginExpired(true);
    } else {
      console.log('[Auth] Performing logout');
      await authStore.logout();
    }
  }

  /**
   * 刷新token逻辑
   */
  async function doRefreshToken() {
    console.log('[Auth] Attempting to refresh token');
    const accessStore = useAccessStore();
    try {
      const resp = await refreshTokenApi();
      const newToken = resp.data;
      console.log('[Auth] Token refresh successful');
      accessStore.setAccessToken(newToken);
      return newToken;
    } catch (error) {
      console.error('[Auth] Token refresh failed:', error);
      throw error;
    }
  }

  function formatToken(token: null | string) {
    return token ? `Bearer ${token}` : null;
  }

  // 请求头处理
  client.addRequestInterceptor({
    fulfilled: async (config) => {
      const accessStore = useAccessStore();

      config.headers.Authorization = formatToken(accessStore.accessToken);
      config.headers['Accept-Language'] = preferences.app.locale;
      return config;
    },
  });

  // response数据解构
  client.addResponseInterceptor<HttpResponse>({
    fulfilled: (response) => {
      const { data: responseData, status } = response;

      const { code, data } = responseData;
      if (status >= 200 && status < 400 && code === 0) {
        return data;
      }
      throw Object.assign({}, response, { response });
    },
  });

  // token过期的处理
  client.addResponseInterceptor(
    authenticateResponseInterceptor({
      client,
      doReAuthenticate,
      doRefreshToken,
      enableRefreshToken: preferences.app.enableRefreshToken,
      formatToken,
    }),
  );

  // 通用的错误处理,如果没有进入上面的错误处理逻辑，就会进入这里
  client.addResponseInterceptor(
    errorMessageResponseInterceptor((msg: string, error) => {
      // 这里可以根据业务进行定制,你可以拿到 error 内的信息进行定制化处理，根据不同的 code 做不同的提示，而不是直接使用 message.error 提示 msg
      // 当前mock接口返回的错误字段是 error 或者 message
      const responseData = error?.response?.data ?? {};
      const errorMessage = responseData?.error ?? responseData?.message ?? '';
      // 如果没有错误信息，则会根据状态码进行提示
      ElMessage.error(errorMessage || msg);
    }),
  );

  client.addResponseInterceptor(
    (response) => {
      // 检查响应体中的 code
      const res = response.data;
      console.log('[Response Interceptor]', res);
      
      if (res && res.code === 401) {
        // 将业务层面的401转换为真实的HTTP 401错误
        return Promise.reject({
          response: {
            status: 401,
            data: res
          }
        });
      }
      
      // 其他业务错误码处理...
      if (res.code && res.code !== 200) {
        return Promise.reject({
          response: {
            status: res.code,
            data: res
          }
        });
      }

      return response;
    },
    (error) => {
      console.error('[Response Error Interceptor]', error);
      return Promise.reject(error);
    }
  );

  return client;
}

export const requestClient = createRequestClient(apiURL);

export const baseRequestClient = new RequestClient({ baseURL: apiURL });
