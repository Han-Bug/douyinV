package model

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	FromUserId int    `gorm:"not null"`
	ToUserId   int    `gorm:"not null"`
	Content    string `gorm:"not null"`
}
