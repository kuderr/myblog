<template>
  <div>
    <tiptap-vuetify
      v-model="content"
      :extensions="extensions"
      placeholder="Текст поста"
      :toolbar-attributes="{ color: 'black', dark: true }"
    />
    <v-btn class="mt-2" block color="secondary">Сохранить</v-btn>
  </div>
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

@Component({
  components: {
    TiptapVuetify,
  },
})
export default class Posts extends Vue {
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
  content = "";
}
</script>
