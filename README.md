## chatroom 

An chatroom application built in Go.

### Why ?
I wanted to understand the building of a conccurent chatroom application to explore in depths the process of concurrent programming and also to deep dive into topics like channels and coroutines.

### Architecture

#### System Architecture (High Level)
1. Server -> listens for incoming TCP connections.For every new client, it spins up a goroutine.
2. Broker / Hub -> A central loop that receives messages ffrom one client and broadcasts them to all others.
3. Client -> A simple CLI that runs 2 concurrent processes: one to read from `stdin` and send it to the server, and one to listen for incoming messages from the server.

#### Low level System Design 
TODO! 


