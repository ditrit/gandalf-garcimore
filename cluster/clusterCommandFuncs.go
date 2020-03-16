package cluster

import (
	"shoset/msg"
	"shoset/net"
)

// HandleCommand :
func HandleCommand(c *net.ShosetConn, message msg.Message) error {
	cmd := message.(msg.Command)
	ch := c.GetCh()
	//dir := c.GetDir()
	thisOne := ch.GetBindAddr()

	//TODO REQUEST
	name := "toto"
	//TODO MAJ MESSAGE ROUTER
	ch.ConnsByName.Get(name).Iterate(
		func(key string, val *net.ShosetConn) {
			if key != "TODO TARGET" && key != thisOne {
				val.SendMessage(cmd)
			}
		},
	)

	return nil
}
