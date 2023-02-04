package services

import (
	"context"
	"douyinV/cmd/user/dal/mysqlDB"
	"douyinV/cmd/user/kitex_gen/user"
	"douyinV/pkg/errno"
	"errors"
	"gorm.io/gorm"
)

func Login(ctx context.Context, req user.LoginRequest) (int64, error) {
	// 获取参数
	username := req.Username
	password := req.Password
	// 查询数据库
	id, err := mysqlDB.UserLogin(ctx, username, password)
	if err != nil {
		// 如果未找到该账号
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, errno.AuthorizationFailedErr
		} else {
			// 其他错误
			return 0, err
		}
	}
	return id, nil
}
