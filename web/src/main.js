import Vue from 'vue'
import './plugins/axios'
import App from './App.vue'
import router from './router'
import store from './store'
import Carousel from './components/Carousel'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import lang from 'element-ui/lib/locale/lang/ko'
import Axios from 'axios'
import MyAccount from './components/MyAccount'
import Board from './components/Board'
import General from './components/General'
import NavMenu from './components/NavMenu'

Vue.config.productionTip = false

// Axios
Vue.prototype.$http = Axios
// Vue.prototype.$api_url = 'http://localhost:9090'

// UI
Vue.use(ElementUI, { lang })

// Component
Vue.component(Carousel.name, Carousel)
Vue.component(MyAccount.name, MyAccount)
Vue.component(Board.name, Board)
Vue.component(General.name, General)
Vue.component(NavMenu.name, NavMenu)

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
