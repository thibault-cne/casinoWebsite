<template>
  <v-app>
    <v-layout>
      <v-app-bar class="header" prominent>
        <v-app-bar-nav-icon
          variant="text"
          @click.stop="drawer = !drawer"
        ></v-app-bar-nav-icon>

        <v-toolbar-title>Casino Website</v-toolbar-title>
      </v-app-bar>

      <v-navigation-drawer
        color="#424549"
        v-model="drawer"
        bottom
        temporary
        class="navbar"
      >
        <v-list>
          <v-list-item v-for="item in items" :key="item.title" :to="item.path">
            <v-list-item class="nav-item">{{ item.title }}</v-list-item>
          </v-list-item>
        </v-list>

        <template v-slot:append>
          <div class="bottom-item">
            <v-btn block> Logout </v-btn>
          </div>
        </template>
      </v-navigation-drawer>

      <v-main class="main">
        <router-view />
      </v-main>
    </v-layout>
    <v-footer app fixed height="50px" class="footer">
      <v-col class="text-center" cols="12">
        <v-btn
          href="https://github.com/thibault-cne/casinoWebsite"
          icon
          variant="text"
        >
          <v-icon>mdi-github</v-icon>
        </v-btn>

        <strong>Fait avec ❤️ par Thibault</strong>
      </v-col>
    </v-footer>
  </v-app>
</template>

<script>
export default {
  name: "App",

  data: () => ({
    drawer: false,
    group: null,
    items: [
      {
        title: "Login",
        path: "/login",
      },
      {
        title: "Roulette",
        path: "/roulette",
      },
      {
        title: "Admin",
        path: "/admin",
      },
    ],
  }),

  watch: {
    group() {
      this.drawer = false;
    },
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
