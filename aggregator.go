package main

import "../shoset/net"

func aggregator(logicalName string, clusterAddress string) {
	chaussette := net.NewShoset(logicalName)
	if clusterAddress != "" {
		chaussette.Connect(clusterAddress)
	}
}
