import Vue from "vue";
import VueMeta from "vue-meta";
import VueRouter, { RouteConfig } from "vue-router";
import { authRequired } from "./authGuard";

Vue.use(VueRouter);
Vue.use(VueMeta, {
  refreshOnceOnNavigation: true,
});

const routes: RouteConfig[] = [
  {
    path: "/",
    name: "Home",
    meta: { title: "Главная" },
    component: () => import("@/views/Home.vue"),
  },
  {
    path: "/login",
    name: "Login",
    meta: { title: "Вход" },
    component: () => import("@/views/Login.vue"),
  },
  {
    path: "/posts/:id",
    name: "Post",
    meta: { title: "Пост" },
    component: () => import("@/views/Post.vue"),
  },
  {
    path: "/your-posts",
    name: "UserPosts",
    meta: { title: "Твои посты" },
    beforeEnter: authRequired,
    component: () => import("@/views/UserPosts.vue"),
  },
  {
    path: "/editor/:id",
    name: "Editor",
    meta: { title: "Редактор" },
    beforeEnter: authRequired,
    component: () => import("@/views/Editor.vue"),
  },
  {
    path: "/about",
    name: "About",
    meta: { title: "Об авторе" },
    component: () => import("@/views/About.vue"),
  },
  {
    path: "/*",
    redirect: "/",
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

router.beforeEach((to, from, next) => {
  document.title = to.meta.title;
  next();
});

export default router;
