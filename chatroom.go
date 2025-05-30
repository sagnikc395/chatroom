package main

import (
	"net"

	"github.com/PullRequestInc/iostream"
	"github.com/gammazero/deque"
	"github.com/golang-collections/collections/set"
)

type Participant interface {
	Deliver(message *Message)
	Write(message *Message)
}

type Room interface {
	Join(participant *Participant)
	Leave(participant *Participant)
	Deliver(participant *Participant, message *Message)
}

type RoomStore struct {
	//since golang interfaces can't store concrete types,
	// defining the private members of the room in a struct type
	MessageQueue deque.Deque[Message]
	Participants set.Set
}

type SessionStore struct {
	ClientSocket net.UDPAddr
	Buffer       iostream.StreamBuffer
	Room         *RoomStore
	MessageQueue deque.Deque[Message]
}

type Session interface {
	NewSession(s net.UDPAddr, room *RoomStore)
	Start()
	Deliver(message *Message)
	Write(message *Message)
	AsyncRead()
	AsyncWrite(messageBody string, messsageLegnth uint)
}
