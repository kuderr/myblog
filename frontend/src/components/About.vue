<template>
  <div id="about">
    <v-layout align-center justify-center>
      <v-flex xs12 sm10 md9>
        <v-card class="mx-auto">
          <div class="text-center py-3">
            <v-avatar center class="profile" color="grey" size="200">
              <v-img
                src="https://avatars0.githubusercontent.com/u/39552217?s=460&u=415aebf4249492578b8b2af663f5e5473c704dd4&v=4"
              ></v-img>
            </v-avatar>

            <div>
              <v-btn
                class="ma-1"
                text
                icon
                color="black"
                href="https://github.com/kuderr"
                target="_blank"
              >
                <v-icon>mdi-github</v-icon>
              </v-btn>

              <v-btn
                class="ma-1"
                text
                icon
                color="blue lighten-2"
                href="https://twitter.com/kuderrr"
                target="_blank"
              >
                <v-icon>mdi-twitter</v-icon>
              </v-btn>

              <v-btn
                class="ma-1"
                text
                icon
                color="blue darken-2"
                href="https://vk.com/kuderrr"
                target="_blank"
              >
                <v-icon>mdi-vk</v-icon>
              </v-btn>

              <v-btn
                class="ma-1"
                text
                icon
                color="pink"
                href="https://www.instagram.com/kuderr_/"
                target="_blank"
              >
                <v-icon>mdi-instagram</v-icon>
              </v-btn>

              <v-btn class="ma-1" text icon color="red" href="mailto:dakudryavcev@gmail.com">
                <v-icon>mdi-gmail</v-icon>
              </v-btn>
            </div>
          </div>
          <v-card-text class="text--primary">
            <div class="body-1">
              <p>
                Привет, я
                <b>Дима Кудрявцев</b>
                <br />
                <b>Программист, жизнелюб</b>
              </p>

              <p>
                <b>Программист.</b> Последний год веду проект в вузе, также являюсь в нём основным разработчиком Full-stack. Пишу на Python, Go и JS/TS. С питоном история любви простая - увидел, познакомился получше и влюбился (принимаю со всеми недостатками и все дела). С js всё сложнее, нужно было срочно делать фронт ...
                <br />Построил
                <b>
                  <a href="https://cabinet.miem.hse.ru/#/project/19/" target="_blank">NVR</a>
                </b>, за который не стыдно, используя Flask на серверной части и Vue.js на фронте (лучший фреймворк кста, не собираюсь слезать). Работаю над проектом автоматизации проверки студенческих работ (не просто лаб по прогам). Тут уже бек делаю на aiohttp, выжимаю максимально возможную мощь с питона.
                Тем не менее, я захотел что-то помощнее. Поэтому этот блог написан на Go, фронт на Vue+TS. Изучение этих технологий и построение приложения с учётом своего уже немалого опыта - это огромная мотивация для меня.
                Много работал в команде и помогал ребятам с факультета по вопросам программирования, поэтому это одно из направлений, о котором буду писать
              </p>
              <p>
                <b>Жизнелюб.</b> Вряд ли встретишь это слово часто на просторах рунета, мне его ещё и гуглить пришлось.
                <br />Одновременно с профессиональными навыками прокачивал себя как личность. Я думаю, у многих людей есть свои страхи, комплексы и тп. Вопрос в том, признают они их и готовы ли от них избавиться? И, видимо, у меня получается справляться с тараканами в своей голове. На психологический темы я буду писать отдельные посты
              </p>
              <p>
                Помимо этого увлекаюсь:
                <b>историей</b>,
                <b>экономикой</b> и
                <b>дизайном</b>. Думаю, этот список ещё будет пополняться
              </p>
              <p>Никогда не любил систему образования. Буду писать и об этом. Ведь я тому пример, как саботировать против стандарта и выбирать свой путь. Ведь никто в школе и универе вас нормально обучить не способен. Нужно лишь понять свои цели и желания и развиваться для их достижения, а учёбу оставить как формальность</p>
              <p>Вообщем, блог задуман для того, чтобы мне практиковаться в чётком изложении своих мыслей, а вам принимать мои идеи, либо отвергать. Самовыражение важно для каждого, своё же я нахожу в этом блоге и общении с вами</p>
              <p>Всем счастья ♥️</p>
            </div>
          </v-card-text>
        </v-card>
      </v-flex>
    </v-layout>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

@Component
export default class About extends Vue {}
</script>

<style lang="scss" scoped>
a {
  text-decoration: none;
}
</style>