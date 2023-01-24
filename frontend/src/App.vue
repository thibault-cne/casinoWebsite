<template>
  <v-app class="bg-slate-700">
    <navBar />
    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>

<script>
import navBar from "@/components/navBar.vue";
import { useDark, useToggle } from "@vueuse/core";

export default {
  name: "App",
  components: { navBar },
  created() {
    const dark = useDark();
    console.log(dark);
    const toggleDark = useToggle(dark);
    toggleDark;
  },
};
</script>
<style scoped lang="scss">
@use "@/assets/styles/scss/standards/colors";

.header {
  background-color: colors.$bg-header;
  color: colors.$header;
}

.bottom-item {
  margin-bottom: 60px;
}

.nav-item {
  color: colors.$text-primary !important;
}

.main {
  padding-bottom: 50px;
}

.footer {
  z-index: 9;
  background-color: colors.$bg-footer;
  color: colors.$footer;
}
</style>
