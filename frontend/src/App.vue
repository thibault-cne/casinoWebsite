<template>
  <div>
    <div
      class="absolute p-10 transition-transform duration-150 ease-in-out right-0 z-10"
      :class="[alert.display ? 'translate-x-0' : 'translate-x-[100%]']"
    >
      <Alert :msg="alert.msg" :type="alert.type" />
    </div>
    <LoginModal
      v-if="!logged && $route.name !== 'register'"
      @login="(u) => login(u)"
      @_alert="
            (msg: string, type: string) => {
                _alert(msg, type);
            }
        "
    />
    <NavBar
      :logged-props="logged"
      :user-props="user"
      @logout="logout"
      :dark-mode="dark"
      @toggle="toggleDark"
    />
    <div class="w-full h-screen bg-base-300">
      <router-view
        :user-props="user"
        @update="
          (u: User) => {
            update(u);
          }
        "
        @refresh="refresh"
        @_alert="
            (msg: string, type: string) => {
                _alert(msg, type);
            }
        "
      />
    </div>
  </div>
</template>

<script lang="ts">
import { getRequest } from "./axios/getRequest";
import { socket } from "./utils/websocket";
import { User } from "./models/user";
import { Alert, NavBar, LoginModal } from "./components";

export default {
  name: "App",
  components: { NavBar, LoginModal, Alert },
  async mounted() {
    await this.refresh();

    if (this.logged) {
      socket.connect();
    }
  },
  data() {
    return {
      logged: false,
      user: {} as User,
      dark: "dark",
      alert: {
        display: false,
        msg: "",
        type: "",
      },
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
      this.user = u;
    },
    toggleDark() {
      if (this.dark === "dark") {
        this.dark = "light";
      } else {
        this.dark = "dark";
      }

      document.querySelector("html")?.setAttribute("data-theme", this.dark);
    },
    async refresh() {
      await getRequest("/auth/connected", "json")
        .then((r) => {
          if (r.status == 200) {
            this.user = r.data.user;
            this.logged = true;
          }
        })
        .catch(() => {});
    },
    _alert(msg: string, type: string) {
      this.alert.msg = msg;
      this.alert.type = type;

      this.alert.display = true;
      setTimeout(() => {
        this.alert.display = false;
      }, 5000);
    },
  },
};
</script>

<style scoped lang="scss"></style>
