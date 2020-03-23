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
		if cmd.GetTenant() == ch.Context["tenant"] {
			ch.Queue["cmd"].Push(cmd, c.ShosetType, c.GetBindAddr())
			if c.GetShosetType() == "cl" {
				ch.ConnsByName.Get(cmd.GetTarget()).Iterate(
					func(key string, val *net.ShosetConn) {
						if key != c.GetBindAddr() && key != thisOne {
							val.SendMessage(cmd)
						}
					},
				)
			} else if c.GetShosetType() == "c" {
				ch.ConnsByType.Get("cl").Iterate(
					func(key string, val *net.ShosetConn) {
						if key != c.GetBindAddr() && key != thisOne {
							val.SendMessage(cmd)
						}
					},
				)
			}

		}
	}

	/* 	if dir == "out" {
		ch.ConnsByType.Get("cl").Iterate(
			func(key string, val *net.ShosetConn) {
				if key != "TODO TARGET" && key != thisOne {
					val.SendMessage(cmd)
				}
			},
		)
	} */
	return nil
}
