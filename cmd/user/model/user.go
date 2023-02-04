package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID            int64 `gorm:"primaryKey"`
	Name          string
	Password      string
	FollowCount   int32
	FollowerCount int32
	CUDTime       `gorm:"embedded"`
}
type Relation struct {
	UserId   int64
	FollowId int64
	CUDTime  `gorm:"embedded"`
}
type CUDTime struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
