export default [
  {
    name: 'PageList',
    path: '/pages',
    meta: {
      title: 'Pages',
      icon: 'fad-file-alt',
      sidebar: true,
    },
    component: () => import('../views/pages/PageList.vue'),
    children: [],
  },
  {
    name: 'PageCreate',
    path: '/pages/create',
    meta: {
      title: 'Page Create',
    },
    component: () => import('../views/pages/PageEdit.vue'),
  },
  {
    name: 'PageEdit',
    path: '/pages/edit/:ulid',
    meta: {
      title: 'Page Edit',
    },
    component: () => import('../views/pages/PageEdit.vue'),
  },
]
