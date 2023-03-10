import { AxiosResponse } from "axios";
import { getAPI } from "./axios";
import { createHeader } from "./createHeader";

function deleteRequest(url: string, params: {}): Promise<AxiosResponse> {
    const headers = createHeader("json");
    
    return new Promise((resolve, reject) => {
        getAPI.delete(url, {headers: headers, params: params})
        .then((r) => resolve(r))
        .catch((e) => reject(e))
    })
}

export {deleteRequest};