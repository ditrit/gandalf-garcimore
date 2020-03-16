package connector

import (
	"shoset/msg"
	"shoset/net"
)

// HandleCommand :
func HandleCommand(c *net.ShosetConn, message msg.Message) error {
	cmd := message.(msg.Command)
	ch := c.GetCh()
	dir := c.GetDir()
	//thisOne := ch.GetBindAddr()

	//fmt.Printf("%s : event 'join' received from %s\n", thisOne, newMember)
	if dir == "in" {
		//QUEUE
		//TODO REMOTE ADD ??
		ch.Queue["cmd"].Push(cmd, "a", "")
	}

	/* 	if dir == "out" {
		//GRPC
	} */
	return nil
}
