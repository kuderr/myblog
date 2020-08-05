<template>
  <v-app>
    <Drawer />

    <v-content>
      <v-container fluid>
        <router-view></router-view>
      </v-container>
    </v-content>

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

<script lang="ts">
import Drawer from "@/components/Drawer.vue";
import { Component, Vue } from "vue-property-decorator";

@Component({
  components: {
    Drawer,
  },
})
export default class App extends Vue {
  get error() {
    return this.$store.state.shared.error;
  }
  closeError() {
    this.$store.dispatch("clearError");
  }

  get message() {
    return this.$store.state.shared.message;
  }
  closeMessage() {
    this.$store.dispatch("clearMessage");
  }

  switchColorMode() {
    this.$vuetify.theme.dark = !this.$vuetify.theme.dark;
    localStorage.darkMode = this.$vuetify.theme.dark;
  }

  mounted() {
    this.$vuetify.theme.dark = localStorage.darkMode === "true" ? true : false;
    this.$store.dispatch("fetchPosts");
    this.$store.dispatch("fetchUserPosts", 1);
  }
}
</script>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
</style>
