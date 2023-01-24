import { getAPI } from "./axios";
import { createHeader } from "./createHeader";

function getRequest(url, headerType, params = {}) {
  const header = createHeader(headerType);
  return new Promise((resolve, reject) => {
    getAPI
      .get(url, { headers: header, params: params })
      .then((response) => {
        resolve(response);
      })
      .catch((error) => {
        if (error.response) {
          if (error.response.status === 401) {
            // refreshToken();
          }
        }
        reject(error);
      });
  });
}

export { getRequest };
