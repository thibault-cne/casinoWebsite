import { createRouter, createWebHistory } from "vue-router";
import { isLogged } from "../axios/logged";
import { isAdmin } from "../axios/admin";

const routes = [
  {
    path: "/register",
    name: "register",
    component: () =>
      import(/* webpackChunkName: "about" */ "../vues/registerVue.vue"),
  },
  {
    path: "/roulette",
    name: "roulette",
    beforeEnter: checkAuth,
    component: () =>
      import(/* webpackChunkName: "about" */ "../vues/rouletteHome.vue"),
  },
  {
    path: "/blackjack",
    name: "Blackjack",
    beforeEnter: checkAuth,
    component: () =>
      import(/* webpackChunkName: "about" */ "../vues/blackJack.vue"),
  },
  {
    path: "/earnings",
    name: "earnings",
    beforeEnter: checkAuth,
    component: () =>
      import(/* webpackChunkName: "about" */ "../vues/userEarnings.vue"),
  },
  {
    path: "/settings",
    name: "settings",
    beforeEnter: checkAuth,
    component: () =>
      import(/* webpackChunkName: "about" */ "../vues/userSettings.vue"),
  },
  {
    path: "/admin/user",
    name: "admin user panel",
    beforeEnter: checkAdmin,
    component: () =>
      import(/* webpackChunkName: "about" */ "../vues/adminUserPanel.vue"),
  },
  {
    path: "/admin/claims",
    name: "admin claims panel",
    beforeEnter: checkAdmin,
    component: () =>
      import(/* webpackChunkName: "about" */ "../vues/adminClaimPanel.vue"),
  },
  {
    path: "/admin/panel",
    name: "admin panel",
    beforeEnter: checkAdmin,
    component: () =>
      import(/* webpackChunkName: "about" */ "../vues/adminPanel.vue"),
  },
  {
    path: "/",
    name: "home",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "../vues/homeVue.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

async function checkAuth(to: any, from: any, next: any) {
  const status = await isLogged();
  if (!status && to.name !== "login") {
    // redirect the user to the login page
    next({ name: "login" });
    return;
  }
  next();
}

async function checkAdmin(to: any, from: any, next: any) {
  const status = await isAdmin();
  if (!status) {
    // redirect the user to the from page
    from();
    return;
  }
  next();
}

router.afterEach(async (to) => {
  // check meta to put title
  if (to.meta.title) {
    document.title = to.meta.title as string;
  }
});

export default router;
