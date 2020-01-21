package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		fmt.Printf("	[options]\n options:\n")
		flag.PrintDefaults()
		fmt.Printf("  modes are exclusive\n")
	}

	var isConnector bool
	var isAggregator bool
	var isCluster bool
	var isTest bool

	flag.BoolVar(&isConnector, "c", false, "Connector mode (shorthand)")
	flag.BoolVar(&isConnector, "connector", false, "Connector mode")
	flag.BoolVar(&isAggregator, "a", false, "Aggregator mode (shorthand)")
	flag.BoolVar(&isAggregator, "aggregator", false, "Aggregator mode")
	flag.BoolVar(&isCluster, "cl", false, "Cluster mode (shorthand)")
	flag.BoolVar(&isCluster, "s", false, "Cluster mode (shorthand)")
	flag.BoolVar(&isCluster, "cluster", false, "Cluster mode")
	flag.BoolVar(&isTest, "t", false, "Test mode (shorthand)")
	flag.BoolVar(&isTest, "test", false, "Test mode")
	flag.Parse()

	args := flag.Args()

	if isTest == true {
		test()
		return
	}

	nbArgs := len(args)
	nbModes := 0

	// 1 unique argument pour les connecteurs : leur nom logique
	if isConnector == true {
		nbModes++
		if nbArgs == 2 {
			logicalName := args[0]
			aggregatorAddress := args[1]
			connector(logicalName, aggregatorAddress)
		} else {
			flag.Usage()
			os.Exit(1)
		}
	}

	// 1 arguments pour un aggrégateur : le tenant et l'adresse d'un nombre
	if isAggregator == true {
		nbModes++
		if nbArgs == 2 {
			tenantName := args[0]
			clusterAddress := args[1]
			aggregator(tenantName, clusterAddress)
		} else {
			flag.Usage()
			os.Exit(1)
		}
	}

	if isCluster == true {
		nbModes++
		if nbArgs == 1 || nbArgs == 2 {
			bindAddress := args[0]
			joinAddress := ""
			if nbArgs == 2 {
				joinAddress = args[1]
			}
			cluster(bindAddress, joinAddress)
		} else {
			flag.Usage()
			os.Exit(1)
		}
	}

	if nbModes != 1 {
		flag.Usage()
		os.Exit(1)
	}
}
