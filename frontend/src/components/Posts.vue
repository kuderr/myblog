<template>
  <div id="posts">
    <el-card>
      <div class="container">
        <img
          src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
          class="container__image"
        />
        <div class="container__text">
          <h2>Ttile</h2>
          <p>
            Lorem ipsum dolor sit amet consectetur adipisicing elit. Ducimus
            dolore animi voluptatum veniam itaque sit aspernatur veritatis.
            Ullam enim, incidunt dolores architecto illum, dicta explicabo
            minus similique debitis numquam consectetur.
          </p>
        </div>
      </div>
    </el-card>
    <el-card>
      <div class="container">
        <img
          src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
          class="container__image"
        />
        <div class="container__text">
          <h2>Ttile</h2>
          <p>
            Lorem ipsum dolor sit amet consectetur adipisicing elit. Ducimus
            dolore animi voluptatum veniam itaque sit aspernatur veritatis.
            Ullam enim, incidunt dolores architecto illum, dicta explicabo
            minus similique debitis numquam consectetur.
          </p>
        </div>
      </div>
    </el-card>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";

@Component
export default class Posts extends Vue {}
</script>

<style scoped lang="scss">
.el-card {
  margin: 15px 0px;
}
.container {
  padding: 10px 60px 20px;
  max-width: 640px;
  &__image {
    display: inline-block;
    vertical-align: top;
    width: 46%;
    margin: 20px 30px 0 0;
  }
  &__text {
    display: inline-block;
    width: 46%;
    @media (max-width: 620px) {
      width: 100%;
    }
  }
}
</style>
