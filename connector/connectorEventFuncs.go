package connector

import (
	"shoset/msg"
	"shoset/net"
)

// HandleEvent :
func HandleEvent(c *net.ShosetConn, message msg.Message) error {
	evt := message.(msg.Event)
	ch := c.GetCh()
	dir := c.GetDir()
	//thisOne := ch.GetBindAddr()

	//fmt.Printf("%s : event 'join' received from %s\n", thisOne, newMember)
	if dir == "in" {
		//QUEUE
		//TODO REMOTE ADD ??
		ch.Queue["evt"].Push(evt, "a", "")
	}

	/* 	if dir == "out" {
		//GRPC
	} */
	return nil
}
