package mysqlDB

import (
	"gorm.io/gorm"
	"time"
)

// 所有数据库结构均未设计外键，而是通过业务逻辑来保证外键约束
// 数据库设计时假设所有微服务使用不同的数据库

type Video struct {
	ID       int64 `gorm:"primaryKey"`
	UserId   int64 `gorm:"index"`
	PlayUrl  string
	CoverUrl string
	Title    string
	CUDTime  `gorm:"embedded"`
}
type CUDTime struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
