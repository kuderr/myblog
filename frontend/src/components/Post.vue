<template>
  <div id="post">
    <v-layout align-center justify-center>
      <v-flex xs12 sm10 md8>
        <v-card class="mx-auto">
          <v-img height="200px" :src="post.img"></v-img>
          <v-card-title class="headline font-weight-bold">
            <strong class="blue--text text--lighten-2">{{post.title}}</strong>
          </v-card-title>

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
  }
}
</script>

<style lang="scss" >
.v-application code {
  background-color: black;
  color: white;
  display: block;
  margin-left: auto;
  margin-right: auto;
}
img {
  max-width: 100%;
  display: block;
  margin-left: auto;
  margin-right: auto;
}

ul,
ol {
  margin-bottom: 10px;
}

h1 {
  margin-top: 30px !important;
}
h2 {
  margin-top: 20px !important;
}
h3 {
  margin-top: 10px !important;
}

blockquote {
  border-left: 0.25em solid #dfe2e5;
  color: #6a737d;
  padding-left: 1em;
  margin: 20px 0 !important;
}

code {
  padding: 0 4px !important;
  margin: 0 5px !important;
  background-color: black;
  color: white;
}

pre code {
  padding: 8px !important;
  margin: 0 5px !important;
}

code:before,
code:after {
  content: none !important;
  letter-spacing: initial !important;
}

p {
  margin-top: 5px !important;
  margin-bottom: 5px !important;
  min-height: 1rem;
}
</style>