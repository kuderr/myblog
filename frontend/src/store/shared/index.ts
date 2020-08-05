export default {
  state: {
    loading: false,
    error: "",
    message: "",
  },
  mutations: {
    switchLoading(state) {
      state.loading = !state.loading;
    },
    setError(state, payload) {
      state.error = payload["response"].data.message;
    },
    clearError(state) {
      state.error = null;
    },
    setMessage(state, payload) {
      state.message = payload.data.message;
    },
    setMessageFromText(state, payload) {
      state.message = payload;
    },
    clearMessage(state) {
      state.message = null;
    },
  },
  actions: {
    switchLoading({ commit }) {
      commit("switchLoading");
    },
    setError({ commit }, payload) {
      commit("setError", payload);
    },
    clearError({ commit }) {
      commit("clearError");
    },
    setMessage({ commit }, payload) {
      commit("setMessage", payload);
    },
    clearMessage({ commit }) {
      commit("clearMessage");
    },
  },
};
