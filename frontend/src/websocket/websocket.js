import io from "socket.io-client";
import { base_backend_domain } from "@/axios/axios";

const mode = `${process.env.MODE}`;
let url = "";

if (mode == "dev") {
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

function sendMsg(endpoint, data) {
  socket.emit(endpoint, { body: data });
}

socket.on("connect", () => {
  console.log("connected");
});

export { sendMsg, socket };
