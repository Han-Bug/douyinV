package main

import (
	"context"
	"douyin_mini/cmd/user/pack"
	"douyin_mini/cmd/user/service"
	user "douyin_mini/kitex_gen/user"
	"douyin_mini/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// RegisterUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) RegisterUser(ctx context.Context, req *user.RegisterUserRequest) (resp *user.RegisterUserResponse, err error) {
	resp = new(user.RegisterUserResponse)

	//参数校验
	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	service.NewRegisterUserService(ctx).CreateUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)

	return resp, nil
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *user.CheckUserRequest) (resp *user.CheckUserResponse, err error) {
	resp = new(user.CheckUserResponse)

	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	uid, err := service.NewCheckUserService(ctx).CheckUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.UserId = uid
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, req *user.UserInfoRequest) (resp *user.UserInfoResponse, err error) {
	resp = new(user.UserInfoResponse)
	if len(req.UserIds) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	users, err := service.NewMGetUserService(ctx).MGetUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Users = users
	return resp, nil
	return
}
