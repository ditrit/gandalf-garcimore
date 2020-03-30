package connector

import (
	"fmt"
	"garcimore/utils"
	"shoset/msg"
	"shoset/net"
)

// HandleCommand :
func HandleCommand(c *net.ShosetConn, message msg.Message) error {
	cmd := message.(msg.Command)
	ch := c.GetCh()
	thisOne := ch.GetBindAddr()

	fmt.Println("HANDLE COMMAND")
	fmt.Println("QUEUEU")
	fmt.Println(ch.Queue["cmd"])
	ch.Queue["cmd"].Push(cmd, c.ShosetType, c.GetBindAddr())
	fmt.Println("QUEUEU2")
	fmt.Println(ch.Queue["cmd"])
	ch.ConnsByAddr.Iterate(
		func(key string, val *net.ShosetConn) {

			if key != thisOne && val.ShosetType == "a" {
				fmt.Println("SEND EVENT VALIE")
				val.SendMessage(utils.CreateValidationEvent(cmd, ch.Context["tenant"].(string)))
				// fmt.Printf("%s : send event new 'member' %s to %s\n", thisOne, newMember, key)
			}
		},
	)

	return nil
}
