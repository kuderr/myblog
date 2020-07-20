<template>
  <v-form ref="form" v-model="valid" :lazy-validation="true">
    <v-row>
      <v-col cols="12" md="4">
        <v-text-field
          v-model.trim="post.title"
          :rules="[(v) => !!v || 'Title is required']"
          label="Title"
          required
        ></v-text-field>
      </v-col>

      <v-col cols="12" md="8">
        <v-textarea
          v-model.trim="post.summary"
          :rules="[(v) => !!v || 'Summary is required']"
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
      <v-btn class="mt-2" block color="secondary" @click="savePost"
        >Сохранить</v-btn
      >
    </div>
  </v-form>
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
  Code,
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
    Code,
    HorizontalRule,
    Paragraph,
    HardBreak,
  ];
}
</script>
