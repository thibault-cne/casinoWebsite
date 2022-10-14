import { getAPI } from "@/axios/axios";
import { refreshToken } from "@/axios/requests/refreshRequests";
import { createHeader } from "@/axios/requests/createHeader";

function postRequest(data, url, times = 0) {
  if (times < 0) {
    return;
  }
  return new Promise((resolve, reject) => {
    getAPI
      .post(url, data)
      .then((response) => {
        resolve(response);
      })
      .catch((error) => {
        reject(error);
      });
  });
}

export { postRequest };
