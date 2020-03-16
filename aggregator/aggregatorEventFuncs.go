package aggregator

import (
	"shoset/msg"
	"shoset/net"
)

// HandleEvent :
func HandleEvent(c *net.ShosetConn, message msg.Message) error {
	evt := message.(msg.Event)
	ch := c.GetCh()
	dir := c.GetDir()
	thisOne := ch.GetBindAddr()

	//fmt.Printf("%s : event 'join' received from %s\n", thisOne, newMember)
	if dir == "in" {
		ch.ConnsByType.Get("c").Iterate(
			func(key string, val *net.ShosetConn) {
				if key != "TODO TARGET" && key != thisOne {
					val.SendMessage(evt)
					// fmt.Printf("%s : send event new 'member' %s to %s\n", thisOne, newMember, key)
				}
			},
		)
	}
	if dir == "out" {
		ch.ConnsByType.Get("cl").Iterate(
			func(key string, val *net.ShosetConn) {
				if key != "TODO TARGET" && key != thisOne {
					val.SendMessage(evt)
					// fmt.Printf("%s : send event new 'member' %s to %s\n", thisOne, newMember, key)
				}
			},
		)
	}
	return nil
}
