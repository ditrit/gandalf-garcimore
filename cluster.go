package main

import (
	"fmt"
	"shoset/msg"
	"shoset/net"
	"time"
)

// ClusterMember :
type ClusterMember struct {
	chaussette   *net.Shoset
	databaseNode *DatabaseNode
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
		thisDBAddr, ok := net.DeltaAddress(ipAddr, 1000)
		if ok {
			//m.databaseClient = NewDatabaseClient()

			id, ok := net.IP2ID(thisDBAddr)
			if ok {
				m.databaseNode = NewDatabaseNode(id, thisDBAddr) //ID
				m.chaussette.Context["node"] = m.databaseNode
				go m.databaseNode.run()
			}
		}
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

func getBrothers(address string, member *ClusterMember) []string {
	bros := []string{address}
	member.chaussette.ConnsJoin.Iterate(
		func(key string, val *net.ShosetConn) {
			bros = append(bros, key)
		})
	return bros
}

func clusterInit(logicalName, bindAddress string) {
	done := make(chan bool)
	member := NewClusterMember(logicalName)
	member.Bind(bindAddress)
<<<<<<< HEAD
	time.Sleep(time.Second * time.Duration(5))
	fmt.Printf("%s.JoinBrothers(%#v)\n", bindAddress, getBrothers(bindAddress, member))
=======
>>>>>>> e303311788882db921a6d844aaa9edd516bf810e
	<-done

}

func clusterJoin(logicalName, bindAddress, joinAddress string) {
	done := make(chan bool)
	member := NewClusterMember(logicalName)
	member.Bind(bindAddress)
	member.Join(joinAddress)
<<<<<<< HEAD
	time.Sleep(time.Second * time.Duration(5))
	fmt.Printf("%s.JoinBrothers(%#v)\n", bindAddress, getBrothers(bindAddress, member))
=======
>>>>>>> e303311788882db921a6d844aaa9edd516bf810e
	<-done
}

// HandleConfigJoin :
func HandleConfigJoin(c *net.ShosetConn, message msg.Message) error {
	cfg := message.(msg.ConfigJoin)
	ch := c.GetCh()
	dir := c.GetDir()
	thisOne := ch.GetBindAddr()
	newMember := cfg.GetBindAddress() // recupere l'adresse distante
	bros := []string{thisOne}
	ch.ConnsJoin.Iterate(
		func(key string, val *net.ShosetConn) {
			bros = append(bros, key)
		})
	fmt.Printf("%s.JoinBrothers(%#v)\n", thisOne, bros)
	switch cfg.GetCommandName() {
	case "join":
		//fmt.Printf("%s : event 'join' received from %s\n", thisOne, newMember)
		if dir == "in" {
			ch.Join(newMember)
		}
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
						// fmt.Printf("%s : send event new 'member' %s to %s\n", thisOne, newMember, key)
					}
				},
			)
<<<<<<< HEAD
			//fmt.Printf("store : %#v\n", store)
=======
			fmt.Printf("store : %#v\n", store)
			node := ch.Context["node"].(*DatabaseNode)
			node.clusterDatabaseClient.Cluster = store
			err := node.addNodesToLeader()
			fmt.Println(err)
>>>>>>> e303311788882db921a6d844aaa9edd516bf810e
		}

		if dir == "out" {
		}

	case "member":
		//fmt.Printf("%s : event 'member' received from %s\n", thisOne, newMember)
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
