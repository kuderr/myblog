import Vue from 'vue'
import Vuex from 'vuex'
import shared from './shared'
import posts from './posts'
import user from './user'

Vue.use(Vuex)

const store = () =>
  new Vuex.Store({
    modules: {
      shared,
      posts,
      user,
    },
  })

export default store
