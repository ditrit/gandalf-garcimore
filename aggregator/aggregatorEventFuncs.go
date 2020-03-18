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

	//TODO VERIF TENANT
	if evt.GetTenant() == ch.Context["tenant"] {
		ch.Queue["evt"].Push(evt, c.ShosetType, c.GetBindAddr())

		//fmt.Printf("%s : event 'join' received from %s\n", thisOne, newMember)
		if dir == "in" {
			ch.ConnsByType.Get(c.GetShosetType()).Iterate(
				func(key string, val *net.ShosetConn) {
					if key != c.GetBindAddr() && key != thisOne {
						val.SendMessage(evt)
						// fmt.Printf("%s : send event new 'member' %s to %s\n", thisOne, newMember, key)
					}
				},
			)

		}
		/* 	if dir == "out" {
			ch.ConnsByType.Get("cl").Iterate(
				func(key string, val *net.ShosetConn) {
					if key != c.GetBindAddr() && key != thisOne {
						val.SendMessage(evt)
						// fmt.Printf("%s : send event new 'member' %s to %s\n", thisOne, newMember, key)
					}
				},
			)
		} */
	}

	return nil
}
