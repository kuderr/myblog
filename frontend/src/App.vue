<template>
  <div id="app">
    <div class="demo-type">
      <el-avatar :size="125">
        <img
          src="https://avatars0.githubusercontent.com/u/39552217?s=460&u=415aebf4249492578b8b2af663f5e5473c704dd4&v=4"
        />
      </el-avatar>
    </div>
    <div>
      <el-link
        v-for="soc in social"
        :key="soc.url"
        :href="soc.url"
        target="_blank"
      >
        <i :class="soc.icon"></i>
      </el-link>
    </div>
    <el-menu
      :default-active="curPath"
      class="el-menu"
      mode="horizontal"
      @select="handleSelect"
      :router="true"
    >
      <el-menu-item index="/">Главная</el-menu-item>
      <el-menu-item index="/posts">Посты</el-menu-item>
      <el-menu-item index="/about">Об авторе</el-menu-item>
    </el-menu>

    <el-container>
      <el-main> <router-view></router-view> </el-main>
    </el-container>
  </div>
</template>

<script>
export default {
  name: "app",
  data() {
    return {
      curPath: localStorage.curPath || "/",
      social: [
        {
          url: "https://github.com/kuderr",
          icon: "el-icon-fa-github",
        },
        {
          url: "https://twitter.com/kuderrr",
          icon: "el-icon-fa-twitter",
        },
        {
          url: "https://vk.com/kuderrr",
          icon: "el-icon-fa-vk",
        },
        {
          url: "https://www.instagram.com/kuderr_/",
          icon: "el-icon-fa-instagram",
        },
        {
          url: "mailto:dakudryavcev@gmail.com",
          icon: "el-icon-fa-envelope",
        },
      ],
    };
  },
  methods: {
    handleSelect(key, keyPath) {
      localStorage.curPath = key;
    },
  },
};
</script>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  width: 60%;
  margin: 30px auto;
}

i {
  margin: 0 10px;
}
</style>
