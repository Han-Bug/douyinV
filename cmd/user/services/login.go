package services

import (
	"context"
	"douyinV/user/dal/mysqlDB"
	"douyinV/user/kitex_gen/user"
	"douyinV/user/pkg/errno"
	utils "douyinV/user/pkg/md5"
	"errors"
	"gorm.io/gorm"
	"strings"
)

func Login(ctx context.Context, req user.LoginRequest) (int64, error) {
	// 获取参数
	username := req.Username
	password := req.Password
	// 密码加密
	pd := utils.EncryptMd5(strings.Join([]string{Md5key, password}, ""))

	// 查询数据库
	id, err := mysqlDB.UserLogin(ctx, username, pd)
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
