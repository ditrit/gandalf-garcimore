package models

import (
	"github.com/jinzhu/gorm"
)

type Application struct {
	gorm.Model
	Name          string        `form:"name" json:"name" binding:"required" gorm:"type:varchar(255);not null"`
	Tenant        Tenant        `gorm:"foreignkey:Name"`
	Aggregator    Aggregator    `gorm:"foreignkey:Name"`
	Connector     Connector     `gorm:"foreignkey:Name"`
	ConnectorType ConnectorType `gorm:"foreignkey:Name"`
	CommandType   CommandType   `gorm:"foreignkey:Name"`
}
