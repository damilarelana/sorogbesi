
// make websocket connection
let socket = new WebSocket("ws://localhost:8080/ws")

let connectViaWebSocket = (callBack) => { // the argument `cb` here is a callBack function
  console.log("Attempting WebSocket connection ... ");

  // listen for connected socket events and take some actions
  socket.onopen = () => {
    console.log("Socket connection successful:");
  };

  socket.onmessage = (msg) => {
    console.log(msg);
    console.log('typeof callBack:', typeof callBack)
    callBack(msg);
  };

  // event for when websocket connection is closed
  socket.onclose = (event) => {
    console.log("Socket connection closed:", event);
  };

  // event for when the websocket connection throws an error
  socket.onerror = (error) => {
    console.log("Socket connection error:", error);
  };
};

let sendMessageViaWebSocket = (msg) => {
  console.log("sending message:", msg);
  socket.send(msg);
};

export {connectViaWebSocket, sendMessageViaWebSocket};