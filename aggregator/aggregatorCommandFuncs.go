package aggregator

import (
	"fmt"
	"garcimore/utils"
	"shoset/msg"
	"shoset/net"
	"time"
)

var sendIndex = 0

// HandleCommand :
func HandleCommand(c *net.ShosetConn, message msg.Message) error {
	cmd := message.(msg.Command)
	ch := c.GetCh()
	dir := c.GetDir()
	//thisOne := ch.GetBindAddr()
	fmt.Println("HANDLE COMMAND")
	fmt.Println(cmd)
	fmt.Println("TYpe")
	fmt.Println(c.GetShosetType())
	fmt.Println("dir")
	fmt.Println(dir)
	if dir == "in" {
		if cmd.GetTenant() == ch.Context["tenant"] {
			ok := ch.Queue["cmd"].Push(cmd, c.ShosetType, c.GetBindAddr())
			if ok {
				if c.GetShosetType() == "c" {
					fmt.Println("CON 1")

					//SEND VALIDATION
					c.SendMessage(utils.CreateValidationEvent(cmd))

					shosets := utils.GetByType(ch.ConnsByAddr, "cl")
					index := getSendIndex(shosets)
					var send = false
					for !send {
						shosets[index].SendMessage(cmd)
						timeoutSend := time.Duration((int(cmd.GetTimeout()) / len(shosets)))
						time.Sleep(timeoutSend * time.Millisecond)

						evt := ch.Queue["evt"].GetByUUID(cmd.GetUUID())
						if evt != nil {
							break
						}
					}
				}
			}
		}
	}
	if dir == "out" {
		if cmd.GetTenant() == ch.Context["tenant"] {
			ok := ch.Queue["cmd"].Push(cmd, c.ShosetType, c.GetBindAddr())
			if ok {
				if c.GetShosetType() == "cl" {
					fmt.Println("CLUSTER1")

					//SEND VALIDATION
					shosets := utils.GetByType(ch.ConnsByName.Get(cmd.GetTarget()), "c")
					index := getSendIndex(shosets)
					var send = false
					for !send {
						shosets[index].SendMessage(cmd)
						timeoutSend := time.Duration((int(cmd.GetTimeout()) / len(shosets)))
						time.Sleep(timeoutSend * time.Millisecond)

						evt := ch.Queue["evt"].GetByUUID(cmd.GetUUID())
						if evt != nil {
							break
						}
					}
					/*
						ch.ConnsByName.Get(cmd.GetTarget()).Iterate(
							func(key string, val *net.ShosetConn) {
								if key != c.GetBindAddr() && key != thisOne && val.ShosetType == "c" {
									val.SendMessage(cmd)
									//WAIT REP
								}
							},
						) */
				}
			}
		}
	}
	return nil
}

func getSendIndex(conns []*net.ShosetConn) int {
	aux := sendIndex
	sendIndex++
	if sendIndex >= len(conns) {
		sendIndex = 0
	}
	return aux
}
