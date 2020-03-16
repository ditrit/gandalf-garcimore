package aggregator

import (
	"shoset/msg"
	"shoset/net"
)

// HandleCommand :
func HandleCommand(c *net.ShosetConn, message msg.Message) error {
	cmd := message.(msg.Command)
	ch := c.GetCh()
	dir := c.GetDir()
	thisOne := ch.GetBindAddr()

	if dir == "in" {
		//TODO GET NAME FROM MESSAGE
		name := "toto"
		ch.ConnsByName.Get(name).Iterate(
			func(key string, val *net.ShosetConn) {
				if key != "TODO TARGET" && key != thisOne {
					val.SendMessage(cmd)
				}
			},
		)
	}

	if dir == "out" {
		ch.ConnsByType.Get("cl").Iterate(
			func(key string, val *net.ShosetConn) {
				if key != "TODO TARGET" && key != thisOne {
					val.SendMessage(cmd)
				}
			},
		)
	}
	return nil
}
