package aggregator

import (
	"fmt"
	"shoset/msg"
	"shoset/net"
)

// HandleEvent :
func HandleEvent(c *net.ShosetConn, message msg.Message) error {
	evt := message.(msg.Event)
	ch := c.GetCh()
	dir := c.GetDir()
	thisOne := ch.GetBindAddr()
	fmt.Println("HANDLE EVENT")
	fmt.Println(evt)
	fmt.Println("TYpe")
	fmt.Println(c.GetShosetType())
	//TODO VERIF TENANT
	if evt.GetTenant() == ch.Context["tenant"] {
		ok := ch.Queue["evt"].Push(evt, c.ShosetType, c.GetBindAddr())
		fmt.Println("ok")
		fmt.Println(ok)
		if ok {
			fmt.Println("dir")
			fmt.Println(dir)
			//fmt.Printf("%s : event 'join' received from %s\n", thisOne, newMember)
			if dir == "in" {
				ch.ConnsByAddr.Iterate(
					func(key string, val *net.ShosetConn) {

						if key != c.GetBindAddr() && key != thisOne && val.ShosetType == "cl" {
							val.SendMessage(evt)
							// fmt.Printf("%s : send event new 'member' %s to %s\n", thisOne, newMember, key)
						}
					},
				)

			}
			if dir == "out" {
				ch.ConnsByAddr.Iterate(
					func(key string, val *net.ShosetConn) {

						if key != c.GetBindAddr() && key != thisOne && val.ShosetType == "c" {
							val.SendMessage(evt)
							// fmt.Printf("%s : send event new 'member' %s to %s\n", thisOne, newMember, key)
						}
					},
				)
			}
		}
	}

	return nil
}
