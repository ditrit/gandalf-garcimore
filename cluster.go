package main

import (
	"shoset/net"
)

func cluster(bindAddress string, joinAddress string) {
	chaussette := net.NewShoset("cluster", "cl")
	chaussette.Bind(bindAddress)
	if joinAddress != "" {
		chaussette.Join(joinAddress)
	}

}
