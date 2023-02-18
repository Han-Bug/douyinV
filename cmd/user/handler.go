package main

import (
	"context"
	user "douyinV/user/kitex_gen/user"
	"douyinV/user/pkg"
	"douyinV/user/pkg/errno"
	"douyinV/user/services"
)

// UserSvrImpl implements the last service interface defined in the IDL.
type UserSvrImpl struct{}

// TODO 统一rpc错误格式

// Register implements the UserSvrImpl interface.
func (s *UserSvrImpl) Register(ctx context.Context, req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	resp = &user.RegisterResponse{
		Id:       0,
		BaseResp: nil,
	}
	id, err := services.Register(ctx, *req)
	if err != nil {
		resp.BaseResp = pkg.BuildBaseResp(err)
		return resp, nil
	}
	resp.Id = id
	resp.BaseResp = pkg.BuildBaseResp(errno.Success)

	return resp, nil
}

// Login implements the UserSvrImpl interface.
func (s *UserSvrImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	resp = &user.LoginResponse{
		Id:       0,
		BaseResp: nil,
	}
	id, err := services.Login(ctx, *req)
	if err != nil {
		resp.BaseResp = pkg.BuildBaseResp(err)
		return resp, nil
	}
	// id为0表示不存在该用户
	resp.Id = id
	resp.BaseResp = pkg.BuildBaseResp(errno.Success)

	return resp, nil
}

// UserInfo implements the UserSvrImpl interface.
func (s *UserSvrImpl) UserInfo(ctx context.Context, req *user.UserInfoRequest) (resp *user.UserInfoResponse, err error) {
	// id为0表示匿名搜索信息
	resp = &user.UserInfoResponse{
		Id:            0,
		Name:          "",
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
		BaseResp:      nil,
	}
	info, err := services.Info(ctx, req)
	if err != nil {
		resp.BaseResp = pkg.BuildBaseResp(err)
		return resp, nil
	}
	resp = &user.UserInfoResponse{
		Id:            info.Id,
		Name:          info.Name,
		FollowCount:   info.FollowCount,
		FollowerCount: info.FollowerCount,
		IsFollow:      info.IsFollow,
		BaseResp:      nil,
	}
	resp.BaseResp = pkg.BuildBaseResp(errno.Success)

	return resp, nil
}
