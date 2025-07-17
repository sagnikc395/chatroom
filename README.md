## chatroom

An implementation of a chatroom application written in Python,
to explore the asynchronous I/O module, efficient data structures
and to handle concurrent client connections.

It allows multiple clients to connect to a single server
and exchange messages in a chat room.

### Key Ideas I want to explore :

- Async I/O with asyncio -> Avoids blocking threads on socket I/O
- TCP sockets -> Most chat apps use TCP to ensure message delivery.
- Pub/Sub pattern -> To efficiently broadcast messages.
- Concurrency -> Using async def , not threads or processes.
- Scalability -> Later stages can add multiprocessing or a Redis backend.


### Running :

1. Run the server:
```shell
python server.py
```

2. Run the client/(s) :
```shell
python client.py
```
