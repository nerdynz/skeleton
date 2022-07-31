import { useUserAccessStore } from '@/store/userAccess'
import Login from '@/views/Login.vue'
import { createRouter, createWebHistory } from 'vue-router'
import personRoute from './personRoute'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'home',
      meta: {
        title: 'Home',
        icon: 'free-house',
        inSidebar: true,
      },
      component: () => import('../views/Home.vue'),
    },
    personRoute,
    {
      name: 'login',
      path: '/login',
      component: Login,
      meta: {
        title: 'Login',
        isHidden: true,
        requiresAuth: false,
      },
    },
    {
      name: 'logout',
      path: '/logout',
      meta: {
        title: 'Logout',
        icon: 'free-logout',
        inSidebar: true,
        cssClass: 'is-bottom',
      },
      component: Login,
    },
  ],
})

router.beforeEach((to, from, next) => {
  const userAccess = useUserAccessStore()
  if (to.name === 'logout') {
    console.log('logging out')
    userAccess.logout()
    next({ name: 'login' })
    return
  }
  if (to.meta.requiresAuth === false) {
    // explictly check the requiresAuth flag is false. assume undefined means auth is required
    next()
    return
  }
  if (!userAccess.isLoggedIn) {
    next({ name: 'login' })
    return
  }
  next()
  return
})

export default router
