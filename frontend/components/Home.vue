<template>
  <div class="d-flex flex-wrap mb-6 justify-center">
    <v-progress-linear v-if="posts.length === 0" stream indeterminate color="primary"></v-progress-linear>
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
        <span class="title font-weight-bold">{{ post.title }}</span>
      </v-card-title>

      <v-card-text class="headline font-weight-bold">
        {{
        post.summary
        }}
      </v-card-text>
    </v-card>
  </div>
</template>

<script>
export default {
  name: 'Home',
  computed: {
    posts() {
      return this.$store.state.posts.posts
    },
  },
  components: {
    VBoilerplate: {
      functional: true,

      render(h, { data, props, children }) {
        return h(
          'v-skeleton-loader',
          {
            ...data,
            props: {
              boilerplate: true,
              elevation: 3,
              ...props,
            },
          },
          children
        )
      },
    },
  },
  head() {
    return {
      title: 'Посты',
      meta: [
        {
          name: 'description',
          content:
            'Привет, я Дима. В своём блоге пишу об интересующих меня вещах. Рад делиться опытом и мыслями, а также узнавать постоянно что-то новое.',
        },
        { name: 'keywords', content: 'блог, blog, tech' },
      ],
    }
  },
}
</script>
