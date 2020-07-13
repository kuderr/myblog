<template>
  <div id="posts">
    <v-row>
      <v-col cols="4">
        <v-textarea filled auto-grow v-model="postText"></v-textarea>
      </v-col>
      <v-col cols="8">
        <div v-html="postText"></div>
      </v-col>
    </v-row>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

@Component
export default class Posts extends Vue {
  postText: string = "";
}
</script>
