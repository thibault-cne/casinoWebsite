<template>
  <div class="bg-base-300">
    <div
      class="flex flex-col items-center justify-center px-6 py-8 mx-auto md:h-screen lg:py-0"
    >
      <div
        class="bg-base-100 w-full rounded-lg shadow md:mt-0 sm:max-w-md xl:p-0"
      >
        <div class="p-6 space-y-4 md:space-y-6 sm:p-8">
          <h1
            class="text-3xl font-bold leading-tight tracking-tight text-primary"
          >
            Create and account
          </h1>
          <form class="space-y-4 md:space-y-6" action="#">
            <div>
              <label
                for="username"
                class="block mb-2 text-xl font-semibold text-primary"
                >Your username</label
              >
              <input
                type="text"
                name="username"
                id="username"
                class="input input-bordered input-primary w-full max-w-xs placeholder-primary-content"
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
                class="block mb-2 text-xl font-semibold text-primary"
                >Password</label
              >
              <input
                type="password"
                name="password"
                id="password"
                placeholder="••••••••"
                v-model="password"
                class="input input-bordered input-primary w-full max-w-xs placeholder-primary-content"
                required
              />
              <span class="font-medium text-red-500" v-if="msg.password">{{
                msg.password
              }}</span>
            </div>
            <div>
              <label
                for="confirm-password"
                class="block mb-2 text-xl font-semibold text-primary"
                >Confirm password</label
              >
              <input
                type="password"
                name="confirm-password"
                id="confirm-password"
                placeholder="••••••••"
                v-model="passConf"
                class="input input-bordered input-primary w-full max-w-xs placeholder-primary-content"
                required
              />
              <span class="font-medium text-red-500" v-if="msg.passConf">{{
                msg.passConf
              }}</span>
            </div>
            <div>
              <label
                for="code"
                class="block mb-2 text-xl font-semibold text-primary"
                >Code</label
              >
              <input
                type="text"
                name="code"
                id="code"
                placeholder="12345"
                v-model="code"
                class="input input-bordered input-primary w-full max-w-xs placeholder-primary-content"
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
            <div class="flex justify-center">
              <button
                :onclick="submit"
                type="button"
                class="btn btn-wide btn-primary"
                data-modal-hide="authentication-modal"
              >
                Register
              </button>
            </div>
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
