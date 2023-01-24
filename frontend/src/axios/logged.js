import { getRequest } from "./getRequest.js";

async function isLogged() {
  const resp = await getRequest("/auth/status", "");
  return resp.status ? 200 : true, false;
}

export { isLogged };
