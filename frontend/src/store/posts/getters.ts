import { GetterTree } from "vuex";
import { PostsState, Post } from "./types";
import { RootState } from "../types";

export const getters: GetterTree<PostsState, RootState> = {
  posts(state): Post[] {
    const { posts } = state;
    return posts;
  },
};
