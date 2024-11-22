import type { RouteRecordRaw } from 'vue-router';

import { BasicLayout } from '#/layouts';
import { $t } from '#/locales';

const routes: RouteRecordRaw[] = [
  {
    component: BasicLayout,
    meta: {
      icon: 'lucide:layout-dashboard',
      order: -1,
      title: "系统管理",
    },
    name: 'Management',
    path: '/',
    children: [
      {
        name: 'Rolemt',
        path: '/rolemt',
        component: () => import('#/views/management/rolemt/index.vue'),
        meta: {
          affixTab: true,
          icon: 'lucide:area-chart',
          title: "角色管理",
        },
      },
      {
        name: 'Usermt',
        path: '/usermt',
        component: () => import('#/views/management/usermt/index.vue'),
        meta: {
          icon: 'carbon:workspace',
          title: "用户管理",
        },
      },
      {
        name: 'Apimt',
        path: '/apimt',
        component: () => import('#/views/management/apimt/index.vue'),
        meta: {
          icon: 'carbon:workspace',
          title: "API管理",
        },
      },
    ],
  },
];

export default routes;
