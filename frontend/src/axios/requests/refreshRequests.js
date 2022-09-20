import { getAPI } from "@/axios/axios";
import { authStore } from "@/store/authStore";

function refreshToken() {
  const token = "Bearer ".concat(authStore.getters.refreshToken);
  const header = {
    Authorization: token,
    "Content-Type": "multipart/form-data",
  };
  return new Promise((resolve, reject) => {
    getAPI
      .get("/oauth/refresh", {
        headers: header,
      })
      .then((response) => {
        authStore.commit("setAccessToken", response.data.accessToken);
        resolve();
      })
      .catch((error) => {
        if (error.response) {
          if (error.response.status === 401) {
            authStore.commit("destroyToken");
          }
        }
        reject(error);
      });
  });
}

export { refreshToken };
