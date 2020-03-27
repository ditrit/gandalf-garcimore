package cluster

import (
	"fmt"
	"garcimore/database"
	"garcimore/models"
	"garcimore/utils"
	"shoset/msg"
	"shoset/net"

	"github.com/jinzhu/gorm"
)

var sendIndex = 0

// HandleCommand :
func HandleCommand(c *net.ShosetConn, message msg.Message) error {
	cmd := message.(msg.Command)
	ch := c.GetCh()
	//dir := c.GetDir()
	//thisOne := ch.GetBindAddr()

	fmt.Println("HANDLE COMMAND")

	mapDatabaseClient := ch.Context["database"].(map[string]*gorm.DB)
	/* fmt.Println("MAp")
	fmt.Println(mapDatabaseClient)

	fmt.Println("MApZ")
	fmt.Println(GetDatabaseClientByTenant(cmd.GetTenant(), mapDatabaseClient))
	*/
	app := GetApplicationContext(cmd, GetDatabaseClientByTenant(cmd.GetTenant(), mapDatabaseClient))
	fmt.Println("app")
	fmt.Println(app)
	if app != (models.Application{}) {
		fmt.Println("CONNECTOR")
		fmt.Println(app.Connector)
		cmd.Target = app.Connector

		fmt.Println("Aggregator")
		fmt.Println(app.Aggregator)

		fmt.Println("NAME")
		fmt.Println(ch.ConnsByName.Get("agg2"))
		fmt.Println("NAME2")
		fmt.Println(ch.ConnsByName.Get(app.Aggregator))

		shosets := utils.GetByType(ch.ConnsByName.Get(app.Aggregator), "a")
		fmt.Println(shosets)
		index := getSendIndex(shosets)
		fmt.Println(index)

		shosets[index].SendMessage(cmd)
		fmt.Println(shosets[index])

		/*
			ch.ConnsByName.Get(app.Aggregator).Iterate(
				func(key string, val *net.ShosetConn) {
					if key != c.GetBindAddr() && key != thisOne && c.GetCh().Context["tenant"] == val.GetCh().Context["tenant"] {
						val.SendMessage(cmd)
						//WAIT REP
					}
				},
			) */
	}

	return nil
}

func getSendIndex(conns []*net.ShosetConn) int {
	aux := sendIndex
	sendIndex++
	if sendIndex >= len(conns) {
		sendIndex = 0
	}
	return aux
}

// GetDatabaseClientByTenant
func GetDatabaseClientByTenant(tenant string, mapDatabaseClient map[string]*gorm.DB) *gorm.DB {
	if _, ok := mapDatabaseClient[tenant]; !ok {
		mapDatabaseClient[tenant] = database.NewDatabaseClient(tenant)
	}
	return mapDatabaseClient[tenant]
}

// GetDatabaseClientByTenant
func GetApplicationContext(cmd msg.Command, client *gorm.DB) (applicationContext models.Application) {
	fmt.Println("TYPE")
	fmt.Println(cmd.GetContext()["ConnectorType"].(string))
	var connectortype models.ConnectorType
	//client.Where("name = ?", cmd.GetContext()["ConnectorType"].(string)).First(&connectortype)
	fmt.Println("connectortype")
	fmt.Println(connectortype)
	//client.Model(&connectortype).Related(&applicationContext)
	client.Where("connector_type = ?", cmd.GetContext()["ConnectorType"].(string)).First(&applicationContext)

	return
}
