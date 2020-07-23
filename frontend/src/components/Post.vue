<template>
  <div id="post">
    <h1>This is post page for {{post_id}}</h1>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

@Component
export default class Post extends Vue {
  post_id: string;

  created() {
    this.post_id = this.$router.currentRoute.params["id"];
  }
}
</script>
