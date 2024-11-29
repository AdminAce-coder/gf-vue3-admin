import type { RouteRecordRaw } from 'vue-router';

import { BasicLayout } from '#/layouts';

const routes: RouteRecordRaw[] = [
  {
    component: BasicLayout,
    meta: {
        badgeType: 'dot',
        order: 5,
        title: "实例管理",
        icon: 'lucide:layout-dashboard',
      },
    name: 'Instance',
    path: '/instance',
    redirect: '/instance/list',
    children: [
        {
            name: '实例列表',
            path: '/instance/list',
            meta: {
              title: '实例列表',
            },
            component: () => import('#/views/instance//index.vue'),
        },
    ],
  },
];

export default routes;
