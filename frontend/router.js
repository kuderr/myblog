import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: () => import('@/components/Home.vue').then((m) => m.default),
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/components/Login.vue').then((m) => m.default),
  },
  {
    path: '/posts/:id',
    name: 'Post',
    component: () => import('@/components/Post.vue').then((m) => m.default),
  },
  {
    path: '/your-posts',
    name: 'UserPosts',
    component: () =>
      import('@/components/UserPosts.vue').then((m) => m.default),
  },
  {
    path: '/editor/:id',
    name: 'Editor',
    component: () => import('@/components/Editor.vue').then((m) => m.default),
  },
  {
    path: '/about',
    name: 'About',
    component: () => import('@/components/About.vue').then((m) => m.default),
  },
  {
    path: '/*',
    redirect: '/',
  },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
})

export function createRouter() {
  return router
}
