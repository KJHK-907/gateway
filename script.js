const net = require("net");

const IP_PORT = 1234;

const server = net.createServer((socket) => {
	console.log("Client connected");
	socket.on("data", (data) => {
		console.log(`Data received: ${data}`);
	});
	socket.on("end", () => {
		console.log("Client disconnected");
	});
});

server.listen(IP_PORT, () => {
	console.log(`Server listening on port ${IP_PORT}`);
});