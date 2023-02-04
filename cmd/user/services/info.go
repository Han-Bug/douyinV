package services

import (
	"context"
	"douyinV/cmd/user/dal/mysqlDB"
	"douyinV/cmd/user/kitex_gen/user"
	"douyinV/pkg/errno"
	"errors"
	"gorm.io/gorm"
)

func Info(ctx context.Context, req *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	infoRes := new(user.UserInfoResponse)
	var err error
	if req.Id == 0 {
		infoRes, err = mysqlDB.UserInfoAnonymous(ctx, req.ToId)
	} else {
		infoRes, err = mysqlDB.UserInfo(ctx, req.Id, req.ToId)
	}
	if err != nil {
		// 如果不存在该id的账号
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return infoRes, errno.AuthorizationFailedErr
		}
		// 其他错误
		return infoRes, err
	}
	return infoRes, nil

}
