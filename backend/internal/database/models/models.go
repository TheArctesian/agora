package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string
	Avatar      string
	Connections []Connection `gorm:"foreignKey:SourceID"`
}

type Connection struct {
	gorm.Model
	SourceID uint
	Source   User `gorm:"foreignKey:SourceID"`
	TargetID uint
	Target   User `gorm:"foreignKey:TargetID"`
	Strength float64
	Type     string
}
