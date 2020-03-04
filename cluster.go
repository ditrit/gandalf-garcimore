package main

import (
	"shoset/net"
)

type clusterMember struct {
	chaussette     *net.Shoset
	databaseNode   *DatabaseNode
	databaseClient *DatabaseClient
}

func clusterInit(logicalName, bindAddress string) {
	clusterMember := new(clusterMember)
	clusterMember.chaussette = net.NewShoset(logicalName, "cl")
	clusterMember.chaussette.Bind(bindAddress)
	//TODO VOIR POUR ID
	clusterMember.databaseNode = NewDatabaseNode(0, bindAddress) //ID
	clusterMember.databaseNode.run()

}

func clusterJoin(logicalName, bindAddress, joinAddress string) {
	clusterMember := new(clusterMember)
	clusterMember.chaussette = net.NewShoset(logicalName, "cl")
	clusterMember.chaussette.Bind(bindAddress)
	clusterMember.chaussette.Join(joinAddress)
	//TODO VOIR POUR ID
	clusterMember.databaseNode = NewDatabaseNode(0, bindAddress) //ID
	clusterMember.databaseNode.run()
}

/* func clusterJoin(bindAddress string, joinAddress string) {
	chaussette := net.NewShoset("cluster")
	chaussette.Bind(bindAddress)
	if joinAddress != "" {
		chaussette.Join(joinAddress)
	}

}

func cluster_old(bindAddress string, joinAddress string) {
	chaussette := net.NewShoset("cluster")
	chaussette.Bind(bindAddress)
	if joinAddress != "" {
		chaussette.Join(joinAddress)
	}

}
*/
