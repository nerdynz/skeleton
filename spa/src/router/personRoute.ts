const personRoute =  {
  name: 'PersonList',
  path: '/users',
  meta: {
    title: 'Users',
    icon: 'free-person',
    inSidebar: true
  },
  component: () => import('../views/people/PersonList.vue'),
  children: [
    {
      name: 'PersonCreate',
      path: '/users/create',
      meta: {
        title: 'Create User',
      },
      component: () => import('../views/people/PersonEdit.vue'),
    },
    {
      name: 'PersonEdit',
      path: '/users/edit/:ulid',
      meta: {
        title: 'Edit User',
      },
      component: () => import('../views/people/PersonEdit.vue'),
    },
  ],
}

export default personRoute