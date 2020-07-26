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
          <v-btn class="mt-2" block color="secondary" @click="savePost">Сохранить</v-btn>
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
  Paragraph,
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

import { addPost } from "../api";

@Component({
  components: {
    TiptapVuetify,
  },
})
export default class Posts extends Vue {
  private valid: boolean = false;
  private post: {
    title: string;
    summary: string;
    body: string;
    authorId: number;
  } = {
    title: "",
    summary: "",
    body: "",
    authorId: 1,
  };

  private savePost(): void {
    if (this.$refs.form.validate()) {
      addPost(this.post);
      this.$router.push("/");
    }
  }

  // tiptap extensions declare
  extensions = [
    History,
    Blockquote,
    Link,
    Underline,
    Strike,
    Italic,
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
    Bold,
    Link,
    CodeBlock,
    HorizontalRule,
    Paragraph,
    HardBreak,
  ];
}
</script>

<style lang="scss" >
.v-application code {
  background-color: black;
  color: white;
}
img {
  max-width: 100%;
  display: block;
  margin-left: auto;
  margin-right: auto;
}
</style>