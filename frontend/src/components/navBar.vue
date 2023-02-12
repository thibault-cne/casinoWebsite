<template>
  <nav
    class="bg-gray-50 border-gray-200 px-2 sm:px-4 py-2.5 rounded dark:bg-gray-900"
  >
    <div class="container flex flex-wrap items-center justify-between mx-auto">
      <router-link to="/" class="flex items-center">
        <img
          src="https://flowbite.com/docs/images/logo.svg"
          class="h-6 mr-3 sm:h-9"
          alt="Flowbite Logo"
        />
        <span
          class="self-center text-xl font-semibold whitespace-nowrap dark:text-gray-50 text-slate-700"
          >Casino website</span
        >
      </router-link>
      <div class="flex items-center md:order-2" v-if="logged">
        <div class="flex items-center">
          <div
            class="flex flex-col p-4 md:flex-row md:mt-0 md:text-sm md:font-medium md:bg-white md:dark:bg-gray-900"
          >
            <counter :num="this.$props.userProps.wallet" />
          </div>
        </div>
        <div class="flex items-center md:order-2">
          <button
            type="button"
            class="flex mr-3 text-sm bg-gray-800 rounded-full md:mr-0 focus:ring-4 focus:ring-gray-300 dark:focus:ring-gray-600"
            id="user-menu-button"
            aria-expanded="false"
            data-dropdown-toggle="user-dropdown"
            data-dropdown-placement="bottom"
          >
            <span class="sr-only">Open user menu</span>
            <img
              class="w-8 h-8 rounded-full"
              :src="
                'http://localhost:8000/api/v1/user/get/picture/' +
                this.$props.userProps.id
              "
              alt="user photo"
            />
          </button>
          <!-- Dropdown menu -->
          <div
            class="z-50 hidden my-4 text-base list-none bg-gray-50 divide-y divide-gray-100 rounded-lg shadow dark:bg-gray-700 dark:divide-gray-600"
            id="user-dropdown"
          >
            <div class="px-4 py-3">
              <span class="block text-sm text-gray-900 dark:text-white">{{
                this.$props.userProps.username
              }}</span>
            </div>
            <ul class="py-2" aria-labelledby="user-menu-button">
              <li
                aria-labelledby="dropdownNavbarLink"
                v-if="this.$props.userProps.status !== 'user'"
              >
                <button
                  id="doubleDropdownButton"
                  data-dropdown-toggle="doubleDropdown"
                  data-dropdown-placement="right-start"
                  type="button"
                  class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 dark:hover:bg-gray-600 dark:text-gray-200 dark:hover:text-white"
                >
                  Administration
                </button>
                <div
                  id="doubleDropdown"
                  class="z-10 hidden bg-slate-50 divide-y divide-gray-100 rounded-lg shadow w-44 dark:bg-gray-700"
                >
                  <ul
                    class="py-2 text-sm text-gray-700 dark:text-gray-200"
                    aria-labelledby="doubleDropdownButton"
                  >
                    <li>
                      <router-link
                        to="/admin/panel"
                        class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 dark:hover:bg-gray-600 dark:text-gray-200 dark:hover:text-white"
                        >Dashboard administrateur</router-link
                      >
                    </li>
                    <li>
                      <router-link
                        to="/admin/user"
                        class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 dark:hover:bg-gray-600 dark:text-gray-200 dark:hover:text-white"
                        >Dashboard utilisteurs</router-link
                      >
                    </li>
                    <li>
                      <router-link
                        to="/admin/claims"
                        class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 dark:hover:bg-gray-600 dark:text-gray-200 dark:hover:text-white"
                        >Dashboard claims</router-link
                      >
                    </li>
                  </ul>
                </div>
              </li>
              <li>
                <router-link
                  to="/settings"
                  class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 dark:hover:bg-gray-600 dark:text-gray-200 dark:hover:text-white"
                  >Settings</router-link
                >
              </li>
              <li>
                <router-link
                  to="/earnings"
                  class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 dark:hover:bg-gray-600 dark:text-gray-200 dark:hover:text-white"
                  >Earnings</router-link
                >
              </li>
              <li>
                <a
                  :onclick="logout"
                  class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 dark:hover:bg-gray-600 dark:text-gray-200 dark:hover:text-white"
                  >Sign out</a
                >
              </li>
            </ul>
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
          <a @click="toggleDark">
            <svg
              viewBox="0 0 24 24"
              fill="none"
              class="w-6 h-6 cursor-pointer"
              v-if="!darkMode"
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
            <svg
              aria-hidden="true"
              id="theme-toggle-light-icon"
              class="w-5 h-5 cursor-pointer"
              fill="white"
              viewBox="0 0 20 20"
              xmlns="http://www.w3.org/2000/svg"
              v-else
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
          class="flex flex-col p-4 mt-4 bg-gray-50 md:flex-row md:space-x-8 md:mt-0 md:text-sm md:font-medium md:bg-white dark:bg-gray-800 md:dark:bg-gray-900"
        >
          <li>
            <a
              data-modal-target="authentication-modal"
              data-modal-show="authentication-modal"
              type="button"
              class="block py-2 pl-3 pr-4 text-white bg-blue-700 rounded md:bg-transparent md:text-blue-700 md:p-0 dark:text-white"
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
          class="flex flex-col p-4 mt-4 bg-gray-50 md:flex-row md:space-x-8 md:mt-0 md:text-sm md:font-medium md:bg-white dark:bg-gray-800 md:dark:bg-gray-900"
        >
          <li v-for="i in index" :key="i.name">
            <router-link
              :to="i.route"
              class="block py-2 pl-3 pr-4 text-gray-700 rounded hover:bg-gray-100 md:hover:bg-transparent md:hover:text-blue-700 md:p-0 md:dark:hover:text-white dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent dark:border-gray-700"
              >{{ i.name }}</router-link
            >
          </li>
        </ul>
      </div>
    </div>
  </nav>
</template>
<script>
import { defineComponent } from "vue";
import { initDropdowns } from "flowbite";
import { getRequest } from "@/axios/getRequest";
import counter from "@/components/counter.vue";

export default defineComponent({
  name: "navBar",
  components: {
    counter,
  },
  props: {
    loggedProps: { type: Boolean, required: true },
    userProps: { type: Object, required: false },
    darkMode: { type: Boolean, required: true },
  },
  watch: {
    loggedProps: function (newVal) {
      this.logged = newVal;
      if (this.logged) {
        initDropdowns();
      }
    },
  },
  mounted() {
    this.logged = this.loggedProps;
    initDropdowns();
  },
  data() {
    return {
      logged: false,
      dropdown: null,
      imageSrc: "",
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
<style lang=""></style>
