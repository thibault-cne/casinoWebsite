<template>
  <v-app class="bg-slate-700">
    <navBar :logged-props="logged" @logout="logout" />
    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>

<script>
import navBar from "@/components/navBar.vue";
import { useDark, useToggle } from "@vueuse/core";
import { isLogged } from "./axios/logged";

export default {
  name: "App",
  components: { navBar },
  created() {
    const dark = useDark();
    const toggleDark = useToggle(dark);
    toggleDark;
  },
  async mounted() {
    await this.update();
  },
  data() {
    return {
      logged: false,
    };
  },
  methods: {
    async update() {
      this.logged = await isLogged();
    },
    logout() {
      this.logged = false;
    },
  },
};
</script>
<style scoped lang="scss"></style>
