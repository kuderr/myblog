import {
  getUserPosts,
  addPost,
  deletePost,
  authenticate,
  updatePost,
  updatePublishedStatus,
} from '@/api'
import { User } from './models'

export default {
  state: {
    userPosts: Array<any>(),
    user: new User(),
    token: '',
  },

  mutations: {
    setUser(state, payload) {
      state.user = payload
    },
    setToken(state, payload) {
      state.token = payload
      localStorage.token = payload
    },
    clearUserData(state, payload) {
      state.token = null
      localStorage.removeItem('token')
    },
    userPostsLoaded(state, payload) {
      state.userPosts = payload
    },
    userPostAdded(state, payload) {
      state.userPosts.unshift(payload)
    },
    userPostDeleted(state, postId: number) {
      let i: number
      state.userPosts.forEach((post, index: number) => {
        if (post.id === postId) {
          i = index
          return
        }
      })
      state.userPosts.splice(i, 1)
    },
    userPostUpdated(state, payload) {
      state.userPosts.forEach((post, index: number) => {
        if (post.id === payload.id) {
          state.userPosts[index] = payload
          return
        }
      })
    },
  },

  actions: {
    async login({ commit }, userData) {
      try {
        commit('switchLoading')
        let res = await authenticate(userData)
        const token = res.data.token
        const tokenParts = token.split('.')
        const body = JSON.parse(atob(tokenParts[1]))
        await commit('setUser', body)
        await commit('setToken', token)
        return body
      } catch (error) {
        console.log(error)
        commit('setError', error)
        return null
      } finally {
        commit('switchLoading')
      }
    },
    logout({ commit }) {
      commit('clearUserData')
    },
    async setDataFromToken({ commit }, token) {
      try {
        commit('switchLoading')
        const tokenParts = token.split('.')
        const body = JSON.parse(atob(tokenParts[1]))
        await commit('setUser', body)
        await commit('setToken', token)
        return body
      } catch (error) {
        commit('setError', error)
        return null
      } finally {
        commit('switchLoading')
      }
    },
    async fetchUserPosts({ state, commit }, userId: number) {
      try {
        let res = await getUserPosts(userId, state.token)
        commit('userPostsLoaded', res.data)
      } catch (error) {
        commit('setError', error)
      }
    },
    async addPost({ state, commit }, postData) {
      let res = await addPost(postData, state.token)
      postData.id = res.data.postId
      commit('userPostAdded', postData)
      return res
    },
    updatePublishedStatus({ state, commit }, post) {
      updatePublishedStatus(post.id, post.published, state.token)
      if (post.published === false) commit('postDeleted', post.id)
      else commit('postAdded', post)
    },
    deletePost({ state, commit }, post) {
      deletePost(post.id, state.token)
      commit('userPostDeleted', post.id)
      if (post.published === true) commit('postDeleted', post.id)
    },
    updatePost({ state, commit }, post) {
      updatePost(post, state.token)
      commit('userPostUpdated', post)
      if (post.published === true) commit('postUpdated', post)
    },
  },
}
