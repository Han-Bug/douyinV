package dal

import (
	"douyinV/user/config"
	"douyinV/user/dal/mysqlDB"
)

func Init(c *config.DataBaseConfig) error {

	err := mysqlDB.Init(c.MySql)
	if err != nil {
		return err
	}
	return nil
}
