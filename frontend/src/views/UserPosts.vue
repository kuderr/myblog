<template>
  <div class="userPosts">
    <UserPosts />
  </div>
</template>

<script>
import UserPosts from "@/components/UserPosts.vue";

export default {
  name: "UserPostsView",
  components: {
    UserPosts,
  },
};
</script>
