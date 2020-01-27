package main

import (
	"shoset/net"
)

func connector(logicalName string, aggregatorAddress string) {
	chaussette := net.NewShoset(logicalName, "c")
	if aggregatorAddress != "" {
		chaussette.Connect(aggregatorAddress)
	}
}
