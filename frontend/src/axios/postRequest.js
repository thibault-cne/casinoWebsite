import { getAPI } from "./axios";
import { createHeader } from "./createHeader";

function postRequest(data, url, headerType) {
  const header = createHeader(headerType);
  return new Promise((resolve, reject) => {
    getAPI
      .post(url, JSON.stringify(data), { headers: header })
      .then((response) => {
        resolve(response);
      })
      .catch((error) => {
        reject(error);
      });
  });
}

export { postRequest };