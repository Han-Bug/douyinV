package mysqlDB

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID       int64 `gorm:"primaryKey"`
	Name     string
	Password string
	CUDTime  `gorm:"embedded"`
}
type Relation struct {
	UserId     int64
	FollowerId int64
	CUDTime    `gorm:"embedded"`
}
type CUDTime struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
