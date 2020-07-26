<template>
  <div id="post">
    <v-layout align-center justify-center>
      <v-flex xs12 sm10 md9>
        <v-card class="mx-auto">
          <v-img class="white--text align-end" height="200px" :src="post.img"></v-img>
          <v-card-title class="headline font-weight-bold">{{post.title}}</v-card-title>

          <v-card-text class="text--primary">
            <div v-html="post.body"></div>
          </v-card-text>
        </v-card>
      </v-flex>
    </v-layout>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

import { getPost } from "../api";

@Component
export default class Post extends Vue {
  post_id: string;
  post: object = {};

  async created() {
    this.post_id = this.$router.currentRoute.params["id"];
    let res = await getPost(this.post_id);
    this.post = res.data;
    this.post["img"] = "https://cdn.vuetifyjs.com/images/cards/docks.jpg";

    console.log(this.post);
  }
}
</script>

<style lang="scss" >
.v-application code {
  background-color: black;
  color: white;
}
img {
  max-width: 100%;
  display: block;
  margin-left: auto;
  margin-right: auto;
}
</style>