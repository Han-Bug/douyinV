package test

import (
	"douyinV/user/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestConnectDB(t *testing.T) {
	dsn := "root:password@tcp(47.108.237.99:3311)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	users := make([]model.User, 16)
	res := db.Find(&users)
	if res.Error != nil {
		panic(res.Error)
	}
}
