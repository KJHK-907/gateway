# Kjhk 90.7 FM API Gateway

This project implements an API gateway and a WebSocket server that listens for metadata from a TCP server and broadcasts it to connected clients. The primary use case is to provide real-time updates to clients using the KJHK App on Android/iOS about new media playing on a Zetta RCS system.

## Features
- **API Gateway**: Connects client to the appropriate service (room for future services such as DJ Scheduling)
- **TCP Server**: Listens for metadata sent from Zetta RCS.
  - Supported metadata:
    - New Media Playing
- **WebSocket Server**: Sends metadata to all connected clients.
    - Example Response: 
    ```json
    {
        "track": "I Know It's Over",
        "album": "The Queen Is Dead",
        "artist": "The Smiths",
        "length": 347893.3,
        "timestampUTC": "2024-09-24T00:32:22Z",
        "timestampCST": "2024-09-23T19:32:22Z"
    }
    ```

## Connection Workflow

1. **Client Connection**: 
   - The client (e.g., an app) connects to the Nginx server on port 80.
   - The client mentions the target endpoint (metadata) in the request.
2. **Nginx Reverse Proxy**:
   - Nginx forwards the WebSocket connection to the API gateway on port 8081.
3. **Subscription Confirmation**:
   - The API gateway then routes the connection to the particular service requested by the client (metadata).
   - The WebSocket server (running on port 8080) sends metadata once a successful connection is established.
   - The server also sends out the most recent metadata to the client once they connect.
4. **Metadata Reception**:
   - The client receives metadata from the WebSocket server.
5. **Heartbeat Mechanism**:
   - The client sends a heartbeat to the Nginx server to keep the connection alive.
6. **Client Disconnection**:
   - The client disconnects from the WebSocket server.
7. **Server Disconnection**:
   - Server has graceful shutdown mechanism to ensure it disconnects all clients and performs necessary cleanups.

## Getting Started

### Prerequisites

- Docker
- Docker Compose

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/KJHK-907/gateway.git
    ```
2. Running the Services:
    ```sh
    docker compose up --build
    ```
3. Client connection example (Node.js)
    ```javascript
    const WebSocket = require('ws');

    const targetEndpoint = "metadata";
    const ws = new WebSocket(`ws://localhost/api/?target=${targetEndpoint}`);

    ws.on('open', function open() {
    console.log('Connected to the server');
    setInterval(() => {
        ws.ping();
    }, 30000);
    });

    ws.on('message', function incoming(data) {
    console.log('Received from server:', data);
    });

    ws.on('error', function error(err) {
    console.error('WebSocket error:', err);
    });

    ws.on('close', function close() {
    console.log('Disconnected from the server');
    });
    ```

### License

This project is licensed under the MIT License.