import Vue from "vue";
import App from "@/App.vue";
import vuetify from "@/plugins/vuetify";
import { TiptapVuetifyPlugin } from "tiptap-vuetify";
import "tiptap-vuetify/dist/main.css";
import router from "@/router";
import store from "@/store";
import { isValidToken } from "@/utils";

Vue.config.productionTip = false;
Vue.use(TiptapVuetifyPlugin, {
  vuetify,
});

if (isValidToken(store.state.user.jwt.token)) {
  store.dispatch("setDataFromToken").then((res) => {
    store.dispatch("fetchUserPosts", 1);
  });
}

new Vue({
  vuetify,
  router,
  store,
  render: (h) => h(App),
}).$mount("#app");
