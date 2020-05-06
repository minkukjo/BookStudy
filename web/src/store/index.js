import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'

Vue.use(Vuex)

const url = 'http://localhost:9090'

export default new Vuex.Store({
  state: {
    posts: []
  },
  mutations: {
    setPosts: function (state, payload) {
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
