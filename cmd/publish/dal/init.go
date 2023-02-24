package dal

import (
	"douyinV/cmd/publish/dal/minio"
	"douyinV/cmd/publish/dal/mysqlDB"
)

func Init() error {
	var err error
	if err = minio.Init(); err != nil {
		return err
	}
	if err = mysqlDB.Init(); err != nil {
		return err
	}
	return nil
}
