import { ActionTree } from "vuex";
import { PostsState, Post } from "./types";
import { RootState } from "../types";
import { getPosts, getPost } from "@/api";

export const actions: ActionTree<PostsState, RootState> = {
  async fetchPosts({ commit }) {
    let res = await getPosts();
    commit("postsLoaded", res.data);
  },
};
