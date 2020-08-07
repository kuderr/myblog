<template>
  <div class="d-flex flex-wrap mb-6 justify-center">
    <v-btn class="mb-4" block color="success" @click="newPost()">Создать пост</v-btn>
    <v-card
      v-for="post in posts"
      :key="post.id"
      class="ma-2 align-self-auto"
      :color="post.color"
      dark
      width="400"
      :to="`/editor/${post.id}`"
    >
      <v-img :src="post.img" height="194"></v-img>

      <v-card-title>
        <span class="title font-weight-bold">{{post.title}}</span>
      </v-card-title>

      <v-card-text class="headline font-weight-bold">{{post.summary}}</v-card-text>
    </v-card>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

@Component
export default class UserPosts extends Vue {
  post = {
    title: "Новый пост",
    summary: "Новый пост",
    authorId: this.user.id,
  };

  get user() {
    return this.$store.state.user.user;
  }

  async newPost() {
    let res = await this.$store.dispatch("addPost", this.post);
    this.$router.push("/editor/" + res.data.postId);
  }

  get posts() {
    return this.$store.state.user.userPosts;
  }
}
</script>
