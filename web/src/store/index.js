import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'

Vue.use(Vuex)

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
      axios.get('http://localhost:9090/api/posts?kind=' + kind)
        .then((response) => {
          console.log(response.data)
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
      console.log(id)
      return state.posts.find(item => item.id === id)
    }
  },
  modules: {
  }
})
