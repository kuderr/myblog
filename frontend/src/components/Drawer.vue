<template>
  <div id="drawer">
    <v-app-bar app clipped-left dark dense flat>
      <v-app-bar-nav-icon @click.stop="switchDrawer()" />
      <v-toolbar-title>Kuder Blog</v-toolbar-title>
    </v-app-bar>

    <v-navigation-drawer v-model="drawer" app clipped dark>
      <v-list dense nav :class="isUserLoggedIn ? 'py-0' : 'py-3'">
        <template v-if="isUserLoggedIn">
          <v-list-item two-line :class="miniVariant && 'px-0'">
            <v-list-item-avatar>
              <img class="avatar" :src="user.avatar" />
            </v-list-item-avatar>

            <v-list-item-content>
              <v-list-item-title>{{ user.fullName }}</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
          <v-divider></v-divider>
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

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import { isValidToken } from "../utils";

@Component
export default class Drawer extends Vue {
  drawer: boolean = localStorage.drawer === "true" ? true : false;

  get items() {
    let items = [
      { title: "Главная", icon: "dashboard", path: "/" },
      { title: "Об авторе", icon: "info", path: "/about" },
    ];

    if (this.isUserLoggedIn) {
      items.splice(1, 0, {
        title: "Твои посты",
        icon: "edit",
        path: "/your-posts",
      });
    }

    return items;
  }

  get isUserLoggedIn() {
    return isValidToken(this.$store.state.user.jwt.token);
  }

  user = {
    fullName: "Dima Kudryavtsev",
    avatar:
      "https://avatars0.githubusercontent.com/u/39552217?s=460&u=415aebf4249492578b8b2af663f5e5473c704dd4&v=4",
  };
  miniVariant = false;

  switchDrawer(): void {
    this.drawer = !this.drawer;
    localStorage.drawer = this.drawer;
  }

  logout(): void {
    this.$store.dispatch("logout");
    this.$router.push("/");
  }
}
</script>
