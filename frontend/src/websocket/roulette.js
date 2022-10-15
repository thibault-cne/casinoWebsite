import { io } from "socket.io-client";

const URL = "http://localhost:5454";
const socket = io(URL, {
  transports: ["websocket"],
  withCredentials: true,
  autoConnect: false,
  path: "/api/v1/roulette/ws",
});

export default socket;
