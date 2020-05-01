import Vue from 'vue'
import VueRouter from 'vue-router'
import Main from '../views/Main.vue'
import General from '../views/General'
import Study from '../views/Study'
import Qna from '../views/Qna'
import User from '../views/User'
import GeneralDetail from '../views/GeneralDetail'

Vue.use(VueRouter)

const routes = [
  {
    path: '/main',
    name: 'Main',
    component: Main
  },
  {
    path: '/main/general',
    name: 'general',
    component: General
  },
  {
    path: '/main/general/:id',
    name: 'generalDetail',
    component: GeneralDetail
  },
  {
    path: '/main/study',
    name: 'study',
    component: Study
  },
  {
    path: '/main/qna',
    name: 'qna',
    component: Qna
  },
  {
    path: '/main/user',
    name: 'user',
    component: User
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
