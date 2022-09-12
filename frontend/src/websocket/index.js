let sendMsg = (ws, msg) => {
  console.log("sending msg: ", msg);
  ws.send(msg);
};

export { sendMsg };
