<template>
  <div id="post">
    <v-layout align-center justify-center>
      <v-flex xs12 sm10 md8>
        <v-skeleton-loader
          :loading="loading"
          transition="fade-transition"
          type="card"
        >
          <v-card class="mx-auto" elevation="3">
            <v-img height="200px" :src="post.img"></v-img>
            <v-card-title
              class="headline font-weight-bold d-flex justify-space-between"
            >
              {{ post.title }}
              <div class="postInfo">
                <div>
                  <time>{{ postDateFormatted }}</time>
                </div>
                <div class="views">
                  <v-icon small> mdi-eye </v-icon>
                  <span>{{ post.views }}</span>
                </div>
              </div>
            </v-card-title>

            <v-card-text class="text--primary">
              <div class="body-1" v-html="post.body"></div>
            </v-card-text>
          </v-card>
        </v-skeleton-loader>
      </v-flex>
    </v-layout>
  </div>
</template>

<script>
import { mapState } from 'vuex'

export default {
  name: 'Post',
  data() {
    return {
      options: {
        year: 'numeric',
        month: 'numeric',
        day: 'numeric',
      },
      loading: null,
    }
  },
  head() {
    return {
      title: this.post.title,
      meta: [
        { name: 'description', content: this.post.summary },
        { name: 'author', content: 'Dmitriy Kudryavtsev' },
        { name: 'article:published_time', content: this.post.dateCreated },
        // { name: "keywords", content: "" }, // place for tags

        {
          name: 'og:url',
          content: 'https://kuderblog.com/posts/' + this.post.id,
        },
        { name: 'og:title', content: this.post.title },
        { name: 'og:type', content: 'article' },
        { name: 'og:site_name', content: 'Kuder Blog' },
        { name: 'og:locale', content: 'ru_RU' },
        { name: 'og:description', content: this.post.summary },
        { name: 'og:image', content: this.post.img },

        { name: 'twitter:card', content: 'summary_large_image' },
        { name: 'twitter:image', content: this.post.img },
        {
          name: 'twitter:url',
          content: 'https://kuderblog.com/posts/' + this.post.id,
        },
        { name: 'twitter:title', content: this.post.title },
        { name: 'twitter:description', content: this.post.summary },
        { name: 'twitter:creator', content: '@kuderrr' },
        { name: 'twitter:site', content: '@kuderrr' },
      ],
    }
  },
  computed: mapState({
    post: (state) => state.posts.currentPost,
    postDateFormatted() {
      let timestamp = Date.parse(this.post.dateCreated)
      let postDate = new Date(timestamp)
      return postDate.toLocaleString('ru', this.options)
    },
  }),
  async created() {
    this.loading = true
    let postId = this.$router.currentRoute.params['id']
    await this.$store.dispatch('fetchPost', postId)
    this.loading = false
  },
}
</script>

<style lang="scss">
.postInfo {
  color: #6a737d;
  font-weight: 400;
  font-size: 14px;
}

.views {
  float: right;
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

pre,
code {
  white-space: pre-wrap;
}
</style>
