import axios from "axios";

const backendDomain = `${process.env.VUE_APP_BACKEND_DOMAIN}:${process.env.VUE_APP_BACKEND_PORT}`;
const base_backend_url = "http://" + backendDomain + "/api/v1";

axios.defaults.withCredentials = true;
const getAPI = axios.create({
  baseURL: base_backend_url,
  headers: {
    "Content-Type": "application/json",
  },
});

export { getAPI, base_backend_url, backendDomain };
