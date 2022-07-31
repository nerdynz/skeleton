const personRoute =  {
  name: 'PersonList',
  path: '/people',
  meta: {
    title: 'Person',
    icon: 'free-person',
    inSidebar: true
  },
  component: () => import('../views/people/PersonList.vue'),
  children: [
    {
      name: 'PersonCreate',
      path: '/people/create',
      meta: {
        title: 'Person Create',
      },
      component: () => import('../views/people/PersonEdit.vue'),
    },
    {
      name: 'PersonEdit',
      path: '/people/edit/:ulid',
      meta: {
        title: 'Person Edit',
      },
      component: () => import('../views/people/PersonEdit.vue'),
    },
  ],
}

export default personRoute