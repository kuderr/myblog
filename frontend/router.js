import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    meta: { title: 'Главная' },
    component: () => import('@/components/Home.vue').then((m) => m.default),
  },
  {
    path: '/login',
    name: 'Login',
    meta: { title: 'Вход' },
    component: () => import('@/components/Login.vue').then((m) => m.default),
  },
  {
    path: '/posts/:id',
    name: 'Post',
    meta: { title: 'Пост' },
    component: () => import('@/components/Post.vue').then((m) => m.default),
  },
  {
    path: '/your-posts',
    name: 'UserPosts',
    meta: { title: 'Твои посты' },
    component: () =>
      import('@/components/UserPosts.vue').then((m) => m.default),
  },
  {
    path: '/editor/:id',
    name: 'Editor',
    meta: { title: 'Редактор' },
    component: () => import('@/components/Editor.vue').then((m) => m.default),
  },
  {
    path: '/about',
    name: 'About',
    meta: { title: 'Об авторе' },
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

router.beforeEach((to, from, next) => {
  if (process.client) {
    document.title = to.meta.title
  }
  next()
})

export function createRouter() {
  return router
}
