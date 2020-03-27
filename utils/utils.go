package utils

import (
	"fmt"
	"shoset/msg"
	"shoset/net"
)

func GetByType(m *net.MapSafeConn, shosetType string) []*net.ShosetConn {
	fmt.Println("GET")
	fmt.Println(m.GetM())
	var result []*net.ShosetConn
	//m.Lock()
	for _, val := range m.GetM() {
		fmt.Println("GET BY TYPE")
		fmt.Println(val)
		if val.ShosetType == shosetType {
			result = append(result, val)
		}
	}
	//m.Unlock()
	return result
}

func GetByTenant(m *net.MapSafeConn, tenant string) []*net.ShosetConn {

	var result []*net.ShosetConn
	m.Lock()
	for _, val := range m.GetM() {
		if val.GetCh().Context["tenant"] == tenant {
			result = append(result, val)
		}
	}
	m.Unlock()
	return result
}

func CreateValidationEvent(command msg.Command) *msg.Event {
	var tab = map[string]string{
		"topic":          "validation",
		"event":          "validation",
		"payload":        "",
		"referencesUUID": command.GetUUID()}

	return msg.NewEvent(tab)
}
