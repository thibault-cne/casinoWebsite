import { getAPI } from "./axios";
import { createHeader } from "./createHeader";

function postRequest(data, url, headerType) {
  const header = createHeader(headerType);

  var formatedData = data;

  if (headerType != "file") {
    formatedData = JSON.stringify(data);
  }

  return new Promise((resolve, reject) => {
    getAPI
      .post(url, formatedData, { headers: header })
      .then((response) => {
        resolve(response);
      })
      .catch((error) => {
        reject(error);
      });
  });
}

export { postRequest };
