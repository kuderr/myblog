<template>
  <div id="about">
    <v-layout align-center justify-center>
      <v-flex xs12 sm10 md8>
        <v-card class="mx-auto" elevation="3">
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

              <v-btn
                class="ma-1"
                text
                icon
                color="red"
                href="mailto:dakudryavcev@gmail.com"
              >
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
                <b>Программирование, психология, инвестиции и дизайн </b>
              </p>

              <v-chip
                v-for="skill in skills"
                :key="skill"
                class="ma-1"
                color="indigo darken-2"
                text-color="white"
              >
                {{ skill }}
              </v-chip>

              <div class="mt-5">
                <p class="mb-3">
                  <b>Программирование</b><br />
                  Пишу на Python, Go и JS. В универе разрабатываю проекты.
                  <br />Построил
                  <b>
                    <a
                      href="https://cabinet.miem.hse.ru/#/project/19/"
                      target="_blank"
                      >NVR</a
                    > </b
                  >, в котором я лидер и главный разработчик. Сервис публикует в
                  LMS пары по расписанию из оборудованных аудиторий, а также все
                  пары проходящие онлайн. Студенты могут пересматривать материал
                  занятий несколько раз. А те, кто из-за пандемии не могут
                  посещать занятия очно, продолжают обучаться. За 2 года сервис
                  стал одним из ключевых, а я — руководителем.
                </p>
                <p class="mb-3">
                  Параллельно работаю над автогрейдерами. Проект даст
                  возможность преподавателям ставить на автомат свои лабы. Всё,
                  что нужно сделать — задать входные и выходные данные,
                  технологию и куда ставить оценку. Но это не контесты. Недавно
                  на наших грейдерах прогнали лабу для своего же потока. Задача
                  была такая: найти самую яркую точку в поле зрения камеры и
                  переместиться на неё.
                </p>
                <p>
                  Хочу развиваться дальше в создании ПО, больше в сторону
                  архитектора, чем менеджера. А также, передавать людям свой
                  опыт. Через 5 лет вижу себя экспертом Python, системным
                  архитектором, ментором, консультантом и лидером. Мечтаю
                  создавать красивые и удобные вещи, пользоваться которыми —
                  удовольствие.
                </p>
              </div>

              <p class="mt-4">
                <b>Психология</b><br />
                Начал интересоваться психологией, чтобы разобраться в себе.
                Разобрался более-менее, но втянулся в тему. Круто понимать о чём
                думают люди.
              </p>

              <p class="mt-4">
                <b>Инвестиции</b><br />
                Год сижу на фондовом рынке. Изучил много литературы по теме. Без
                лишних нервов пережил пандемию 2020 и вышел в +30% при немалом
                портфеле. Подробнее в постах или в Тинькофф Пульс.
              </p>

              <p class="mt-4">
                <b>Дизайн</b><br />
                Я всегда любил красивые вещи. Для меня дизайн — всё приятное
                глазу и мозгу: цвет, форма, композиция, информативный текст,
                быстрый фронт и тд. Но мне недостаточно наблюдать, хочу
                создавать сам.
              </p>

              <p class="mt-4">
                <b>Моя цель</b> — сделать мир лучше. Я этого добьюсь, если буду
                создавать новое и обучать. Для этого и создал блог. Здесь
                самовыражаюсь, делюсь опытом и практикуюсь в чётком изложении
                своих мыслей.
              </p>

              <p class="mt-4">
                Первый компьютер в нашей семье появился, когда мне было 4 года.
                Типичный офисный набор нулевых: белый строгий корпус и
                толстенный монитор. Годы шли, компьютеры менялись. А страсть к
                технологиям только росла.
              </p>
            </div>
          </v-card-text>
        </v-card>
      </v-flex>
    </v-layout>
  </div>
</template>
<script>
export default {
  name: 'About',
  data() {
    return {
      skills: [
        'Python',
        'Flask',
        'SQLAlchemy',
        'Tortoise ORM',
        'asyncpg',
        'asyncio',
        'aiohttp',
        'aiofiles',
        'FastAPI',
        'WebSocket',
        'HTTP',
        'JS',
        'Vuejs',
        'Go',
        'PostgreSQL',
        'Nginx',
        'Docker',
        'Linux',
        'Git',
        'Google APIs',
      ],
    }
  },
  head() {
    return {
      title: 'Об авторе',
      meta: [
        { name: 'description', content: 'Блог Димы Кудрявцева' },
        { name: 'author', content: 'Dmitriy Kudryavtsev' },
        { name: 'keywords', content: 'blog, author, блог, автор' },

        {
          name: 'og:url',
          content: 'https://kuderblog.com/about',
        },
        { name: 'og:title', content: 'Обо мне' },
        { name: 'og:type', content: 'profile' },
        { name: 'og:site_name', content: 'Kuder Blog' },
        { name: 'og:locale', content: 'ru_RU' },
        { name: 'og:description', content: 'Блог Димы Кудрявцева' },
        {
          name: 'og:image',
          content:
            'https://avatars0.githubusercontent.com/u/39552217?s=460&u=415aebf4249492578b8b2af663f5e5473c704dd4&v=4',
        },

        { name: 'twitter:card', content: 'summary' },
        {
          name: 'twitter:image',
          content:
            'https://avatars0.githubusercontent.com/u/39552217?s=460&u=415aebf4249492578b8b2af663f5e5473c704dd4&v=4',
        },
        {
          name: 'twitter:url',
          content: 'https://kuderblog.com/about',
        },
        { name: 'twitter:title', content: 'Обо мне' },
        { name: 'twitter:description', content: 'Блог Димы Кудрявцева' },
        { name: 'twitter:creator', content: '@kuderrr' },
        { name: 'twitter:site', content: '@kuderrr' },
      ],
    }
  },
}
</script>

<style lang="scss" scoped>
a {
  text-decoration: none;
}
</style>
