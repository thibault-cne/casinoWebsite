import { getAPI } from "@/axios/axios";
import { refreshToken } from "@/axios/requests/refreshRequests";
import { createHeader } from "@/axios/requests/createHeader";

function getRequest(url, params = {}, times = 0) {
  if (times < 0) {
    return;
  }
  let header = createHeader();
  return new Promise((resolve, reject) => {
    getAPI
      .get(url, { headers: header, params: params })
      .then((response) => {
        resolve(response);
      })
      .catch((error) => {
        if (error.response) {
          if (error.response.status === 401) {
            refreshToken();
            getRequest(url, params, times - 1);
          }
        }
        reject(error);
      });
  });
}

export { getRequest };
