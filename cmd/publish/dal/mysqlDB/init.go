package mysqlDB

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dsn = ""
)

var DB *gorm.DB

// 打印错误并上抛
// TODO 将fmt输出改为日志输出

func Init() error {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("mysql初始化失败")
		return err
	}
	DB = db
	return nil
}
