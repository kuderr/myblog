<template>
  <div class="d-flex flex-wrap mb-6 justify-center">
    <v-card
      v-for="post in posts"
      :key="post.id"
      class="ma-2 align-self-auto"
      color="#26c6da"
      dark
      width="400"
      :to="`/posts/${post.id}`"
    >
      <v-img
        src="https://cdn.vuetifyjs.com/images/cards/mountain.jpg"
        height="194"
      ></v-img>

      <v-card-title>
        <span class="title font-weight-bold">{{ post.title }}</span>
      </v-card-title>

      <v-card-text class="headline font-weight-bold">{{
        post.summary
      }}</v-card-text>
    </v-card>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import { State, Action, Getter } from "vuex-class";
import { PostsState, Post } from "../store/posts/types";
const namespace: string = "posts";

@Component
export default class Home extends Vue {
  @State("posts")
  postsState: PostsState;
  @Action("fetchData", { namespace })
  fetchData: any;
  @Getter("posts", { namespace })
  posts: Post[];

  mounted() {
    this.fetchData();
  }
}
</script>
