import io from "socket.io-client";
import { base_backend_domain } from "../axios/axios";

let url = "";

if (import.meta.env.DEV) {
  url = "ws://" + base_backend_domain;
} else {
  url = "wss://" + base_backend_domain;
}

const socket = io(url, {
  transports: ["websocket"],
  withCredentials: true,
  autoConnect: false,
  path: "/api/v1/ws/",
});

function sendMsg(endpoint: string, data: any) {
  socket.emit(endpoint, { body: data });
}

socket.on("connect", () => {
});

export { sendMsg, socket };
