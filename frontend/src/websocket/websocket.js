import io from "socket.io-client";
import { base_backend_domain } from "@/axios/axios";

const socket = io("ws://" + base_backend_domain, {
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
