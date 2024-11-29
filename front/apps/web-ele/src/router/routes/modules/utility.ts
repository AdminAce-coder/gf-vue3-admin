import type { RouteRecordRaw } from 'vue-router';

import { BasicLayout } from '#/layouts';

const routes: RouteRecordRaw[] = [
  {
    component: BasicLayout,
    meta: {
        badgeType: 'dot',
        order: 3,
        title: "工具",
        icon: 'lucide:layout-dashboard',
      },
    name: 'Utility',
    path: '/utility',
    redirect: '/utility/oss',
    children: [
        {
            name: 'Oss',
            path: '/utility/oss',
            component: () => import('#/views/utility/oss.vue'),
        },
    ],
  },
];

export default routes;
