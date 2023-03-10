<template>
  <div>
    <loginModal v-if="!logged" @login="(u) => login(u)" />
    <navBar
      :logged-props="logged"
      :user-props="user"
      @logout="logout"
      :dark-mode="dark"
      @toggle="toggleDark"
    />
    <div>
      <router-view
        :user-props="user"
        @update="
          (u: User) => {
            update(u);
          }
        "
        @refresh="refresh"
      />
    </div>
  </div>
</template>

<script lang="ts">
import loginModal from "./components/loginModal.vue";
import navBar from "./components/navBar.vue";
import { useDark, useToggle } from "@vueuse/core";
import { getRequest } from "./axios/getRequest";
import { initModals } from "flowbite";
import { socket } from "./utils/websocket";
import { User } from "./models/user";

export default {
  name: "App",
  components: { navBar, loginModal },
  mounted() {
    initModals();
    this.refresh();

    if (this.logged) {
      socket.connect();
    }
  },
  data() {
    return {
      logged: false,
      user: {} as User,
      dark: useDark(),
    };
  },
  methods: {
    logout() {
      this.logged = false;
    },
    login(u: User) {
      this.user = u;
      this.logged = true;
      socket.connect();
    },
    update(u: User) {
      console.log("Update ", u);
      this.user = u;
      console.log(this.user);
    },
    toggleDark() {
      const dark = useDark();
      const toggle = useToggle(dark);

      toggle();
    },
    refresh() {
      getRequest("/auth/connected", "json").then((r) => {
        if (r.status == 200) {
          this.user = r.data.user;
          this.logged = true;
        }
      });
    },
  },
};
</script>

<style scoped lang="scss"></style>
