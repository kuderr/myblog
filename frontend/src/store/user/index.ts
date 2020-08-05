import { getUserPosts, addPost, deletePost } from "@/api";
import { User } from "./models";
import { Post } from "..posts/models";

export default {
  state: {
    userPosts: Array<Post>(),
  },

  mutations: {
    userPostsLoaded(state, payload: Post[]) {
      state.userPosts = payload;
    },
    postAdded(state, payload: Post) {
      state.userPosts.unshift(payload);
    },
    postDeleted(state, postId: number) {
      let i: number;
      state.userPosts.forEach((post: Post, index: number) => {
        if (post.id === postId) {
          i = index;
          return;
        }
      });

      state.userPosts.splice(i, 1);
    },
  },

  actions: {
    async fetchUserPosts({ commit }, userId: number) {
      let res = await getUserPosts(userId);
      commit("userPostsLoaded", res.data);
    },
    async addPost({ commit }, postData: Post) {
      let res = await addPost(postData);
      postData.id = res.data.postId;
      commit("postAdded", postData);
      return res;
    },
    deletePost({ commit }, postId: number) {
      deletePost(postId);
      commit("postDeleted", postId);
    },
  },
};
