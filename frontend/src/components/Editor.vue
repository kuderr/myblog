<template>
  <v-layout align-center justify-center>
    <v-flex xs12 sm10 md9>
      <v-form ref="form" v-model="valid" :lazy-validation="true">
        <v-row>
          <v-col cols="12" md="4">
            <v-text-field
              v-model.trim="post.title"
              :rules="[v => !!v || 'Title is required',
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
              :rules="[v => !!v || 'Summary is required',
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
            @keydown.meta.s.prevent="updatePost()"
            @keydown.ctrl.s.prevent="updatePost()"
            :toolbar-attributes="{ color: 'black', dark: true }"
            placeholder="Текст поста"
          />
          <template>
            <v-tooltip right>
              <template v-slot:activator="{ on, attrs }">
                <v-btn
                  class="mt-4"
                  v-bind="attrs"
                  v-on="on"
                  block
                  color="secondary"
                  :loading="saveLoader"
                  @click="updatePost()"
                >Сохранить</v-btn>
              </template>
              <span>Ctrl+S (Cmd+S)</span>
            </v-tooltip>
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

import { updatePost } from "../api";

@Component({
  components: {
    TiptapVuetify,
  },
})
export default class Editor extends Vue {
  // loaders
  private saveLoader: boolean = false;
  // form data
  private valid: boolean = false;

  mounted() {
    let postId = this.$router.currentRoute.params["id"];
    this.$store.dispatch("fetchPost", postId);
  }

  get post() {
    return this.$store.state.posts.currentPost;
  }

  // TODO
  private async updatePost() {
    if (this.$refs.form.validate()) {
      this.saveLoader = true;
      await updatePost(this.post);
      let images = this.post.body.match(/"data:image[^ >]*/g);
      this.saveLoader = false;
    }
  }

  private switchPublishedStatus(): void {
    this.post.published = !this.post.published;
    this.$store.dispatch("updatePublishedStatus", this.post);
  }

  private deletePost() {
    this.$store.dispatch("deletePost", this.post.id);
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