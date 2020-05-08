import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'
import createPersistedState from 'vuex-persistedstate'

Vue.use(Vuex)

const url = 'http://localhost:9090'

export default new Vuex.Store({
  plugins: [createPersistedState({
    storage: window.sessionStorage
  })],
  state: {
    posts: []
  },
  mutations: {
    setPosts: function (state, payload) {
      if (payload.length > 1) {
        payload.sort(function (a, b) {
          return a.id > b.id ? -1 : a.id < b.id ? 1 : 0
        })
      }
      state.posts = payload
    }
  },
  actions: {
    loadPostsFromServer: function (context, kind) {
      axios.get(url + '/api/posts?kind=' + kind)
        .then((response) => {
          context.commit('setPosts', response.data)
        }).catch(err => {
          console.log(err)
        })
    }
  },
  getters: {
    getPosts: (state) => {
      return state.posts
    },
    getPostById: (state) => (id) => {
      return state.posts.find(item => item.id === id)
    }
  },
  modules: {
  }
})
