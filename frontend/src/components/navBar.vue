<template>
  <nav class="bg-base-100 border-gray-200 px-2 sm:px-4 py-2.5 rounded">
    <div class="container flex flex-wrap items-center justify-between mx-auto">
      <router-link to="/" class="flex items-center">
        <span
          class="self-center text-xl whitespace-nowrap text-primary font-bold"
          >Casino website</span
        >
      </router-link>
      <div class="flex items-center md:order-2" v-if="logged">
        <div class="flex items-center">
          <div
            class="flex flex-col p-4 md:flex-row md:mt-0 md:text-sm md:font-medium"
          >
            <Counter :num="$props.userProps.wallet" />
          </div>
        </div>
        <div class="flex items-center md:order-2">
          <!-- Dropdown menu -->
          <div class="dropdown dropdown-end dropdown-hover">
            <label tabindex="0">
              <span class="sr-only">Open user menu</span>
              <img
                class="w-8 h-8 rounded-full"
                :src="
                  base_backend_url +
                  '/user/get/picture?userId=' +
                  $props.userProps.id
                "
                alt="user photo"
              />
            </label>
            <div
              class="dropdown-content menu p-2 shadow rounded-box w-52 bg-primary"
            >
              <div v-if="userProps.status !== 'user'">
                <ul>
                  <li>
                    <router-link
                      to="/admin/panel"
                      class="block px-4 py-2 text-sm text-primary-content hover:bg-primary-focus"
                      >Dashboard administrateur</router-link
                    >
                  </li>
                  <li>
                    <router-link
                      to="/admin/user"
                      class="block px-4 py-2 text-sm text-primary-content hover:bg-primary-focus"
                      >Dashboard utilisateurs</router-link
                    >
                  </li>
                  <li>
                    <router-link
                      to="/admin/claims"
                      class="block px-4 py-2 text-sm text-primary-content hover:bg-primary-focus"
                      >Dashboard claims</router-link
                    >
                  </li>
                </ul>
              </div>
              <div class="divider" v-if="userProps.status !== 'user'"></div>
              <div>
                <ul>
                  <li>
                    <router-link
                      to="/settings"
                      class="block px-4 py-2 text-sm text-primary-content hover:bg-primary-focus"
                      >Settings</router-link
                    >
                  </li>
                  <li>
                    <router-link
                      to="/earnings"
                      class="block px-4 py-2 text-sm text-primary-content hover:bg-primary-focus"
                      >Earnings</router-link
                    >
                  </li>
                </ul>
              </div>
              <div class="divider"></div>
              <div>
                <ul>
                  <li>
                    <a
                      :onclick="logout"
                      class="block px-4 py-2 text-sm text-primary-content hover:bg-primary-focus"
                      >Sign out</a
                    >
                  </li>
                </ul>
              </div>
            </div>
          </div>
          <button
            data-collapse-toggle="mobile-menu-2"
            type="button"
            class="inline-flex items-center p-2 ml-1 text-sm text-gray-500 rounded-lg md:hidden hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200 dark:text-gray-400 dark:hover:bg-gray-700 dark:focus:ring-gray-600"
            aria-controls="mobile-menu-2"
            aria-expanded="false"
          >
            <span class="sr-only">Open main menu</span>
            <svg
              class="w-6 h-6"
              aria-hidden="true"
              fill="currentColor"
              viewBox="0 0 20 20"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                fill-rule="evenodd"
                d="M3 5a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 10a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 15a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1z"
                clip-rule="evenodd"
              ></path>
            </svg>
          </button>
          <a @click="toggleDark" class="swap swap-rotate p-4">
            <!-- this hidden checkbox controls the state -->
            <input type="checkbox" v-model="dark" />

            <!-- sun icon -->
            <svg
              viewBox="0 0 24 24"
              fill="none"
              class="cursor-pointer swap-off w-6 h-6"
            >
              <path
                fill-rule="evenodd"
                clip-rule="evenodd"
                d="M17.715 15.15A6.5 6.5 0 0 1 9 6.035C6.106 6.922 4 9.645 4 12.867c0 3.94 3.153 7.136 7.042 7.136 3.101 0 5.734-2.032 6.673-4.853Z"
                class="fill-transparent"
              ></path>
              <path
                d="m17.715 15.15.95.316a1 1 0 0 0-1.445-1.185l.495.869ZM9 6.035l.846.534a1 1 0 0 0-1.14-1.49L9 6.035Zm8.221 8.246a5.47 5.47 0 0 1-2.72.718v2a7.47 7.47 0 0 0 3.71-.98l-.99-1.738Zm-2.72.718A5.5 5.5 0 0 1 9 9.5H7a7.5 7.5 0 0 0 7.5 7.5v-2ZM9 9.5c0-1.079.31-2.082.845-2.93L8.153 5.5A7.47 7.47 0 0 0 7 9.5h2Zm-4 3.368C5 10.089 6.815 7.75 9.292 6.99L8.706 5.08C5.397 6.094 3 9.201 3 12.867h2Zm6.042 6.136C7.718 19.003 5 16.268 5 12.867H3c0 4.48 3.588 8.136 8.042 8.136v-2Zm5.725-4.17c-.81 2.433-3.074 4.17-5.725 4.17v2c3.552 0 6.553-2.327 7.622-5.537l-1.897-.632Z"
                class="fill-slate-400 dark:fill-slate-500"
              ></path>
              <path
                fill-rule="evenodd"
                clip-rule="evenodd"
                d="M17 3a1 1 0 0 1 1 1 2 2 0 0 0 2 2 1 1 0 1 1 0 2 2 2 0 0 0-2 2 1 1 0 1 1-2 0 2 2 0 0 0-2-2 1 1 0 1 1 0-2 2 2 0 0 0 2-2 1 1 0 0 1 1-1Z"
                class="fill-slate-400 dark:fill-slate-500"
              ></path>
            </svg>

            <!-- moon icon -->
            <svg
              aria-hidden="true"
              id="theme-toggle-light-icon"
              class="w-6 h-6 cursor-pointer swap-on"
              fill="white"
              viewBox="0 0 20 20"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                d="M10 2a1 1 0 011 1v1a1 1 0 11-2 0V3a1 1 0 011-1zm4 8a4 4 0 11-8 0 4 4 0 018 0zm-.464 4.95l.707.707a1 1 0 001.414-1.414l-.707-.707a1 1 0 00-1.414 1.414zm2.12-10.607a1 1 0 010 1.414l-.706.707a1 1 0 11-1.414-1.414l.707-.707a1 1 0 011.414 0zM17 11a1 1 0 100-2h-1a1 1 0 100 2h1zm-7 4a1 1 0 011 1v1a1 1 0 11-2 0v-1a1 1 0 011-1zM5.05 6.464A1 1 0 106.465 5.05l-.708-.707a1 1 0 00-1.414 1.414l.707.707zm1.414 8.486l-.707.707a1 1 0 01-1.414-1.414l.707-.707a1 1 0 011.414 1.414zM4 11a1 1 0 100-2H3a1 1 0 000 2h1z"
                fill-rule="evenodd"
                clip-rule="evenodd"
              ></path>
            </svg>
          </a>
        </div>
      </div>
      <div v-else class="flex items-center md:order-2">
        <ul
          class="flex flex-col p-4 mt-4 md:flex-row md:space-x-8 md:mt-0 md:text-sm md:font-medium"
        >
          <li>
            <a
              data-modal-target="authentication-modal"
              data-modal-show="authentication-modal"
              type="button"
              class="block py-2 pl-3 pr-4 text-primary hover:text-primary-focus rounded md:bg-transparent md:p-0 cursor-pointer"
              >Login</a
            >
          </li>
        </ul>
      </div>
      <div
        class="items-center justify-between hidden w-full md:flex md:w-auto md:order-1"
        id="mobile-menu-2"
      >
        <ul
          class="flex flex-col p-4 mt-4 md:flex-row md:space-x-8 md:mt-0 md:text-sm md:font-medium"
        >
          <li v-for="i in index" :key="i.name">
            <router-link
              :to="i.route"
              class="text-primary hover:text-primary-focus block py-2 pl-3 pr-4 rounded md:hover:bg-transparent md:p-0"
              >{{ i.name }}</router-link
            >
          </li>
        </ul>
      </div>
    </div>
  </nav>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { getRequest } from "../axios/getRequest";
import { Counter } from ".";
import { base_backend_url } from "../axios/axios";
import { User } from "../models/user";

export default defineComponent({
  name: "navBar",
  components: {
    Counter,
  },
  props: {
    loggedProps: { type: Boolean, required: true },
    userProps: { type: Object as () => User, required: true },
    darkMode: { type: String, required: true },
  },
  watch: {
    loggedProps: function (newVal) {
      this.logged = newVal;
    },
    darkMode: function () {
      this.dark = !this.dark;
    },
  },
  mounted() {
    this.logged = this.loggedProps;
  },
  data() {
    return {
      base_backend_url: base_backend_url,
      logged: false,
      dropdown: null,
      dark: true,
      index: [
        {
          name: "Roulette",
          route: "/roulette",
        },
      ],
    };
  },
  methods: {
    logout() {
      getRequest("/auth/logout", "");
      this.$emit("logout");
    },
    toggleDark() {
      this.$emit("toggle");
    },
  },
});
</script>

<style lang="scss"></style>
