package database

import (
	"gandalf-ui/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var databaseClient *gorm.DB

func GetDatabaseClient() *gorm.DB {
	var err error
	if databaseClient == nil {
		databaseClient, err = gorm.Open("sqlite3", "context.db")
		if err != nil {
			panic("failed to connect database")
		}
	}
	return databaseClient
}

func InitDatabaseCluster() (err error) {

	GetDatabaseClient().AutoMigrate(&models.Aggregator{}, &models.Application{}, &models.Cluster{}, &models.CommandType{},
		&models.ConnectorType{}, &models.Connector{}, &models.Organisation{}, &models.Role{}, &models.Tenant{}, &models.User{})

	GetDatabaseClient().Create(&models.Aggregator{Name: "Aggregator1"})
	GetDatabaseClient().Create(&models.Aggregator{Name: "Aggregator2"})

	GetDatabaseClient().Create(&models.Application{Name: "Application1"})
	GetDatabaseClient().Create(&models.Application{Name: "Application2"})

	GetDatabaseClient().Create(&models.Cluster{Name: "Cluster1"})
	GetDatabaseClient().Create(&models.Cluster{Name: "Cluster2"})

	GetDatabaseClient().Create(&models.CommandType{Name: "Command_type1"})
	GetDatabaseClient().Create(&models.CommandType{Name: "Command_type1"})

	GetDatabaseClient().Create(&models.ConnectorType{Name: "Connector_type1"})
	GetDatabaseClient().Create(&models.ConnectorType{Name: "Connector_type2"})

	GetDatabaseClient().Create(&models.Connector{Name: "Connector1"})
	GetDatabaseClient().Create(&models.Connector{Name: "Connector2"})

	GetDatabaseClient().Create(&models.Tenant{Name: "Tenant1"})
	GetDatabaseClient().Create(&models.Tenant{Name: "Tenant2"})

	return
}
