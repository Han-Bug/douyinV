package service

import (
	"context"
	"crypto/md5"
	"douyin_mini/cmd/user/dal/db"
	"douyin_mini/kitex_gen/user"
	"douyin_mini/pkg/errno"
	"fmt"

	"io"
)

type CheckUserService struct {
	ctx context.Context
}

// NewCheckUserService new CheckUserService
// 初始化
func NewCheckUserService(ctx context.Context) *CheckUserService {
	return &CheckUserService{
		ctx: ctx,
	}
}

// CheckUser check user info
// 检查用户信息的业务逻辑
func (s *CheckUserService) CheckUser(req *user.CheckUserRequest) (int64, error) {
	//创建一个md5
	h := md5.New()
	//是不是匹配md5格式
	if _, err := io.WriteString(h, req.Password); err != nil {
		return 0, err
	}
	Password := fmt.Sprintf("%x", h.Sum(nil))
	//rpc请求中直接获取用户名
	Username := req.Username
	//根据用户名查询用户
	users, err := db.QueryUser(s.ctx, Username)
	//处理异常
	if err != nil {
		return 0, err
	}
	//如果没有用户，异常
	if len(users) == 0 {
		return 0, errno.AuthorizationFailedErr
	}
	u := users[0]
	//比对密码
	if u.Password != Password {
		return 0, errno.AuthorizationFailedErr
	}
	//返回id
	return int64(u.ID), nil
}
