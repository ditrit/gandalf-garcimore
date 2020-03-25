package cluster

import (
	"fmt"
	"shoset/msg"
	"shoset/net"
)

// HandleEvent :
func HandleEvent(c *net.ShosetConn, message msg.Message) error {
	evt := message.(msg.Event)
	ch := c.GetCh()
	thisOne := ch.GetBindAddr()
	fmt.Println("HANDLE EVENT")
	dir := c.GetDir()
	fmt.Println(dir)

	ok := ch.Queue["evt"].Push(evt, c.ShosetType, c.GetBindAddr())
	fmt.Println("ok")
	fmt.Println(ok)
	if ok {
		ch.ConnsByAddr.Iterate(
			func(key string, val *net.ShosetConn) {
				if key != c.GetBindAddr() && key != thisOne && val.ShosetType == "a" && c.GetCh().Context["tenant"] == val.GetCh().Context["tenant"] {
					val.SendMessage(evt)
				}
			},
		)
	}

	return nil
}
