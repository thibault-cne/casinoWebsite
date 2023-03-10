import { AxiosResponse } from "axios";
import { getAPI } from "./axios";
import { createHeader } from "./createHeader";

function getRequest(url: string, headerType: string, params = {}): Promise<AxiosResponse> {
  const header = createHeader(headerType);
  return new Promise((resolve, reject) => {
    getAPI
      .get(url, { headers: header, params: params })
      .then((response) => {
        resolve(response);
      })
      .catch((error) => {
        reject(error);
      });
  });
}

export { getRequest };
