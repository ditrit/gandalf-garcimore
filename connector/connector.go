package connector

import (
	"fmt"
	"shoset/msg"
	"shoset/net"
	"time"
)

// ConnectorMember :
type ConnectorMember struct {
	chaussette *net.Shoset
}

// NewClusterMember :
func NewConnectorMember(logicalName string) *ConnectorMember {
	member := new(ConnectorMember)
	member.chaussette = net.NewShoset(logicalName, "ag")
	member.chaussette.Handle["cfgjoin"] = HandleConfigJoin

	return member
}

// Bind :
func (m *ConnectorMember) Bind(addr string) error {
	ipAddr, err := net.GetIP(addr)
	if err == nil {
		err = m.chaussette.Bind(ipAddr)
	}
	return err
}

// Join :
func (m *ConnectorMember) Join(addr string) (*net.ShosetConn, error) {
	return m.chaussette.Join(addr)
}

// Link :
func (m *ConnectorMember) Link(addr string) (*net.ShosetConn, error) {
	return m.chaussette.Link(addr)
}

func getBrothers(address string, member *ConnectorMember) []string {
	bros := []string{address}
	member.chaussette.ConnsJoin.Iterate(
		func(key string, val *net.ShosetConn) {
			bros = append(bros, key)
		})
	return bros
}

func ConnectorMemberInit(logicalName, bindAddress string) (connectorMember *ConnectorMember) {
	member := NewConnectorMember(logicalName)
	member.Bind(bindAddress)

	time.Sleep(time.Second * time.Duration(5))
	fmt.Printf("%s.JoinBrothers Init(%#v)\n", bindAddress, getBrothers(bindAddress, member))

	return member
}

func ConnectorMemberJoin(logicalName, bindAddress, joinAddress string) (connectorMember *ConnectorMember) {

	member := NewConnectorMember(logicalName)
	member.Bind(bindAddress)
	member.Join(joinAddress)

	time.Sleep(time.Second * time.Duration(5))
	fmt.Printf("%s.JoinBrothers Join(%#v)\n", bindAddress, getBrothers(bindAddress, member))

	return member
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
