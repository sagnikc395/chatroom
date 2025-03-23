## chatroom

An basic chatroom application written in Go, to explore high performance communication and to handle concurrent client connections.
It allows multiple clients to connect to a single server and exchange messages in a chat room - think of it like a mini discord , with a single server.

### Key Features:

- Async Communication
- Multi-Threading
- Room Management
- Message Handling

### Architecture:

This contains of the following key components:

1. Session : Represents a client session, handling communication with a specific client. It is responsible for reading incoming messages, delivering outgoing messages, and managing the client socket.
2. Room : Represents a chat room, managinng a set of connected clients. It is responsible for adding and removing clients, and delivering messages to all participants in the room.
3. Message : Represents a message exchanged between clients.It is responsible for encoding and decoding the message header and body.

### Critical Functionalities

#### Async I/O

Using channel and wait group, which allows the server to handle multiple client connections without blocking.

#### Threading

The application uses threading to handle concurrent client connections. Each client session is run in a seperate thread, allowing the server to handle multiple clients simulatenously.

#### Session Management

Session interface manages a client's connection to the chat room. It handles the async reading and writing of messages, and ensures that messages are delivered to the correct recipients.

#### Room Management

Room interface manages the chat room itself. It keeps track of all connected clients, and provides methods for adding and removing clients, and for delivering messages to all participants in the room.

### Build:

- To build and run the project:
  - Clone the repo.
  - Install `task` and run as task build to generate the binary
  - Run `task run server <port>` to start the server , replacing `<port>` with the port number you want to use.
  - Run `task run client <port>` to start a client, replacing `<port>` with the port number the server is running on.
