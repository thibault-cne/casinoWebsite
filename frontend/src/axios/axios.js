import axios from "axios";

const base_backend_domain = `${process.env.VUE_APP_BACKEND_DOMAIN}`;
const base_backend_url = `${process.env.VUE_APP_BACKEND_URL}` + "/api/v1";

axios.defaults.withCredentials = true;
const getAPI = axios.create({
  baseURL: base_backend_url,
  headers: {
    "Content-Type": "application/json",
  },
});

export { getAPI, base_backend_domain, base_backend_url };
