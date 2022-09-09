import { authStore } from "@/store/authStore";

function createHeader() {
  let header;
  let contentType = "multipart/form-data";
  if (authStore.getters.loggedIn) {
    const token = "Bearer " + authStore.getters.accessToken;
    header = {
      Authorization: token,
      "Content-Type": contentType,
    };
  } else {
    header = {
      "Content-Type": contentType,
    };
  }

  return header;
}

export { createHeader };
