<template>
  <!-- Main modal -->
  <div
    id="authentication-modal"
    tabindex="-1"
    aria-hidden="true"
    class="fixed top-0 left-0 right-0 z-50 hidden w-full p-4 overflow-x-hidden overflow-y-auto md:inset-0 h-modal md:h-full"
  >
    <div class="relative w-full h-full max-w-md md:h-auto">
      <!-- Modal content -->
      <div class="relative rounded-lg shadow bg-base-100">
        <button
          type="button"
          class="absolute top-3 right-2.5 bg-transparent hover:bg-base-300 rounded-lg text-sm p-1.5 ml-auto inline-flex items-center"
          data-modal-hide="authentication-modal"
        >
          <svg
            aria-hidden="true"
            class="w-5 h-5"
            fill="currentColor"
            viewBox="0 0 20 20"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              fill-rule="evenodd"
              d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z"
              clip-rule="evenodd"
            ></path>
          </svg>
          <span class="sr-only">Close modal</span>
        </button>
        <div class="px-6 py-6 lg:px-8">
          <h3 class="mb-4 text-3xl font-bold text-primary">
            Sign in to our platform
          </h3>
          <form class="space-y-6" action="#">
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
                class="input input-bordered input-primary w-full max-w-xs placeholder-base-content"
                placeholder="Vladimir Petrienko"
                required
                v-model="username"
              />
            </div>
            <div>
              <label
                for="password"
                class="block mb-2 text-xl font-semibold text-primary"
                >Your password</label
              >
              <input
                type="password"
                name="password"
                id="password"
                placeholder="••••••••"
                class="input input-bordered input-primary w-full max-w-xs placeholder-base-content"
                required
                v-model="password"
              />
            </div>
            <div class="flex justify-center">
              <button
                :onclick="submit"
                type="button"
                class="btn btn-wide btn-primary"
                data-modal-hide="authentication-modal"
              >
                Login to your account
              </button>
            </div>
            <div class="text-sm font-medium text-gray-500 dark:text-gray-300">
              Not registered?
              <a
                href="/register"
                class="text-blue-700 hover:underline dark:text-blue-500 font-medium"
                >Create account</a
              >
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { initModals } from "flowbite";
import { postRequest } from "../axios/postRequest";

export default defineComponent({
  name: "loginModal",
  mounted() {
    initModals();
  },
  data() {
    return {
      username: "",
      password: "",
    };
  },
  methods: {
    submit() {
      const data = {
        username: this.username,
        password: this.password,
      };
      postRequest(data, "/auth/login", "json")
        .then((r) => {
          if (r.status == 200) {
            this.$emit("login", r.data.user);
            this.$emit("_alert", "You successfully logged in.", "success");
          } else {
            this.$emit("_alert", "An error occured during the login.", "error");
          }
        })
        .catch(() => {
          this.$emit("_alert", "An error occured during the login.", "error");
        });
    },
  },
});
</script>

<style lang="scss"></style>
