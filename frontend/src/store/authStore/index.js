import { createStore } from "vuex";
import { actions } from "./actions";
import { mutations } from "./mutations";
import { getters } from "@/store/authStore/getters";
import VuexPersistence from "vuex-persist";

const vuexLocal = new VuexPersistence({
  storage: window.localStorage,
});

const authStore = createStore({
  state: {
    accessToken: null,
    refreshToken: null,
  },
  getters,
  mutations,
  actions,
  plugins: [vuexLocal.plugin],
});

export { authStore };
