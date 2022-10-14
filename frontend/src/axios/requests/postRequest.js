import { getAPI } from "@/axios/axios";

function postRequest(data, url, times = 0) {
  if (times < 0) {
    return;
  }
  const header = {
    "Content-Type": "multipart/form-data",
  };
  return new Promise((resolve, reject) => {
    getAPI
      .post(url, data, { headers: header })
      .then((response) => {
        resolve(response);
      })
      .catch((error) => {
        reject(error);
      });
  });
}

export { postRequest };
