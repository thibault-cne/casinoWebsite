<template>
  <v-app>
    <v-layout>
      <v-app-bar color="primary" prominent>
        <v-app-bar-nav-icon
          variant="text"
          @click.stop="drawer = !drawer"
        ></v-app-bar-nav-icon>

        <v-toolbar-title>Casino Website</v-toolbar-title>
      </v-app-bar>

      <v-navigation-drawer v-model="drawer" bottom temporary>
        <v-list>
          <v-list-item v-for="item in items" :key="item.title" :to="item.path">
            <v-list-item>{{ item.title }}</v-list-item>
          </v-list-item>
        </v-list>
      </v-navigation-drawer>

      <v-main>
        <router-view />
      </v-main>
    </v-layout>
    <v-footer padless>
      <v-col class="text-center" cols="12">
        Fait avec ❤️ par <strong>Thibault</strong>
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
