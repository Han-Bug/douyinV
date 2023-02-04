package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"tiktok/biz/model"
	"tiktok/config"
)

var DB *gorm.DB

func Init() {
	db, err := gorm.Open(mysql.Open(config.MysqlDsn), &gorm.Config{})
	if err != nil {
		panic("mysql连接失败！" + err.Error())
	}

	if err = db.AutoMigrate(&model.Message{}); err != nil {
		panic("mysql自动迁移失败！" + err.Error())
	}

	DB = db
}
