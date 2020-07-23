<template>
  <div :class="`d-flex flex-wrap mb-6`">
    <v-card
      v-for="post in posts"
      :key="post.id"
      class="ma-2 align-self-start"
      :color="post.color"
      dark
      max-width="400"
      :to="`/posts/${post.id}`"
    >
      <v-img :src="post.img" height="194"></v-img>

      <v-card-title>
        <span class="title font-weight-bold">{{post.title}}</span>
      </v-card-title>

      <v-card-text class="headline font-weight-bold">{{post.summary}}</v-card-text>

      <v-card-actions>
        <v-list-item class="grow">
          <v-list-item-avatar color="grey darken-3">
            <v-img class="elevation-6" :src="post.author_avatar"></v-img>
          </v-list-item-avatar>

          <v-list-item-content>
            <v-list-item-title>{{post.author_name}}</v-list-item-title>
          </v-list-item-content>

          <v-row align="center" justify="end">
            <v-icon class="mr-1">mdi-heart</v-icon>
            <span class="subheading mr-2">{{post.likes}}</span>
            <span class="mr-1">·</span>
            <v-icon class="mr-1">mdi-share-variant</v-icon>
          </v-row>
        </v-list-item>
      </v-card-actions>
    </v-card>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

import { getPosts } from "../api";

@Component
export default class Home extends Vue {
  posts = [
    {
      id: 1,
      color: "#26c6da",
      img: "https://cdn.vuetifyjs.com/images/cards/mountain.jpg",
      likes: 256,
      author_name: "Evan You",
      author_avatar:
        "https://avataaars.io/?avatarStyle=Transparent&topType=ShortHairShortCurly&accessoriesType=Prescription02&hairColor=Black&facialHairType=Blank&clotheType=Hoodie&clotheColor=White&eyeType=Default&eyebrowType=DefaultNatural&mouthType=Default&skinColor=Light",
      title: "Twitter",
      summary: `Turns out semicolon-less style is easier and safer in TS because most
        gotcha edge cases are type invalid as well."`
    },
    {
      id: 2,
      color: "#1F7087",
      img: "https://cdn.vuetifyjs.com/images/cards/foster.jpg",
      title: "Supermodel",
      summary: "Foster the People"
    },
    {
      id: 3,
      color: "#952175",
      img: "https://cdn.vuetifyjs.com/images/cards/halcyon.png",
      title: "Halcyon Days",
      summary: "Ellie Goulding"
    },
    {
      id: 4,
      color: "#952175",
      img: "https://cdn.vuetifyjs.com/images/cards/halcyon.png",
      title: "Halcyon Days",
      summary: "Ellie Goulding"
    }
  ];

  created() {
    let posts;
    getPosts().then(res => {
      posts = res.data;
      this.posts = [...this.posts, ...posts];
    });
  }
}
</script>
