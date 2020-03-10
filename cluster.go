package main

import (
	"fmt"
	"garcimore/database"
	"shoset/msg"
	"shoset/net"
	"strconv"
	"time"
)

// ClusterMember :
type ClusterMember struct {
	chaussette   *net.Shoset
	databaseNode *database.DatabaseNode
}

// NewClusterMember :
func NewClusterMember(logicalName string) *ClusterMember {
	member := new(ClusterMember)
	member.chaussette = net.NewShoset(logicalName, "cl")
	member.chaussette.Handle["cfgjoin"] = HandleConfigJoin

	//member.databaseNode = new(DatabaseNodeCluster)
	//database := database.NewDatabaseCluster("/tmp/", []string{"127.0.0.1:9000", "127.0.0.1:9001", "127.0.0.1:9002"})
	//database.Run()
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

func database2(add, id string) {
	done := make(chan bool)

	//id, _ := net.IP2ID(add)
	idi, _ := strconv.Atoi(id)
	databaseNode := database.NewDatabaseNode("/home/orness/db/", add, uint64(idi))
	databaseNode.Run()
	//time.Sleep(time.Second * time.Duration(5))

	<-done
}

func database3(add, id string) *database.DatabaseNode {

	//id, _ := net.IP2ID(add)
	idi, _ := strconv.Atoi(id)
	databaseNode := database.NewDatabaseNode("/home/orness/db/", add, uint64(idi))
	//time.Sleep(time.Second * time.Duration(5))
	return databaseNode
}

func database4(add, id string) {

	//id, _ := net.IP2ID(add)
	idi, _ := strconv.Atoi(id)
	databaseNode := database.NewDatabaseNode("/home/orness/db/", add, uint64(idi))
	//time.Sleep(time.Second * time.Duration(5))
	databaseNode.Run()
}

func clusterInit(logicalName, bindAddress string) {
	done := make(chan bool)
	member := NewClusterMember(logicalName)
	member.Bind(bindAddress)

	time.Sleep(time.Second * time.Duration(5))
	fmt.Printf("%s.JoinBrothers Init(%#v)\n", bindAddress, getBrothers(bindAddress, member))

	add, _ := net.DeltaAddress(bindAddress, 1000)
	id, _ := net.IP2ID(add)
	fmt.Println("TOTO")
	fmt.Println(id)
	fmt.Println(add)
	//databaseNode := database.NewDatabaseNode("/home/orness/db/", add, uint64(id))
	//member.databaseNode.Run()
	time.Sleep(time.Second * time.Duration(5))

	//member.databaseNode.DatabaseClient.DatabaseClientCluster = CreateStore(getBrothers(bindAddress, member))
	//err := member.databaseNode.addNodesToLeader()
	//fmt.Println(err)
	<-done

}

func clusterJoin(logicalName, bindAddress, joinAddress string) {
	done := make(chan bool)
	member := NewClusterMember(logicalName)
	member.Bind(bindAddress)
	member.Join(joinAddress)
	time.Sleep(time.Second * time.Duration(5))

	fmt.Printf("%s.JoinBrothers Join(%#v)\n", bindAddress, getBrothers(bindAddress, member))
	add, _ := net.DeltaAddress(bindAddress, 1000)
	id, _ := net.IP2ID(add)
	fmt.Println("TOTO")
	fmt.Println(id)
	fmt.Println(add)
	//databaseNode := database.NewDatabaseNode("/home/orness/db/", add, uint64(id))

	member.databaseNode.Run()
	time.Sleep(time.Second * time.Duration(5))

	//member.databaseNode.DatabaseClient.DatabaseClientCluster = CreateStore(getBrothers(bindAddress, member))
	//err := member.databaseNode.addNodesToLeader()
	//fmt.Println(err)

	<-done
}

func CreateStore(bros []string) []string {
	store := []string{}

	for _, bro := range bros {
		thisDBBro, ok := net.DeltaAddress(bro, 1000)
		if ok {
			store = append(store, thisDBBro)
		}
	}

	return store
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
