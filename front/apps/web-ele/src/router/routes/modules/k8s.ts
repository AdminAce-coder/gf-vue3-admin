import type { RouteRecordRaw } from 'vue-router';

import { BasicLayout } from '#/layouts';

const routes: RouteRecordRaw[] = [
  {
    component: BasicLayout,
    meta: {
        badgeType: 'dot',
        order: 4,
        title: "K8s",
        icon: 'lucide:layout-dashboard',
      },
    name: 'K8s',
    path: '/k8s',
    redirect: '/k8s/dashboard',
    children: [
        {
            name: 'Dashboard',
            path: '/dashboard',
            component: () => import('#/views/k8s/index.vue'),
        },
    ],
  },
];

export default routes;
