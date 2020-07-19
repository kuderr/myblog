import Vue from "vue";
import App from "@/App.vue";
import vuetify from "@/plugins/vuetify";
import { TiptapVuetifyPlugin } from "tiptap-vuetify";
import "tiptap-vuetify/dist/main.css";
import router from "@/router";
import store from "@/store";

Vue.config.productionTip = false;
Vue.use(TiptapVuetifyPlugin, {
  vuetify,
});

new Vue({
  vuetify,
  router,
  store,
  render: (h) => h(App),
}).$mount("#app");
