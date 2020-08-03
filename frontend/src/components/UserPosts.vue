<template>
  <div class="d-flex flex-wrap mb-6 justify-center">
    <v-btn class="mb-4" block color="success" @click="newPost()">Новый пост</v-btn>
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
import { addPost, getUserPosts } from "../api";

@Component
export default class UserPosts extends Vue {
  posts = [];
  post = {
    title: "Новый пост",
    summary: "Новый пост",
    body: "",
    authorId: 1,
  };

  async newPost() {
    let res = await addPost(this.post);
    this.$router.push("/editor/" + res.data.postId);
  }

  async created() {
    let res = await getUserPosts(1);
    res.data.forEach((element) => {
      element.color = "#26c6da";
      element.img = "https://cdn.vuetifyjs.com/images/cards/mountain.jpg";
      element.likes = 256;
      element.author_name = "Evan You";
      element.author_avatar =
        "https://avataaars.io/?avatarStyle=Transparent&topType=ShortHairShortCurly&accessoriesType=Prescription02&hairColor=Black&facialHairType=Blank&clotheType=Hoodie&clotheColor=White&eyeType=Default&eyebrowType=DefaultNatural&mouthType=Default&skinColor=Light";
    });
    this.posts = res.data;
  }
}
</script>
