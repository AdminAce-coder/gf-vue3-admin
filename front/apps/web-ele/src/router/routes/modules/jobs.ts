import type { RouteRecordRaw } from 'vue-router';

import { BasicLayout } from '#/layouts';

const routes: RouteRecordRaw[] = [
  {
    component: BasicLayout,
    meta: {
        badgeType: 'dot',
        order: 6,
        title: "任务管理",
        icon: 'lucide:layout-dashboard',
      },
    name: 'Jobs',
    path: '/jobs',
    redirect: '/jobs/list',
    children: [
        {
            name: '任务列表',
            path: '/jobs/list',
            meta: {
              title: '任务列表',
            },
            component: () => import('#/views/jobs/index.vue'),
        },
    ],
  },
];

export default routes;
