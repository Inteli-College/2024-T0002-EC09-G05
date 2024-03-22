import { createRouter, createWebHistory } from 'vue-router'
import LoginView from '../views/LoginView.vue'
import PathNotFound from '../views/PathNotFound.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/', 
      name: 'home',
      component: () => import('../views/DefaultView.vue')

    },
    {
      path: '/auth', 
      name: 'auth',
      component: LoginView
    },
    { path: '/:pathMatch(.*)*', component: PathNotFound },
    
  ],strict: true
})

export default router
