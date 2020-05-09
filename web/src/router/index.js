import Vue from 'vue'
import VueRouter from 'vue-router'
import Main from '../views/Main.vue'
import General from '../views/General'
import Study from '../views/Study'
import Qna from '../views/Qna'
import User from '../views/User'
import GeneralDetail from '../views/GeneralDetail'
import Write from '../views/Write'
import QnaDetail from '../views/QnaDetail'
import StudyDetail from '../views/StudyDetail'
import Edit from '../views/Edit'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Main',
    component: Main
  },
  {
    path: '/general',
    name: 'general',
    component: General
  },
  {
    path: '/general/:id',
    name: 'generalDetail',
    component: GeneralDetail
  },
  {
    path: '/study',
    name: 'study',
    component: Study
  },
  {
    path: '/study/:id',
    name: 'studyDetail',
    component: StudyDetail
  },
  {
    path: '/qna',
    name: 'qna',
    component: Qna
  },
  {
    path: '/qna/:id',
    name: 'qnaDetail',
    component: QnaDetail
  },
  {
    path: '/user',
    name: 'user',
    component: User
  },
  {
    path: '/write',
    name: 'write',
    component: Write
  },
  {
    path: '/edit/:id',
    name: 'edit',
    component: Edit
  }
]

const router = new VueRouter({
  base: process.env.BASE_URL,
  routes
})

export default router
