import { getUserPosts, addPost, deletePost, authenticate } from "@/api";
import { User } from "./models";
import { Post } from "..posts/models";

export default {
  state: {
    userPosts: Array<Post>(),
    user: new User(),
    jwt: { token: localStorage.token || "" },
  },

  mutations: {
    setJwtToken(state, payload) {
      localStorage.token = payload.jwt.token;
      state.jwt = payload.jwt;
    },
    clearUserData(state) {
      state.jwt = "";
      localStorage.token = "";
    },

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
    async login({ state, commit }, userData) {
      try {
        commit("switchLoading");
        let res = await authenticate(userData);
        commit("setJwtToken", { jwt: res.data });
        const tokenParts = res.data.token.split(".");
        const body = JSON.parse(atob(tokenParts[1]));
        state.user.username = body.sub.username;
        return body;
      } catch (error) {
        commit("setError", error);
        return null;
      } finally {
        commit("switchLoading");
      }
    },
    logout({ commit }) {
      commit("clearUserData");
    },
    async setDataFromToken({ commit, state }) {
      try {
        commit("switchLoading");
        const tokenParts = localStorage.token.split(".");
        const body = JSON.parse(atob(tokenParts[1]));
        state.user.username = body.sub.username;
        return body;
      } catch (error) {
        commit("setError", error);
        return null;
      } finally {
        commit("switchLoading");
      }
    },

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
