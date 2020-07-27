<template>
  <div class="d-flex flex-wrap mb-6 justify-center">
    <v-card
      v-for="post in posts"
      :key="post.id"
      class="ma-2 align-self-auto"
      :color="post.color"
      dark
      width="400"
      :to="`/posts/${post.id}`"
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
        gotcha edge cases are type invalid as well."`,
    },
    {
      id: 2,
      color: "#1F7087",
      img: "https://cdn.vuetifyjs.com/images/cards/foster.jpg",
      title: "Supermodel",
      summary: "Foster the People",
    },
    {
      id: 3,
      color: "#952175",
      img: "https://cdn.vuetifyjs.com/images/cards/halcyon.png",
      title: "Halcyon Days",
      summary: "Ellie Goulding",
    },
    {
      id: 4,
      color: "#952175",
      img: "https://cdn.vuetifyjs.com/images/cards/halcyon.png",
      title: "Halcyon Days",
      summary: "Ellie Goulding",
    },
  ];

  test() {
    console.log("kek");
  }

  async created() {
    let res = await getPosts();
    res.data.forEach((element) => {
      element.color = "#26c6da";
      element.img = "https://cdn.vuetifyjs.com/images/cards/mountain.jpg";
      element.likes = 256;
      element.author_name = "Evan You";
      element.author_avatar =
        "https://avataaars.io/?avatarStyle=Transparent&topType=ShortHairShortCurly&accessoriesType=Prescription02&hairColor=Black&facialHairType=Blank&clotheType=Hoodie&clotheColor=White&eyeType=Default&eyebrowType=DefaultNatural&mouthType=Default&skinColor=Light";
    });
    this.posts = [...res.data, ...this.posts];
  }
}
</script>
