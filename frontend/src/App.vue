<template>
  <div id="app">
    <Header />
    <Menu />

    <el-container>
      <el-main>
        <router-view></router-view>
      </el-main>
    </el-container>
  </div>
</template>

<script lang="ts">
import Header from "@/components/Header.vue";
import Menu from "@/components/Menu.vue";
import { Component, Prop, Vue } from "vue-property-decorator";

@Component({
  components: {
    Header,
    Menu
  }
})
export default class App extends Vue {}
</script>

<style >
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  width: 65%;
  margin: 30px auto;
}
</style>
