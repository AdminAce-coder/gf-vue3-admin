import type { RouteRecordRaw } from 'vue-router';

import { BasicLayout } from '#/layouts';

const routes: RouteRecordRaw[] = [
  {
    component: BasicLayout,
    meta: {
        badgeType: 'dot',
        order: -1,
        title: "首页",
        icon: 'lucide:layout-dashboard',
      },
    name: 'Homepage',
    path: '/',
    redirect: '/dashboard',
    children: [
        {
            name: 'Dashboard',
            path: '/dashboard',
            component: () => import('#/views/homepage/index.vue'),
        },
    ],
  },
];

export default routes;
