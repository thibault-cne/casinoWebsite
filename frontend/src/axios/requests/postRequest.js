import { getAPI } from "@/axios/axios";
import { refreshToken } from "@/axios/requests/refreshRequests";
import { createHeader } from "@/axios/requests/createHeader";

function postRequest(data, url, times = 0) {
  if (times < 0) {
    return;
  }
  const header = createHeader();
  return new Promise((resolve, reject) => {
    getAPI
      .post(url, data, { headers: header })
      .then((response) => {
        resolve(response);
      })
      .catch((error) => {
        if (error.response) {
          if (error.response.status === 401) {
            refreshToken();
            postRequest(data, url, times - 1);
          }
        }
        reject(error);
      });
  });
}

export { postRequest };
