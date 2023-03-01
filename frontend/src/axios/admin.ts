import { getRequest } from "./getRequest";

async function isAdmin() {
  const resp = await getRequest("/admin/isAdmin", "");
  return resp.status === 200;
}

export { isAdmin };
