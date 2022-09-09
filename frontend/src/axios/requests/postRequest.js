import { getAPI } from "@/axios/axios";
import { refreshToken } from "@/axios/requests/refreshRequests";
import { createHeader } from "@/axios/requests/createHeader";

function postRequest(data, url) {
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
          }
        }
        reject(error);
      });
  });
}

export { postRequest };
