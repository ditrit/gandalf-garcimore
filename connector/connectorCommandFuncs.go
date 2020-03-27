package connector

import (
	"fmt"
	"shoset/msg"
	"shoset/net"
)

// HandleCommand :
func HandleCommand(c *net.ShosetConn, message msg.Message) error {
	cmd := message.(msg.Command)
	ch := c.GetCh()

	fmt.Println("HANDLE COMMAND")
	fmt.Println("QUEUEU")
	fmt.Println(ch.Queue["cmd"])
	ch.Queue["cmd"].Push(cmd, c.ShosetType, c.GetBindAddr())
	fmt.Println("QUEUEU2")
	fmt.Println(ch.Queue["cmd"])

	return nil
}
