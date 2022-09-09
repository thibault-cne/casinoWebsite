import { refreshToken } from "@/axios/requests/refreshRequests";
import { authStore } from "@/store/authStore";

let uri = "";

if (authStore.getters.loggedIn) {
  uri =
    "ws://localhost:5454/api/v1/roulette/connect?accessToken=" +
    authStore.getters.accessToken;
} else {
  uri = "ws://localhost:5454/api/v1/roulette/connect";
}

var socket = new WebSocket(uri);

let connect = () => {
  console.log("Attempting Connection...");

  socket.onopen = () => {
    console.log("Successfully Connected");
  };

  socket.onmessage = (msg) => {
    let data = JSON.parse(msg.data);
    if (data.errorType === 401) {
      refreshToken().then(() => {
        connect();
      });
      return;
    }
  };

  socket.onclose = (event) => {
    console.log("Socket Closed Connection: ", event);
  };

  socket.onerror = (error) => {
    console.log("Socket Error: ", error);
  };
};

let sendMsg = (msg) => {
  console.log("sending msg: ", msg);
  socket.send(msg);
};

export { connect, sendMsg, socket };
