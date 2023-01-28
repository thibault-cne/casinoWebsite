<template>
  <v-app class="bg-slate-700">
    <loginModal v-if="!logged" @login="(u) => login(u)" />
    <navBar :logged-props="logged" :user-props="user" @logout="logout" />
    <v-main>
      <router-view
        :user-props="user"
        @update="
          (u) => {
            update(u);
          }
        "
        @refresh="refresh"
      />
    </v-main>
  </v-app>
</template>

<script>
import loginModal from "@/components/loginModal.vue";
import navBar from "@/components/navBar.vue";
import { useDark, useToggle } from "@vueuse/core";
import { getRequest } from "./axios/getRequest";
import { initModals } from "flowbite";
import { socket } from "@/websocket/websocket";

export default {
  name: "App",
  components: { navBar, loginModal },
  created() {
    const dark = useDark();
    const toggleDark = useToggle(dark);
    toggleDark;
  },
  mounted() {
    initModals();
    this.refresh();
    socket.connect();
  },
  data() {
    return {
      logged: false,
      user: {},
    };
  },
  methods: {
    logout() {
      this.logged = false;
    },
    login(u) {
      this.user = u;
      this.logged = true;
    },
    update(u) {
      this.u = u;
    },
    refresh() {
      getRequest("/auth/connected").then((r) => {
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
