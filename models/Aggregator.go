package models

import (
	"github.com/jinzhu/gorm"
)

type Aggregator struct {
	gorm.Model `form:"ID" json:"ID" binding:"required"`
	Name       string `form:"name" json:"name" binding:"required" gorm:"type:varchar(255);not null"`
}
