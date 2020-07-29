import Vue from "vue";
import VueRouter, { RouteConfig } from "vue-router";

Vue.use(VueRouter);

const routes: RouteConfig[] = [
  {
    path: "/",
    name: "Home",
    component: () => import("@/views/Home.vue"),
  },
  {
    path: "/posts/:id",
    name: "Post",
    component: () => import("@/views/Post.vue"),
  },
  {
    path: "/your-posts",
    name: "UserPosts",
    component: () => import("@/views/UserPosts.vue"),
  },
  {
    path: "/editor/:id",
    name: "Editor",
    component: () => import("@/views/Editor.vue"),
  },
  {
    path: "/about",
    name: "About",
    component: () => import("@/views/About.vue"),
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

export default router;
