package mysqlDB

import (
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	ID      int64 `gorm:"primaryKey"`
	VideoId int64 `gorm:"index"`
	UserId  int64 `gorm:"index"`
	Content string
	CUDTime `gorm:"embedded"`
}

type CUDTime struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
