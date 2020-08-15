<template>
  <v-app>
    <Drawer />

    <v-main>
      <v-container fluid>
        <nuxt></nuxt>
      </v-container>
    </v-main>

    <v-fab-transition>
      <v-btn small fab fixed bottom right @click="switchColorMode()">
        <v-icon>invert_colors</v-icon>
      </v-btn>
    </v-fab-transition>

    <!-- Ошибки и сообщнеия: -->
    <template v-if="error">
      <v-snackbar
        :multi-line="true"
        :timeout="3000"
        @input="closeError"
        :value="true"
        color="error"
      >
        {{ error }}
        <v-btn dark text @click="closeError">Close</v-btn>
      </v-snackbar>
    </template>

    <template v-if="message">
      <v-snackbar
        :multi-line="true"
        :timeout="3000"
        @input="closeMessage"
        :value="true"
        color="primary"
      >
        {{ message }}
        <v-btn dark text @click="closeMessage">Close</v-btn>
      </v-snackbar>
    </template>
    <!-- ------------------------ -->
  </v-app>
</template>

<script>
import Drawer from '@/components/Drawer.vue'
import { isValidToken } from '@/utils'

export default {
  name: 'Main',
  components: {
    Drawer,
  },
  async mounted() {
    this.$vuetify.theme.dark = localStorage.darkMode === 'true' ? true : false
    await this.$store.dispatch('fetchPosts')
    console.log(this.$auth.$storage)
    const token = this.$auth.getToken('token') || ''
    if (token) {
      const tokenParts = token.split('.')
      const body = JSON.parse(atob(tokenParts[1]))
      console.log(body)
      this.$auth.setUser(body)
    }
  },
  computed: {
    error() {
      return this.$store.state.shared.error
    },
    message() {
      return this.$store.state.shared.message
    },
  },
  methods: {
    closeError() {
      this.$store.dispatch('clearError')
    },

    closeMessage() {
      this.$store.dispatch('clearMessage')
    },

    switchColorMode() {
      this.$vuetify.theme.dark = !this.$vuetify.theme.dark
      if (process.browser) {
        localStorage.darkMode = this.$vuetify.theme.dark
      }
    },
  },
}
</script>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
</style>
