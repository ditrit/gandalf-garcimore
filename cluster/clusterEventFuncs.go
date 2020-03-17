package cluster

import (
	"shoset/msg"
	"shoset/net"
)

// HandleEvent :
func HandleEvent(c *net.ShosetConn, message msg.Message) error {
	evt := message.(msg.Event)
	ch := c.GetCh()
	thisOne := ch.GetBindAddr()

	ch.Queue["evt"].Push(evt, c.ShosetType, c.GetBindAddr())

	ch.ConnsByType.Get("a").Iterate(
		func(key string, val *net.ShosetConn) {
			if key != c.GetBindAddr() && key != thisOne {
				val.SendMessage(evt)
			}
		},
	)
	return nil
}
