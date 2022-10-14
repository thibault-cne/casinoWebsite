import { getAPI } from "@/axios/axios";

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
