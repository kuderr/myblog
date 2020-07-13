<template>
  <v-app>
    <Drawer />

    <v-content>
      <v-container fluid>
        <router-view></router-view>
      </v-container>
    </v-content>

    <!-- Ошибки и сообщнеия: -->
    <!-- <template v-if="error">
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
    </template> -->
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
export default class App extends Vue {}
</script>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
</style>
