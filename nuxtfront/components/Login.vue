<template>
  <v-container fluid fill-height class="mt-10">
    <v-layout align-center justify-center>
      <v-flex xs12 sm8 md6>
        <v-card>
          <v-card-text>
            <v-form ref="form" validation>
              <v-text-field
                prepend-icon="person"
                label="Email"
                type="email"
                @keyup.enter="onSubmit"
                v-model.trim="email"
                :rules="[v => !!v || 'required']"
              ></v-text-field>
              <v-text-field
                prepend-icon="lock"
                label="Password"
                type="password"
                @keyup.enter="onSubmit"
                v-model.trim="password"
                :rules="[v => !!v || 'required']"
              ></v-text-field>
            </v-form>
          </v-card-text>
          <v-card-actions>
            <v-btn
              color="black"
              class="white--text"
              block
              @click="onSubmit"
              :loading="loading"
            >Войти</v-btn>
          </v-card-actions>
        </v-card>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
export default {
  name: 'Login',
  data() {
    return { email: '', password: '' }
  },
  computed: {
    loading() {
      return this.$store.state.shared.loading
    },
  },
  methods: {
    async onSubmit() {
      let res = await this.$store.dispatch('login', {
        email: this.email,
        password: this.password,
      })

      if (res) {
        this.$router.push('/')
        this.$store.dispatch('fetchUserPosts', res.id)
      }
    },
  },
}
</script>

