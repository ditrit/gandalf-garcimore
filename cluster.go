package main

import (
	"fmt"
	"shoset/msg"
	"shoset/net"
)

// ClusterMember :
type ClusterMember struct {
	chaussette     *net.Shoset
	databaseNode   *DatabaseNode
	databaseClient *DatabaseClient
}

// NewClusterMember :
func NewClusterMember(logicalName string) *ClusterMember {
	member := new(ClusterMember)
	member.chaussette = net.NewShoset(logicalName, "cl")
	member.chaussette.Handle["cfgjoin"] = HandleConfigJoin
	return member
}

// Bind :
func (m ClusterMember) Bind(addr string) error {
	ipAddr, err := net.GetIP(addr)
	if err == nil {
		err = m.chaussette.Bind(ipAddr)
		//	m.databaseNode = NewDatabaseNode(0, addr) //ID
		//	m.databaseNode.run()
	}
	return err
}

// Join :
func (m ClusterMember) Join(addr string) (*net.ShosetConn, error) {
	return m.chaussette.Join(addr)
}

// Link :
func (m ClusterMember) Link(addr string) (*net.ShosetConn, error) {
	return m.chaussette.Link(addr)
}

func clusterInit(logicalName, bindAddress string) {
	done := make(chan bool)
	member := NewClusterMember(logicalName)
	member.Bind(bindAddress)

	<-done

}

func clusterJoin(logicalName, bindAddress, joinAddress string) {
	done := make(chan bool)
	member := NewClusterMember(logicalName)
	member.Bind(bindAddress)
	member.Join(joinAddress)

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
		thisOne := ch.GetBindAddr()
		cfgNewMember := msg.NewCfgMember(newMember)

		thisDBAddr, ok := net.DeltaAddress(thisOne, 1000)
		if ok {
			store := []string{thisDBAddr}
			ch.ConnsJoin.Iterate(
				func(key string, val *net.ShosetConn) {
					dbKey, ok := net.DeltaAddress(key, 1000)
					if ok {
						store = append(store, dbKey)
					}
					if key != newMember && key != thisOne {
						val.SendMessage(cfgNewMember)
					}
				},
			)
			fmt.Printf("store : %#v\n", store)
		}

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
