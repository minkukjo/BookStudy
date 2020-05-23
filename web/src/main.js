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
import Detail from './components/Detail'
import Title from './components/Title'
import Axios from 'axios'

Vue.config.productionTip = false

// Axios
Vue.prototype.$http = Axios
// Vue.prototype.$http_url = 'http://localhost:9090'
Vue.prototype.$http_url = 'http://34.67.130.46:9090'

// UI
Vue.use(ElementUI, { lang })

// Component
Vue.component(Carousel.name, Carousel)
Vue.component(UserInform.name, UserInform)
Vue.component(Board.name, Board)
Vue.component(Detail.name, Detail)
Vue.component(Title.name, Title)

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
