package main

import (
	"../chaussette/net"
)

func connector(logicalName string, aggregatorAddress string) {
	chaussette := net.NewChaussette(logicalName)
	if aggregatorAddress != "" {
		chaussette.Connect(aggregatorAddress)
	}
}
