import {
  getUserPosts,
  addPost,
  deletePost,
  authenticate,
  updatePost,
  updatePublishedStatus,
} from '@/api'
import { User } from './models'
import { Post } from '..posts/models'

export default {
  state: {
    userPosts: Array<Post>(),
  },

  mutations: {
    userPostsLoaded(state, payload: Post[]) {
      state.userPosts = payload
    },
    userPostAdded(state, payload: Post) {
      state.userPosts.unshift(payload)
    },
    userPostDeleted(state, postId: number) {
      let i: number
      state.userPosts.forEach((post: Post, index: number) => {
        if (post.id === postId) {
          i = index
          return
        }
      })
      state.userPosts.splice(i, 1)
    },
    userPostUpdated(state, payload: Post) {
      state.userPosts.forEach((post: Post, index: number) => {
        if (post.id === payload.id) {
          state.userPosts[index] = payload
          return
        }
      })
    },
  },

  actions: {
    async fetchUserPosts({ state, commit }, userId: number) {
      try {
        let res = await getUserPosts(userId, state.token)
        commit('userPostsLoaded', res.data)
      } catch (error) {
        commit('setError', error)
      }
    },
    async addPost({ state, commit }, postData: Post) {
      let res = await addPost(postData, state.token)
      postData.id = res.data.postId
      commit('userPostAdded', postData)
      return res
    },
    updatePublishedStatus({ state, commit }, post: Post) {
      updatePublishedStatus(post.id, post.published, state.token)
      if (post.published === false) commit('postDeleted', post.id)
      else commit('postAdded', post)
    },
    deletePost({ state, commit }, post: Post) {
      deletePost(post.id, state.token)
      commit('userPostDeleted', post.id)
      if (post.published === true) commit('postDeleted', post.id)
    },
    updatePost({ state, commit }, post: Post) {
      updatePost(post, state.token)
      commit('userPostUpdated', post)
      if (post.published === true) commit('postUpdated', post)
    },
  },
}
