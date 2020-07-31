<template>
  <v-container fluid fill-height>
    <v-layout align-center justify-center>
      <v-flex xs12 sm8 md6>
        <v-card>
          <v-card-text>
            <v-form ref="form" validation>
              <v-text-field
                prepend-icon="person"
                label="Username"
                type="username"
                @keyup.enter="onSubmit"
                v-model.trim="username"
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
            <v-btn color="black" class="white--text" block @click="onSubmit">Войти</v-btn>
          </v-card-actions>
        </v-card>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

@Component
export default class Login extends Vue {
  private username: string = "";
  private password: string = "";

  async onSubmit() {}
}
</script>

