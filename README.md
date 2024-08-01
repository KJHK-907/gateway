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

- Blockers:
    - How can we send out initial metadata to client when they first connect?
        - Since the TCP Server only sends out data when it recives it
        - We will need to store current metadata in memory and send it out to client when they first connect

Program scheduling 
    - history of programs (changes every semester) - update json every semester
        - Name of program
        - DJ
        - description
        - time
        - day
    - history JSON current sem JSON or each sem JSON
    - route to fetch programs for the semester

- Notes:
    - If my websocket server would handle 10,000 clients then how do I decide what buffer size should the channels in the Pool struct be?
    - gracefully shutdown the websocket server in case of interrupt signal timeout and context
    - load testing with 10,000 clients
    - clear old entries in the broadcast buffer when the buffer is full
    - How do I send out metadata to the client when they first connect?
    - Check for duplicate entries before sending out metadata in the broadcast buffer
    - Check if it is a Link or a Song
