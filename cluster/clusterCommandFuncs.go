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
	fmt.Println("HANDLE COMMAND")
	fmt.Println(cmd)

	mapDatabaseClient := ch.Context["database"].(map[string]*gorm.DB)

	app := GetApplicationContext(cmd, GetDatabaseClientByTenant(cmd.GetTenant(), mapDatabaseClient))

	if app != (models.Application{}) {

		cmd.Target = app.Connector
		shosets := utils.GetByType(ch.ConnsByName.Get(app.Aggregator), "a")
		index := getSendIndex(shosets)
		shosets[index].SendMessage(cmd)
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

	client.Where("connector_type = ?", cmd.GetContext()["ConnectorType"].(string)).First(&applicationContext)

	return
}
