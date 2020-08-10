<template>
  <div id="post">
    <v-layout align-center justify-center>
      <v-flex xs12 sm10 md8>
        <v-card class="mx-auto">
          <v-img height="200px" :src="post.img"></v-img>
          <v-card-title class="headline font-weight-bold d-flex justify-space-between">
            {{ post.title }}
            <time class="dateCreated">{{ postDateFormatted }}</time>
          </v-card-title>

          <v-card-text class="text--primary">
            <div class="body-1" v-html="post.body"></div>
          </v-card-text>
        </v-card>
      </v-flex>
    </v-layout>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import { MetaInfo } from "vue-meta";

@Component({
  metaInfo(): MetaInfo {
    return {
      meta: [
        { name: "description", content: this.post.summary },
        { name: "author", content: "Dmitriy Kudryavtsev" },
        { name: "article:published_time", content: this.post.dateCreated },
        // { name: "keywords", content: "" }, // place for tags

        {
          name: "og:url",
          content: location.origin + "/posts/" + this.post.id,
        },
        { name: "og:title", content: this.post.title },
        { name: "og:type", content: "article" },
        { name: "og:site_name", content: "Kuder Blog" },
        { name: "og:locale", content: "ru_RU" },
        { name: "og:description", content: this.post.summary },
        { name: "og:image", content: this.post.img },

        // { name: "twitter:card", content: "summary_large_image" },
        {
          name: "twitter:url",
          content: location.origin + "/posts/" + this.post.id,
        },
        { name: "twitter:title", content: this.post.title },
        { name: "twitter:description", content: this.post.summary },
        { name: "twitter:creator", content: "@kuderrr" },
        { name: "twitter:site", content: "@kuderrr" },
        { name: "twitter:image", content: this.post.img },
      ],
    };
  },
})
export default class PostDetail extends Vue {
  private options = {
    year: "numeric",
    month: "numeric",
    day: "numeric",
    // hour: "numeric",
    // minute: "numeric",
  };

  mounted() {
    let postId = this.$router.currentRoute.params["id"];
    this.$store.dispatch("fetchPost", postId);
  }

  get post() {
    let post = this.$store.state.posts.currentPost;
    return post;
  }

  get postDateFormatted() {
    let timestamp = Date.parse(this.post.dateCreated);
    let postDate = new Date(timestamp);
    return postDate.toLocaleString("ru", this.options);
  }
}
</script>

<style lang="scss">
.dateCreated {
  color: #6a737d;
  font-weight: 400;
  font-size: 16px;
}
.v-application code {
  background-color: black;
  color: white;
  display: block;
  margin-left: auto;
  margin-right: auto;
}
img:not(.avatar) {
  max-width: 100%;
  padding: 5px 0;
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
