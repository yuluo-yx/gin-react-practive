package model

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Id     uint   `gorm:"not null"`
	Text   string `gorm:"type:text"`
	UserId uint
}
