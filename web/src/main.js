import Vue from 'vue'
import './plugins/axios'
import App from './App.vue'
import router from './router'
import store from './store'
import Carousel from './components/Carousel'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import lang from 'element-ui/lib/locale/lang/ko'
import UserInform from './components/UserInform'
import Board from './components/Board'
import NavMenu from './components/NavMenu'
import Detail from './components/Detail'
import Title from './components/Title'

Vue.config.productionTip = false

// UI
Vue.use(ElementUI, { lang })

// Component
Vue.component(Carousel.name, Carousel)
Vue.component(UserInform.name, UserInform)
Vue.component(Board.name, Board)
Vue.component(NavMenu.name, NavMenu)
Vue.component(Detail.name, Detail)
Vue.component(Title.name, Title)

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
