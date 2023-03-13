<template>
  <div>
    <div
      class="absolute p-10 transition-transform duration-150 ease-in-out right-0"
      :class="[alert.display ? 'translate-x-0' : 'translate-x-[100%]']"
    >
      <Alert :msg="alert.msg" :type="alert.type" />
    </div>
    <LoginModal
      v-if="!logged && $route.name !== 'register'"
      @login="(u) => login(u)"
    />
    <NavBar
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
import { useDark, useToggle } from "@vueuse/core";
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
      dark: useDark(),
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
      const dark = useDark();
      const toggle = useToggle(dark);

      toggle();
    },
    async refresh() {
      await getRequest("/auth/connected", "json").then((r) => {
        if (r.status == 200) {
          this.user = r.data.user;
          this.logged = true;
        }
      });
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

<style scoped lang="scss">
.hidden {
  right: -100%;
  opacity: 0;
}

.show {
  right: 0;
  opacity: 1;
  transition-property: tween;
  transition-timing-function: ease-out;
  transition-delay: 0.2s;
  transition-duration: 1.5s;
}
</style>
