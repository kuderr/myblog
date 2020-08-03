import { MutationTree } from "vuex";
import { PostsState, Post } from "./types";

export const mutations: MutationTree<PostsState> = {
  postsLoaded(state, payload: Post[]) {
    state.error = false;
    state.posts = payload;
  },
  postsError(state) {
    state.error = true;
    state.posts = undefined;
  },
};
