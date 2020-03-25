package connector

import (
	"fmt"
	"shoset/msg"
	"shoset/net"
)

// HandleEvent :
func HandleEvent(c *net.ShosetConn, message msg.Message) error {
	evt := message.(msg.Event)
	ch := *c.GetCh()
	//dir := c.GetDir()
	//thisOne := ch.GetBindAddr()
	fmt.Println("HANDLE EVENT")
	fmt.Println("QUEUEU")
	fmt.Println(ch.Queue["evt"])
	ch.Queue["evt"].Push(evt, c.ShosetType, c.GetBindAddr())
	fmt.Println("QUEUEU2")
	fmt.Println(ch.Queue["evt"])
	//ch.Queue["evt"].Print()
	/* 	//fmt.Printf("%s : event 'join' received from %s\n", thisOne, newMember)
	   	if dir == "in" {
	   		//QUEUE
	   		//TODO REMOTE ADD ??
	   		ch.Queue["evt"].Push(evt, c.ShosetType, c.GetBindAddr())
	   	} */

	/* 	if dir == "out" {
		//GRPC
	} */
	return nil
}
