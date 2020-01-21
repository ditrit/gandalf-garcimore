package main

import "../chaussette/net"

func aggregator(logicalName string, clusterAddress string) {
	chaussette := net.NewChaussette(logicalName)
	if clusterAddress != "" {
		chaussette.Connect(clusterAddress)
	}
}
