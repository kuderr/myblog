import { Module } from "vuex";
import { PostsState } from "./types";
import { RootState } from "../types";
import { getters } from "./getters";
import { actions } from "./actions";
import { mutations } from "./mutations";

export const state: PostsState = {
  posts: [],
  error: false,
};

const namespaced: boolean = true;

export const posts: Module<PostsState, RootState> = {
  namespaced,
  state,
  getters,
  actions,
  mutations,
};
