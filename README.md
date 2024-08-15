- TCP Server listens to metadata sent out from Zetta RCS
    - Supported metadata:
        - New Media Playing

- TCP Server sends out metadata to all connected clients via WebSocket

- Connection:
    - Client (App) connects to WebSocket Server and sends message to subscribe to metadata via Opcodes
    - WebSocket Server sends back message to confirm subscription and sends out metadata including heartbeat interval
    - Client (App) receives metadata from WebSocket Server
    - Client (App) sends heartbeat to WebSocket Server to keep connection alive
    - WebSocket Server sends heartbeat opcode to client occasionally to keep connection alive
    - Client (App) disconnects from WebSocket Server

- Opcodes:
    - Heartbeat
    - HeartbeatAck

=====

Program scheduling 
- history of programs (changes every semester) - update json every semester
    - Name of program
    - DJ
    - description
    - time
    - day
- history JSON current sem JSON or each sem JSON
- route to fetch programs for the semester

### Notes:
- (TODO) If my websocket server would handle 10,000 clients then how do I decide what buffer size should the channels in the Pool struct be?
- (TODO) load testing with 10,000 clients
- (TODO) Check if it is a Link or a Song

- (DONE) Test client connection by creating a client script
- (DONE) How do I send out metadata to the client when they first connect?
- (DONE) Should the websocket server be a goroutine or a main routine
- (DONE) Check for duplicate entries in the recentTrackInfo map before sending out metadata in the broadcast buffer - time period is 10 minutes
- (DONE) clear old entries in the recentTrackInfo map when the map is full 
- (DONE) If Zetta did not send anything in the last 10 minutes then attempt reconnection
- (DONE) gracefully shutdown the websocket server in case of interrupt signal timeout and context

- (Note) Use caffeinate to prevent the system from going to sleep causing connection issues