import { getPosts, getPost, addView } from '@/api'
import { Post } from './models'

export default {
  state: {
    posts: Array<Post>(),
    currentPost: new Post(),
  },

  mutations: {
    postsLoaded(state, payload: Post[]) {
      state.posts = payload
    },
    postLoaded(state, payload: Post) {
      state.currentPost = payload
    },
    postAdded(state, post: Post) {
      state.posts.unshift(post)
    },
    postDeleted(state, postId: number) {
      let i: number
      state.posts.forEach((post: Post, index: number) => {
        if (post.id === postId) {
          i = index
          return
        }
      })

      state.posts.splice(i, 1)
    },
    postUpdated(state, payload: Post) {
      state.posts.forEach((post: Post, index: number) => {
        if (post.id === payload.id) {
          state.posts[index] = payload
          return
        }
      })
    },
  },

  actions: {
    async fetchPosts({ commit }) {
      let res = await getPosts()
      commit('postsLoaded', res.data)
    },
    async fetchPost({ state, commit }, postId) {
      let res = await getPost(postId)
      commit('postLoaded', res.data)

      let postsVisited = sessionStorage.getItem('postsVisited')
      if (postsVisited){
        var posts = postsVisited.split(';')
      }
      else{
        var posts = Array<string>()
      }
      
      if(posts.indexOf(postId.toString()) === -1){
        await addView(postId)
        posts.push(postId)
        sessionStorage.setItem('postsVisited', posts.join(';'))
      }

    },
  },
}
