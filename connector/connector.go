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
func NewConnectorMember(logicalName, tenant string) *ConnectorMember {
	member := new(ConnectorMember)
	member.chaussette = net.NewShoset(logicalName, "c")
	member.chaussette.Context["tenant"] = tenant
	member.chaussette.Handle["cfgjoin"] = HandleConfigJoin
	member.chaussette.Handle["cmd"] = HandleCommand
	member.chaussette.Handle["event"] = HandleEvent
	//member.connectorGrpc = NewConnectorGrpc("", member.chaussette.)
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

func ConnectorMemberInit(logicalName, tenant, bindAddress, linkAddress string) (connectorMember *ConnectorMember) {
	member := NewConnectorMember(logicalName, tenant)
	member.Bind(bindAddress)
	member.Link(linkAddress)

	time.Sleep(time.Second * time.Duration(5))
	fmt.Printf("%s.JoinBrothers Init(%#v)\n", bindAddress, getBrothers(bindAddress, member))

	return member
}

func ConnectorMemberJoin(logicalName, tenant, bindAddress, linkAddress, joinAddress string) (connectorMember *ConnectorMember) {

	member := NewConnectorMember(logicalName, tenant)
	member.Bind(bindAddress)
	member.Link(linkAddress)
	member.Join(joinAddress)

	time.Sleep(time.Second * time.Duration(5))
	fmt.Printf("%s.JoinBrothers Join(%#v)\n", bindAddress, getBrothers(bindAddress, member))

	return member
}
