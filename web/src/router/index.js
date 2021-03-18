import { createRouter, createWebHistory } from 'vue-router'
import User from '../composables/user'

import Login from '../views/auth/Login.vue'
import Signup from '../views/auth/Signup.vue'
import MainPage from '../views/MainPage.vue'
import Profile from '../views/Profile.vue'

const requireNotAuth = async (to, from, next) => {
  const { user, getUser } = User()
  await getUser()
  if(!user.value) {
    next()
  } else {
    next({name: 'MainPage'})
  }
}

const requireAuth = async (to, from, next) => {
  const { user, getUser } = User()
  await getUser()
  if(!user.value) {
    next({name: 'Login'})
  } else {
    next()
  }
}

const routes = [
  {
    path: '/user/:id',
    name: 'Profile',
    component: Profile,
    props: true,
  },
  {
    path: '/',
    name: 'MainPage',
    component: MainPage,
    beforeEnter: requireAuth,
  },
  {
    path: '/login',
    name: 'Login',
    component: Login,
    beforeEnter:requireNotAuth,
  },
  {
    path: '/signup',
    name: 'Signup',
    component: Signup,
    beforeEnter:requireNotAuth,
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
