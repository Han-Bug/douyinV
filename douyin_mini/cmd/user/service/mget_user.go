package service

import (
	"douyin_mini/cmd/user/dal/db"
	"douyin_mini/cmd/user/pack"
	"douyin_mini/kitex_gen/user"
	"golang.org/x/net/context"
)

// MGetUserService 初始化service
type MGetUserService struct {
	ctx context.Context
}

// NewMGetUserService new MGetUserService
// 初始化
func NewMGetUserService(ctx context.Context) *MGetUserService {
	return &MGetUserService{ctx: ctx}
}

// MGetUser multiple get list of user info
// 具体的业务逻辑
func (s *MGetUserService) MGetUser(req *user.UserInfoRequest) ([]*user.User, error) {
	modelUsers, err := db.MGetUsers(s.ctx, req.UserIds)
	if err != nil {
		return nil, err
	}
	return pack.Users(modelUsers), nil
}
