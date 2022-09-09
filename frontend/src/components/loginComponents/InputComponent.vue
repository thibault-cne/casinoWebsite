<template>
  <v-sheet class="bg-deep-purple pa-12" height="90vh" width="100vw">
    <v-card class="mx-auto px-6 py-8" max-width="344">
      <v-card-title class="pa-4">Login</v-card-title>
      <v-form ref="form" v-model="valid" lazy-validation>
        <v-text-field
          v-model="username"
          :readonly="loading"
          :rules="[required]"
          class="mb-2"
          clearable
          label="Username"
        ></v-text-field>

        <v-text-field
          v-model="password"
          :readonly="loading"
          :rules="[required]"
          clearable
          label="Password"
          placeholder="Enter your password"
          :append-icon="show ? 'mdi-eye' : 'mdi-eye-off'"
          :type="show ? 'text' : 'password'"
          @click:append="show = !show"
        ></v-text-field>

        <br />

        <v-card-actions>
          <v-spacer></v-spacer>

          <v-btn color="success" :disabled="!valid" @click="validate">
            Complete Registration

            <v-icon icon="mdi-chevron-right" end></v-icon>
          </v-btn>
        </v-card-actions>
      </v-form>
    </v-card>
  </v-sheet>
</template>
<script>
import { postRequest } from "@/axios/requests/postRequest";
import { authStore } from "@/store/authStore";
export default {
  data() {
    return {
      username: "",
      password: "",
      show: false,
      loading: false,
      valid: false,
    };
  },
  methods: {
    submit() {
      let data = { username: this.username, password: this.password };

      console.log(data);

      postRequest(data, "/oauth/login").then((response) => {
        let credentials = {
          accessToken: response.data.access_token,
          refreshToken: response.data.refresh_token,
        };

        authStore.commit("updateStorage", credentials);
        console.log(authStore.getters.loggedIn);
      });
    },
    required(v) {
      return !!v || "Field is required";
    },
    validate() {
      this.$refs.form.validate();

      this.loading = true;
      this.submit();
      setTimeout(() => (this.loading = false), 2000);
    },
  },
};
</script>
<style lang="css"></style>
