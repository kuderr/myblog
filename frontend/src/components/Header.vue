<template>
  <div id="header">
    <div>
      <el-avatar :size="125">
        <img
          src="https://avatars0.githubusercontent.com/u/39552217?s=460&u=415aebf4249492578b8b2af663f5e5473c704dd4&v=4"
        />
      </el-avatar>
    </div>
    <div>
      <el-link v-for="soc in social" :key="soc.url" :href="soc.url" target="_blank">
        <i :class="soc.icon"></i>
      </el-link>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";

@Component
export default class Header extends Vue {
  social: { url: string; icon: string }[] = [
    {
      url: "https://github.com/kuderr",
      icon: "el-icon-fa-github"
    },
    {
      url: "https://twitter.com/kuderrr",
      icon: "el-icon-fa-twitter"
    },
    {
      url: "https://vk.com/kuderrr",
      icon: "el-icon-fa-vk"
    },
    {
      url: "https://www.instagram.com/kuderr_/",
      icon: "el-icon-fa-instagram"
    },
    {
      url: "mailto:dakudryavcev@gmail.com",
      icon: "el-icon-fa-envelope"
    }
  ];
}
</script>

<style scoped>
i {
  margin: 0 10px;
}
</style>
