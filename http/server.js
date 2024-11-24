const WebSocket = require('ws');
const socket = new WebSocket('ws://localhost:7001/ws');

// Connection established
socket.onopen = function(event) {
  console.log('Connected to WebSocket server!');
  socket.send('Hello, server!');
};

// Message received from the server
socket.onmessage = function(event) {
  console.log('Message from server: ', event.data);
};

// Error handling
socket.onerror = function(error) {
  console.error('WebSocket Error: ', error);
};

// Connection closed
socket.onclose = function(event) {
  if (event.wasClean) {
    console.log('Closed cleanly, code= ' + event.code);
  } else {
    console.error('Connection error');
  }
};
