import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import pageRoutes from './pageRoutes'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
      meta: {
        sidebar: true,
        icon: 'fad-home-alt',
      },
    },
    ...pageRoutes,
  ],
})

export default router
