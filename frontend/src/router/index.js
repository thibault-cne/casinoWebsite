import { createRouter, createWebHistory } from "vue-router";
import { isLogged } from "@/axios/logged";

const routes = [
  {
    path: "/login",
    name: "login",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "@/views/LoginView.vue"),
  },
  {
    path: "/roulette",
    name: "roulette",
    beforeEnter: checkAuth,
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "@/views/RouletteView.vue"),
  },
  {
    path: "/admin",
    name: "admin",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "@/views/AdminView.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

async function checkAuth(to, from, next) {
  const status = await isLogged();
  if (!status.logged && to.name !== "login") {
    // redirect the user to the login page
    next({ name: "login" });
    return;
  }
  next();
}

router.afterEach(async (to) => {
  // check meta to put title
  if (to.meta.title) {
    document.title = to.meta.title;
  }
});

export default router;
