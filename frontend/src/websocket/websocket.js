import io from "socket.io-client";
import { backendDomain } from "@/axios/axios";

const socket = io("ws://" + backendDomain, {
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
