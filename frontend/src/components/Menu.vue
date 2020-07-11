<template>
  <div id="menu">
    <div>
      <el-menu
        :default-active="curPath"
        class="el-menu"
        mode="horizontal"
        @select="handleSelect"
        :router="true"
      >
        <el-menu-item index="/">Главная</el-menu-item>
        <el-menu-item index="/about">Об авторе</el-menu-item>
      </el-menu>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";

@Component
export default class Menu extends Vue {
  curPath: string = localStorage.curPath || "/";

  handleSelect(key: string, keyPath: string): void {
    localStorage.curPath = key;
  }
}
</script>


