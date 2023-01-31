package rpc

import (
	"context"
	"douyin_mini/kitex_gen/user"
	"douyin_mini/kitex_gen/user/userservice"
	"douyin_mini/pkg/constants"
	"douyin_mini/pkg/errno"
	"douyin_mini/pkg/middleware"
	etcd "github.com/kitex-contrib/registry-etcd"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var userClient userservice.Client

func initUserRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := userservice.NewClient(
		constants.UserServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

// CreateUser create user info
func CreateUser(ctx context.Context, req *user.RegisterUserRequest) error {
	//发出rpc请求
	resp, err := userClient.RegisterUser(ctx, req)
	//处理错误
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

// CheckUser check user info
//检查用户信息的接口

func CheckUser(ctx context.Context, req *user.CheckUserRequest) (int64, error) {
	//发送rpc请求
	resp, err := userClient.CheckUser(ctx, req)
	//处理两种错误
	if err != nil {
		return 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.UserId, nil
}

func GetUserInfo(ctx context.Context, req *user.UserInfoRequest) ([]*user.User, error) {
	resp, err := userClient.GetUserInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.Users, nil

}
