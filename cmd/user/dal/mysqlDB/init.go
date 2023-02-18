package mysqlDB

import (
	"douyinV/user/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
)

var DB *gorm.DB

func Init(c *config.MySqlConfig) error {
	var err error
	//dsn2 := c.GetDSN()
	//fmt.Println(dsn2)
	dsn := "root:password@tcp(47.108.237.99:3311)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		})
	if err != nil {
		return err
	}
	if err = DB.Use(gormopentracing.New()); err != nil {
		return err
	}
	return nil
}
