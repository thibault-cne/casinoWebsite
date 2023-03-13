<template>
  <div class="bg-gray-50 dark:bg-gray-900">
    <div
      class="flex flex-col items-center justify-center px-6 py-8 mx-auto md:h-screen lg:py-0"
    >
      <div
        class="w-full bg-gray-50 rounded-lg shadow dark:border md:mt-0 sm:max-w-md xl:p-0 dark:bg-gray-800 dark:border-gray-700"
      >
        <div class="p-6 space-y-4 md:space-y-6 sm:p-8">
          <h1
            class="text-xl font-bold leading-tight tracking-tight text-gray-900 md:text-2xl dark:text-white"
          >
            Create and account
          </h1>
          <form class="space-y-4 md:space-y-6" action="#">
            <div>
              <label
                for="username"
                class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                >Your username</label
              >
              <input
                type="text"
                name="username"
                id="username"
                class="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                placeholder="Vladimir Petrienko"
                v-model="username"
                required
              />
              <span class="font-medium text-red-500" v-if="msg.username">{{
                msg.username
              }}</span>
            </div>
            <div>
              <label
                for="password"
                class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                >Password</label
              >
              <input
                type="password"
                name="password"
                id="password"
                placeholder="••••••••"
                v-model="password"
                class="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                required
              />
              <span class="font-medium text-red-500" v-if="msg.password">{{
                msg.password
              }}</span>
            </div>
            <div>
              <label
                for="confirm-password"
                class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                >Confirm password</label
              >
              <input
                type="password"
                name="confirm-password"
                id="confirm-password"
                placeholder="••••••••"
                v-model="passConf"
                class="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                required
              />
              <span class="font-medium text-red-500" v-if="msg.passConf">{{
                msg.passConf
              }}</span>
            </div>
            <div>
              <label
                for="code"
                class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                >Code</label
              >
              <input
                type="text"
                name="code"
                id="code"
                placeholder="12345"
                v-model="code"
                class="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                required
              />
              <span class="font-medium text-red-500" v-if="msg.code">{{
                msg.code
              }}</span>
            </div>
            <!--
            <div class="flex items-start">
              <div class="flex items-center h-5">
                <input
                  id="terms"
                  aria-describedby="terms"
                  type="checkbox"
                  class="w-4 h-4 border border-gray-300 rounded bg-gray-50 focus:ring-3 focus:ring-primary-300 dark:bg-gray-700 dark:border-gray-600 dark:focus:ring-primary-600 dark:ring-offset-gray-800"
                  required
                />
              </div>
              <div class="ml-3 text-sm">
                <label
                  for="terms"
                  class="font-light text-gray-500 dark:text-gray-300"
                  >I accept the
                  <a
                    class="font-medium text-slate-600 hover:underline dark:text-primary-500"
                    href="#"
                    >Terms and Conditions</a
                  ></label
                >
              </div>
            </div>
        -->
            <button
              type="button"
              :onclick="submit"
              class="w-full text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-primary-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800"
            >
              Create an account
            </button>
            <p class="text-sm font-light text-gray-500 dark:text-gray-400">
              Already have an account?
              <a
                href="/"
                class="font-medium text-blue-700 hover:underline dark:text-blue-500"
                >Login here</a
              >
            </p>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { postRequest } from "../axios/postRequest";

export default {
  name: "registerView",
  data() {
    return {
      username: "",
      password: "",
      passConf: "",
      code: "",
      msg: {
        username: "",
        password: "",
        passConf: "",
        code: "",
      },
    };
  },
  watch: {
    username: function (newVal: string) {
      this.username = newVal;
      this.validUsername(newVal);
    },
    password: function (newVal: string) {
      this.password = newVal;
      this.validPassword(newVal);
    },
    passConf: function (newVal: string) {
      this.passConf = newVal;
      this.validPasswordConf(newVal);
    },
  },
  methods: {
    submit() {
      if (
        !this.validUsername(this.username) ||
        !this.validPassword(this.password) ||
        !this.validPasswordConf(this.passConf)
      ) {
        return;
      }

      const data = {
        username: this.username,
        password: this.password,
        passwordVerify: this.passConf,
        code: this.code,
      };
      postRequest(data, "/auth/register", "json")
        .then((r) => {
          if (r.status === 200) {
            this.$emit(
              "_alert",
              "Your account has successfully been created.",
              "success"
            );
            this.$router.push("/");
          } else {
            this.$emit("_alert", "An error occured.", "error");
          }
        })
        .catch(() => {
          this.$emit("_alert", "An error occured.", "error");
        });
    },
    validUsername(username: string): boolean {
      if (username.length < 4 && username.length > 0) {
        this.msg.username = "Username must have a length >= 4";
        return false;
      }

      this.msg.username = "";
      return true;
    },
    validPassword(password: string): boolean {
      if (password.length < 8 && password.length > 0) {
        this.msg.password = "Password must have a length >= 8";
        return false;
      }

      this.msg.password = "";
      return true;
    },
    validPasswordConf(passConf: string): boolean {
      if (passConf !== this.password) {
        this.msg.passConf = "Password and confirmation must be identical";
        return false;
      }

      this.msg.passConf = "";
      return true;
    },
  },
};
</script>
<style lang="scss"></style>
