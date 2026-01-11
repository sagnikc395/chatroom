package server

import (
	"bufio"
	"fmt"
	"net"
)

//uses the "Hub" pattern
// it will manage the clients using channels rather than mutexes to avoid manual memory locking

// outoging message channel
type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	//the incoming messages
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			// broadcast the incoming messages to all cleints outgoing channels
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {
	//outgoing client messages
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived 🥳"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	leaving <- ch
	messages <- who + " has left the chat!"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func main() {
	fmt.Println("Chatroom server starting on :8080 ...")
	listener, _ := net.Listen("tcp", "localhost:8080")

	go broadcaster()
	for {
		//event loop working
		conn, _ := listener.Accept()
		go handleConn(conn)
	}
}
