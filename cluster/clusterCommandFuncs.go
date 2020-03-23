package cluster

import (
	"garcimore/database"
	"garcimore/models"
	"shoset/msg"
	"shoset/net"

	"github.com/jinzhu/gorm"
)

// HandleCommand :
func HandleCommand(c *net.ShosetConn, message msg.Message) error {
	cmd := message.(msg.Command)
	ch := c.GetCh()
	//dir := c.GetDir()
	thisOne := ch.GetBindAddr()
	mapDatabaseClient := ch.Context["database"].(map[string]*gorm.DB)

	app := GetApplicationContext(cmd, mapDatabaseClient[cmd.GetTenant()])
	cmd.Target = app.Connector.Name

	ch.ConnsByName.Get(app.Aggregator.Name).Iterate(
		func(key string, val *net.ShosetConn) {
			if key != c.GetBindAddr() && key != thisOne && c.GetCh().Context["tenant"] == val.GetCh().Context["tenant"] {
				val.SendMessage(cmd)
			}
		},
	)

	return nil
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
	client.Where("tenant = ? AND connectorType = ?", cmd.Tenant, cmd.GetContext()["ConnectorType"]).First(&applicationContext)

	return
}
