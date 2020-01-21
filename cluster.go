package main

import (
	"../chaussette/net"
)

func cluster(bindAddress string, joinAddress string) {
	chaussette := net.NewChaussette("cluster")
	chaussette.Bind(bindAddress)
	if joinAddress != "" {
		chaussette.Join(joinAddress)
	}

}
