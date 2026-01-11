package client

import (
	"io"
	"log"
	"net"
	"os"
)

//client handles 2 things concurrently
// reading from the network
// and reading from the keyboard

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})

	//goroutine to handle the background server output
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()

	io.Copy(conn, os.Stdin)
	conn.Close()
	<-done // waiting for the background goroutine to finish
}
