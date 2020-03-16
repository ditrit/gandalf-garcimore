package main

import (
	"fmt"
	"garcimore/database"
	"shoset/msg"
	"shoset/net"
	"time"
)

// ClusterMember :
type ClusterMember struct {
	chaussette   *net.Shoset
	databaseNode *database.DatabaseNode
	Store        *[]string
}

// NewClusterMember :
func NewClusterMember(logicalName string) *ClusterMember {
	member := new(ClusterMember)
	member.chaussette = net.NewShoset(logicalName, "cl")
	member.chaussette.Handle["cfgjoin"] = HandleConfigJoin

	return member
}

// Bind :
func (m *ClusterMember) Bind(addr string) error {
	ipAddr, err := net.GetIP(addr)
	if err == nil {
		err = m.chaussette.Bind(ipAddr)
	}
	return err
}

// Join :
func (m *ClusterMember) Join(addr string) (*net.ShosetConn, error) {
	return m.chaussette.Join(addr)
}

// Link :
func (m *ClusterMember) Link(addr string) (*net.ShosetConn, error) {
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

func databaseInit(add string, id int) {
	databaseNode := database.NewDatabaseNode(database.DefaultNodeDirectory, add, uint64(id))
	databaseNode.Run()
}

func clusterInit(logicalName, bindAddress string) (clusterMember *ClusterMember) {
	member := NewClusterMember(logicalName)
	member.Bind(bindAddress)

	time.Sleep(time.Second * time.Duration(5))
	fmt.Printf("%s.JoinBrothers Init(%#v)\n", bindAddress, getBrothers(bindAddress, member))

	return clusterMember
}

func clusterJoin(logicalName, bindAddress, joinAddress string) (clusterMember *ClusterMember) {

	member := NewClusterMember(logicalName)
	member.Bind(bindAddress)
	member.Join(joinAddress)

	time.Sleep(time.Second * time.Duration(5))
	fmt.Printf("%s.JoinBrothers Join(%#v)\n", bindAddress, getBrothers(bindAddress, member))

	member.Store = CreateStore(getBrothers(bindAddress, member))
	fmt.Println("Store")
	fmt.Println(member.Store)

	return member
}

func CreateStore(bros []string) *[]string {
	store := []string{}

	for _, bro := range bros {
		thisDBBro, ok := net.DeltaAddress(bro, 1000)
		if ok {
			store = append(store, thisDBBro)
		}
	}

	return &store
}

// HandleConfigJoin :
func HandleConfigJoin(c *net.ShosetConn, message msg.Message) error {
	cfg := message.(msg.ConfigJoin)
	ch := c.GetCh()
	dir := c.GetDir()
	thisOne := ch.GetBindAddr()
	newMember := cfg.GetBindAddress() // recupere l'adresse distante

	switch cfg.GetCommandName() {
	case "join":
		//fmt.Printf("%s : event 'join' received from %s\n", thisOne, newMember)
		if dir == "in" {
			ch.Join(newMember)
		}
		cfgNewMember := msg.NewCfgMember(newMember)
		ch.ConnsJoin.Iterate(
			func(key string, val *net.ShosetConn) {
				if key != newMember && key != thisOne {
					val.SendMessage(cfgNewMember)
					// fmt.Printf("%s : send event new 'member' %s to %s\n", thisOne, newMember, key)
				}
			},
		)
		if dir == "out" {
		}

	case "member":
		//fmt.Printf("%s : event 'member' received from %s\n", thisOne, newMember)
		ch.Join(newMember)
	}
	return nil
}
