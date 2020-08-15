<template>
  <div id="drawer">
    <v-app-bar app clipped-left dark dense flat>
      <v-app-bar-nav-icon @click.stop="switchDrawer()" />
      <v-toolbar-title>Kuder Blog</v-toolbar-title>
    </v-app-bar>

    <v-navigation-drawer v-model="drawer" app clipped dark>
      <v-list dense nav :class="isUserLoggedIn ? 'py-0' : 'py-3'">
        <template v-if="isUserLoggedIn">
          <v-list-item two-line class="px-1">
            <v-list-item-avatar>
              <img v-if="user.avatar" class="avatar" :src="user.avatar" />
              <v-icon v-else large>account_circle</v-icon>
            </v-list-item-avatar>

            <v-list-item-content class="px-1">
              <v-list-item-title>{{ user.fullName }}</v-list-item-title>
              <v-list-item-subtitle>{{ user.role }}</v-list-item-subtitle>
            </v-list-item-content>
          </v-list-item>
          <v-divider class="mb-1"></v-divider>
        </template>

        <v-list-item v-for="item in items" :key="item.title" :to="item.path">
          <v-list-item-icon>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-icon>

          <v-list-item-content>
            <v-list-item-title>{{ item.title }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
      <template v-slot:append>
        <div class="pa-2">
          <v-btn block @click="logout()" v-if="isUserLoggedIn">Logout</v-btn>
        </div>
      </template>
    </v-navigation-drawer>
  </div>
</template>

<script>
import { isValidToken } from '../utils'

export default {
  name: 'Drawer',
  data() {
    return {
      drawer: () => {
        if (process.browser) {
          return localStorage.drawer === 'true' ? true : false
        }
      },
    }
  },
  computed: {
    items() {
      let items = [
        { title: 'Главная', icon: 'dashboard', path: '/' },
        { title: 'Об авторе', icon: 'info', path: '/about' },
      ]

      if (this.isUserLoggedIn) {
        items.splice(1, 0, {
          title: 'Твои посты',
          icon: 'edit',
          path: '/your-posts',
        })
      }

      return items
    },
    isUserLoggedIn() {
      if (process.browser) {
        console.log(localStorage.token)
        return isValidToken(localStorage.token)
      }
    },
    user() {
      return this.$store.state.user.user
    },
  },
  methods: {
    switchDrawer() {
      this.drawer = !this.drawer
      if (process.browser) {
        localStorage.drawer = this.drawer
      }
    },

    logout() {
      this.$store.dispatch('logout')
      if (
        this.$router.currentRoute.fullPath !== '/' &&
        this.$router.currentRoute.fullPath !== '/about'
      ) {
        this.$router.push('/')
      }
    },
  },
}
</script>
