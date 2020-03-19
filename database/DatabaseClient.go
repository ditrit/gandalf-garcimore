package database

import (
	"gandalf-ui/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//var databaseClient *gorm.DB

func NewDatabaseClient(tenant string) *gorm.DB {
	databaseClient, err := gorm.Open("sqlite3", tenant+".db")
	if err != nil {
		panic("failed to connect database")
	}
	InitTenantDatabase(databaseClient)

	return databaseClient
}

/* func GetDatabaseClient *gorm.DB {
	var err error
	if databaseClient == nil {
		databaseClient, err = gorm.Open("sqlite3", "context.db")
		if err != nil {
			panic("failed to connect database")
		}
	}
	return databaseClient
} */

func InitTenantDatabase(databaseClient *gorm.DB) (err error) {

	databaseClient.AutoMigrate(&models.Aggregator{}, &models.Application{}, &models.Cluster{}, &models.CommandType{},
		&models.ConnectorType{}, &models.Connector{}, &models.Organisation{}, &models.Role{}, &models.Tenant{}, &models.User{})

	databaseClient.Create(&models.Aggregator{Name: "Aggregator1"})
	databaseClient.Create(&models.Aggregator{Name: "Aggregator2"})

	databaseClient.Create(&models.Application{Name: "Application1"})
	databaseClient.Create(&models.Application{Name: "Application2"})

	databaseClient.Create(&models.Cluster{Name: "Cluster1"})
	databaseClient.Create(&models.Cluster{Name: "Cluster2"})

	databaseClient.Create(&models.CommandType{Name: "Command_type1"})
	databaseClient.Create(&models.CommandType{Name: "Command_type1"})

	databaseClient.Create(&models.ConnectorType{Name: "Connector_type1"})
	databaseClient.Create(&models.ConnectorType{Name: "Connector_type2"})

	databaseClient.Create(&models.Connector{Name: "Connector1"})
	databaseClient.Create(&models.Connector{Name: "Connector2"})

	databaseClient.Create(&models.Tenant{Name: "Tenant1"})
	databaseClient.Create(&models.Tenant{Name: "Tenant2"})

	return
}
