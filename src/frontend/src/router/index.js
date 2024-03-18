import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import PathNotFound from '../views/PathNotFound.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/', 
      name: 'home',
      component: () => import('../views/ComumView.vue')

    },
    {
      path: '/auth', 
      name: 'auth',
      component: LoginView
    },
    {
      path: '/default', 
      name: 'default',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/DefaultView.vue')
    },
    { path: '/:pathMatch(.*)*', component: PathNotFound },
    
  ],strict: true
})

export default router
