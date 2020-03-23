package database

import (
	"garcimore/models"

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
		&models.ConnectorType{}, &models.Connector{}, &models.Tenant{})

	databaseClient.Create(&models.Application{Name: "Application1",
		ConnectorType: models.ConnectorType{Name: "Connector_type1"},
		CommandType:   models.CommandType{Name: "Command_type1"},
		Aggregator:    models.Aggregator{Name: "Aggregator1"},
		Connector:     models.Connector{Name: "Connector1"}})

	databaseClient.Create(&models.Application{Name: "Application2",
		ConnectorType: models.ConnectorType{Name: "Connector_type2"},
		CommandType:   models.CommandType{Name: "Command_type2"},
		Aggregator:    models.Aggregator{Name: "Aggregator2"},
		Connector:     models.Connector{Name: "Connector2"}})

	databaseClient.Create(&models.Aggregator{Name: "Aggregator1"})
	databaseClient.Create(&models.Aggregator{Name: "Aggregator2"})

	databaseClient.Create(&models.Connector{Name: "Connector1"})
	databaseClient.Create(&models.Connector{Name: "Connector2"})

	databaseClient.Create(&models.CommandType{Name: "Command_type1"})
	databaseClient.Create(&models.CommandType{Name: "Command_type1"})

	databaseClient.Create(&models.ConnectorType{Name: "Connector_type1"})
	databaseClient.Create(&models.ConnectorType{Name: "Connector_type2"})

	return
}
