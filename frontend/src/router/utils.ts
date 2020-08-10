import store from "@/store";
import { isValidToken } from "@/utils";

export function authRequired(to, from, next) {
  if (!isValidToken(store.state.user.token)) {
    next("/");
  } else {
    next();
  }
}

export async function fetchMeta(to, from, next) {
  await store.dispatch("fetchPost", to.params.id);

  createMetaTags();

  next();
}

function createMetaTags() {
  let post = store.state.posts.currentPost;
  let metaTags = [
    { name: "description", content: post.summary },
    { name: "author", content: "Dmitriy Kudryavtsev" },
    { name: "article:published_time", content: post.dateCreated },
    // { name: "keywords", content: "" }, // place for tags

    {
      name: "og:url",
      content: location.origin + "/posts/" + post.id,
    },
    { name: "og:title", content: post.title },
    { name: "og:type", content: "article" },
    { name: "og:site_name", content: "Kuder Blog" },
    { name: "og:locale", content: "ru_RU" },
    { name: "og:description", content: post.summary },
    { name: "og:image", content: post.img },

    { name: "twitter:card", content: "summary_large_image" },
    {
      name: "twitter:url",
      content: location.origin + "/posts/" + post.id,
    },
    { name: "twitter:title", content: post.title },
    { name: "twitter:description", content: post.summary },
    { name: "twitter:creator", content: "@kuderrr" },
    { name: "twitter:site", content: "@kuderrr" },
    { name: "twitter:image", content: post.img },
  ];

  for (let tag of metaTags) {
    var meta = document.createElement("meta");
    meta.name = tag.name;
    meta.content = tag.content;
    document.head.prepend(meta);
  }
}
