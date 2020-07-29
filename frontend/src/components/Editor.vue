<template>
  <v-layout align-center justify-center>
    <v-flex xs12 sm10 md9>
      <v-form ref="form" v-model="valid" :lazy-validation="true">
        <v-row>
          <v-col cols="12" md="4">
            <v-text-field
              v-model.trim="post.title"
              :rules="[(v) => !!v || 'Title is required',
          v => v.length <= 50 || 'Title must be less than 50 characters']"
              :counter="50"
              label="Title"
              required
            ></v-text-field>
          </v-col>

          <v-col cols="12" md="8">
            <v-textarea
              v-model.trim="post.summary"
              :counter="255"
              :rules="[(v) => !!v || 'Summary is required',
                    v => v.length <= 255 || 'Summary must be less than 255 characters']"
              label="Summary"
              required
              auto-grow
              rows="1"
            ></v-textarea>
          </v-col>
        </v-row>
        <div>
          <tiptap-vuetify
            v-model="post.body"
            :extensions="extensions"
            placeholder="Текст поста"
            :toolbar-attributes="{ color: 'black', dark: true }"
          />
          <template>
            <v-btn class="mt-3" block color="secondary" @click="updatePost()">Сохранить</v-btn>
            <v-btn
              class="mt-1"
              block
              color="primary"
              v-if="post.published"
              @click="switchPublishedStatus()"
            >Отменить публикацию</v-btn>
            <v-btn
              v-else
              class="mt-1"
              block
              color="primary"
              @click="switchPublishedStatus()"
            >Публиковать</v-btn>
            <v-btn class="mt-1" block color="error" @click="deletePost()">Удалить</v-btn>
          </template>
        </div>
      </v-form>
    </v-flex>
  </v-layout>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

import {
  // component
  TiptapVuetify,
  // extensions
  Heading,
  Bold,
  Italic,
  Strike,
  Underline,
  CodeBlock,
  BulletList,
  OrderedList,
  ListItem,
  Link,
  Blockquote,
  HardBreak,
  HorizontalRule,
  History,
  Image,
} from "tiptap-vuetify";

import { getPost, updatePost, updatePublishedStatus, deletePost } from "../api";

@Component({
  components: {
    TiptapVuetify,
  },
})
export default class Editor extends Vue {
  private valid: boolean = false;
  private postId: string;
  private post: {
    title: string;
    summary: string;
    body: string;
    authorId: number;
    published?: boolean;
    dateCreated?: string;
    dateUpdated?: string;
    id?: number;
  } = {
    title: "",
    summary: "",
    body: "",
    authorId: 1,
  };

  async created() {
    this.postId = this.$router.currentRoute.params["id"];
    let res = await getPost(this.postId);
    this.post = res.data;
    this.post["img"] = "https://cdn.vuetifyjs.com/images/cards/docks.jpg";
  }

  private async updatePost() {
    if (this.$refs.form.validate()) {
      console.log(this.post);
      await updatePost(this.post);
    }
  }

  private switchPublishedStatus(): void {
    this.post.published = !this.post.published;
    updatePublishedStatus(this.postId, this.post.published);
  }

  private deletePost(): void {
    deletePost(this.postId);
    this.$router.push("/your-posts");
  }

  // tiptap extensions declare
  extensions = [
    History,
    Bold,
    Italic,
    Underline,
    Strike,
    ListItem,
    BulletList,
    OrderedList,
    Image,
    [
      Heading,
      {
        options: {
          levels: [1, 2, 3],
        },
      },
    ],
    CodeBlock,
    Blockquote,
    Link,
    HorizontalRule,
    HardBreak,
  ];
}
</script>

<style lang="scss" >
.v-application code {
  background-color: black;
  color: white;
  display: block;
  margin-left: auto;
  margin-right: auto;
}
img {
  max-width: 100%;
  display: block;
  margin-left: auto;
  margin-right: auto;
}
</style>