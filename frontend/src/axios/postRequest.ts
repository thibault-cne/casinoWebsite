import { AxiosResponse } from "axios";
import { getAPI } from "./axios";
import { createHeader } from "./createHeader";

function postRequest(data: any, url: string, headerType: string): Promise<AxiosResponse> {
  const header = createHeader(headerType);

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
