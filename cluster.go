package main

import (
	"fmt"
	"shoset/msg"
	"shoset/net"
	"time"
)

type clusterMember struct {
	chaussette     *net.Shoset
	databaseNode   *DatabaseNode
	databaseClient *DatabaseClient
}

func clusterInit(logicalName, bindAddress string) {
	done := make(chan bool)

	clusterMember := new(clusterMember)
	clusterMember.chaussette = net.NewShoset(logicalName, "cl")
	clusterMember.chaussette.Handle["cfgjoin"] = HandleConfigJoin
	clusterMember.chaussette.Bind(bindAddress)
	//TODO VOIR POUR ID
	clusterMember.databaseNode = NewDatabaseNode(0, bindAddress) //ID
	clusterMember.databaseNode.run()

	<-done

}

func clusterJoin(logicalName, bindAddress, joinAddress string) {
	done := make(chan bool)

	clusterMember := new(clusterMember)
	clusterMember.chaussette = net.NewShoset(logicalName, "cl")
	clusterMember.chaussette.Handle["cfgjoin"] = HandleConfigJoin
	clusterMember.chaussette.Bind(bindAddress)
	time.Sleep(time.Millisecond * time.Duration(50))
	clusterMember.chaussette.Join(joinAddress)
	//TODO VOIR POUR ID
	clusterMember.databaseNode = NewDatabaseNode(0, bindAddress) //ID
	clusterMember.databaseNode.run()

	<-done
}

// HandleConfigJoin :
func HandleConfigJoin(c *net.ShosetConn, message msg.Message) error {
	fmt.Println("NEW HANDLE")

	cfg := message.(msg.ConfigJoin)
	ch := c.GetCh()
	dir := c.GetDir()
	switch cfg.GetCommandName() {
	case "join":
		newMember := cfg.GetBindAddress() // recupere l'adresse distante

		if dir == "in" {
			ch.Join(newMember)
		}
		thisOne := c.GetBindAddr()
		cfgNewMember := msg.NewCfgMember(newMember)
		ch.ConnsJoin.Iterate(
			func(key string, val *net.ShosetConn) {
				if key != newMember && key != thisOne {
					//
					fmt.Printf("- %s : %s \n", thisOne, key)
					fmt.Println(key)
					val.SendMessage(cfgNewMember)
				}
			},
		)

		if dir == "out" {
		}

	case "member":
		newMember := cfg.GetBindAddress()
		ch.Join(newMember)
	}
	return nil
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
