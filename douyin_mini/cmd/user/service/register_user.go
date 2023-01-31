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

type RegisterUserService struct {
	ctx context.Context
}

// NewRegisterUserService RegisterUserService 注册用户
// 初始化
func NewRegisterUserService(ctx context.Context) *RegisterUserService {
	return &RegisterUserService{ctx: ctx}
}

// CreateUser create user info.
// 具体的CreateUser的业务逻辑
func (s *RegisterUserService) CreateUser(req *user.RegisterUserRequest) error {
	//查询有没有用户
	users, err := db.QueryUser(s.ctx, req.Username)
	//错误处理
	if err != nil {
		return err
	}
	//有用户就不能创建
	if len(users) != 0 {
		return errno.UserAlreadyExistErr
	}
	//创建一个md5工具
	h := md5.New()
	//用h加密工具去加密密码，最终过的加密结果存在h工具当中
	if _, err = io.WriteString(h, req.Password); err != nil {
		return err
	}
	//去h工具中获取passWord
	passWord := fmt.Sprintf("%x", h.Sum(nil))
	//创建用户，调用db的接口
	return db.CreateUser(s.ctx, []*db.User{{
		Name:     req.Username,
		Password: passWord,
	}})
}
