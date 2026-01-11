package server

import (
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

}

func handleConn(conn net.Conn) {

}

func clientWriter(conn net.Conn, ch <-chan string) {

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
