import { getPosts, getPost, updatePublishedStatus } from "@/api";
import { Post } from "./models";

export default {
  state: {
    posts: Array<Post>(),
    currentPost: new Post(),
  },

  mutations: {
    postsLoaded(state, payload: Post[]) {
      state.posts = payload;
    },
    postLoaded(state, payload: Post) {
      state.currentPost = payload;
    },
    deletePost(state, postId: number) {
      let i: number;
      state.posts.forEach((post: Post, index: number) => {
        if (post.id === postId) {
          i = index;
          return;
        }
      });

      state.posts.splice(i, 1);
    },
    addPost(state, post: Post) {
      state.posts.unshift(post);
    },
  },

  actions: {
    async fetchPosts({ commit }) {
      let res = await getPosts();
      commit("postsLoaded", res.data);
    },
    async fetchPost({ commit }, postId) {
      let res = await getPost(postId);
      commit("postLoaded", res.data);
    },
    updatePublishedStatus({ commit }, post: Post) {
      updatePublishedStatus(post.id, post.published);
      if (post.published === false) commit("deletePost", post.id);
      else commit("addPost", post);
    },
  },
};
