import { getAPI } from "@/axios/axios";
import { refreshToken } from "@/axios/requests/refreshRequests";
import { createHeader } from "@/axios/requests/createHeader";

function getRequest(url, params = {}, times = 0) {
  if (times < 0) {
    return;
  }
  return new Promise((resolve, reject) => {
    getAPI
      .get(url, { params: params })
      .then((response) => {
        resolve(response);
      })
      .catch((error) => {
        reject(error);
      });
  });
}

export { getRequest };
