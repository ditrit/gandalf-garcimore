package connector

import (
	"fmt"
	"shoset/net"
	"time"
)

// ConnectorMember :
type ConnectorMember struct {
	chaussette    *net.Shoset
	connectorGrpc ConnectorGrpc
}

// NewClusterMember :
func NewConnectorMember(logicalName string) *ConnectorMember {
	member := new(ConnectorMember)
	member.chaussette = net.NewShoset(logicalName, "c")
	member.chaussette.Handle["cfgjoin"] = HandleConfigJoin
	member.chaussette.Handle["cmd"] = HandleCommand
	member.chaussette.Handle["event"] = HandleEvent

	return member
}

// Bind :
func (m *ConnectorMember) Bind(addr string) error {
	ipAddr, err := net.GetIP(addr)
	if err == nil {
		err = m.chaussette.Bind(ipAddr)
	}
	//TODO
	//member.connectorGrpc = NewConnectorGrpc()

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
