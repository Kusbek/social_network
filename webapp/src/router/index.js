import { createRouter, createWebHistory } from 'vue-router'
import User from '../composables/user'

import Login from '../views/auth/Login.vue'
import Signup from '../views/auth/Signup.vue'
import MainPage from '../views/MainPage.vue'
import Profile from '../views/Profile.vue'
import Group from '../views/group/Group.vue'
import Groups from '../views/group/Groups.vue'
import InviteForm from '../views/group/InviteForm.vue'
import CreateGroupForm from '../views/group/CreateGroupForm.vue'
import Test from '../views/Test.vue'

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
    path: '/:id',
    name: 'Profile',
    component: Profile,
    props: true,
    beforeEnter:requireAuth,
  },
  {
    path: '/group/:id',
    name: 'Group',
    component: Group,
    props: true,
    beforeEnter:requireAuth,
  },
  {
    path: '/group/:id/invite',
    name: 'InviteForm',
    component: InviteForm,
    props: true,
    beforeEnter:requireAuth,
  },
  {
    path: '/groups',
    name: 'Groups',
    component: Groups,
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
  },
  {
    path: '/creategroup',
    name: 'CreateGroupForm',
    component: CreateGroupForm,
    beforeEnter:requireAuth,
  },
  {
    path: '/test',
    name: 'Test',
    component: Test,
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
