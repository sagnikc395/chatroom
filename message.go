package main

import "fmt"

/**
header -> 4 bytes and maxBytes can be stored as 512 bytes
header will store the body length , which is the current body length
data will be stored the header+bodyLength with maximum size of header+maxBytes

Client will attempt to send message -> It will encode header and put message into the data and send data
server gets the message, decodes the header, get the bodyLength from the header and hence complete body
then server sends the message to all the clients connected to that room
*/

type Message struct {
	Data       string
	BodyLength uint
}

func NewMessage(message string) *Message {
	bodyLength := NewBodyLength(len(message))
	EncodeHeader()
	Message{
		Data:       message,
		BodyLength: bodyLength + Header,
	}
}

func (m Message) PrintMessage() {
	message := m.GetData()
	fmt.Printf("Message recieved: %s\n", message)
}

func EncodeHeader() {
	
}

func NewBodyLength(newLength int) int {
	if newLength > int(MaxBytes) {
		return int(MaxBytes)
	}
	return newLength
}
