package utils

import (
	"shoset/msg"
	"shoset/net"
)

func GetByType(m *net.MapSafeConn, shosetType string) []*net.ShosetConn {

	var result []*net.ShosetConn
	m.Lock()
	for _, val := range m.GetM() {
		if val.ShosetType == shosetType {
			result = append(result, val)
		}
	}
	m.Unlock()
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
