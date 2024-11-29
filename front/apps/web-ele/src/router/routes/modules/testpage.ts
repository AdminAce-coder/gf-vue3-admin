import type { RouteRecordRaw } from 'vue-router';

import { BasicLayout } from '#/layouts';

const routes: RouteRecordRaw[] = [
  {
    component: BasicLayout,
    meta: {
        badgeType: 'dot',
        order: 6,
        title: "测试页面",
        icon: 'lucide:layout-dashboard',
      },
    name: 'TestPage',
    path: '/testpage',
    redirect: '/testpage/index',
    children: [
        {
            name: '测试页面',
            path: '/testpage/index',
              component: () => import('#/views/testviews/index.vue'),
        },
    ],
  },
];

export default routes;
