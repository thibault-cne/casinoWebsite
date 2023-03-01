import { getRequest } from "./getRequest";

async function isLogged() {
  const resp = await getRequest("/auth/connected", "");
  return resp.status === 200;
}

export { isLogged };
