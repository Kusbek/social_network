import { createRouter, createWebHistory } from 'vue-router'
import getUser from '../composables/getUser'
import Profile from '../views/Profile.vue'
import Login from '../views/auth/Login.vue'
import Signup from '../views/auth/Signup.vue'


const requireAuth = (to, from, next) => {
  const { user } = getUser()
  if(!user.value) {
    next({name: 'Login'})
  } else {
    next()
  }
}

const routes = [
  {
    path: '/',
    name: 'Profile',
    component: Profile,
    beforeEnter: requireAuth,
  },
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/signup',
    name: 'Signup',
    component: Signup
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
