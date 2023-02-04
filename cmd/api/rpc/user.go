package rpc

import (
	"context"
	"douyinV/cmd/user/kitex_gen/user"
	"douyinV/cmd/user/kitex_gen/user/usersvr"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

const (
	EtcdAddress = ""
	UserServiceName
)

var userClient usersvr.Client

func initUserRpc() {
	r, err := etcd.NewEtcdResolver([]string{EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := usersvr.NewClient(
		UserServiceName,
		//client.WithMiddleware(middleware.CommonMiddleware),
		//client.WithInstanceMW(middleware.ClientMiddleware),
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

func UserRegister(ctx context.Context, req *user.RegisterRequest) (*user.RegisterResponse, error) {
	//发出rpc请求
	resp, err := userClient.Register(ctx, req)
	// TODO 处理错误
	if err != nil {
		return nil, err
	}

	//	TODO 待修正
	//if resp.BaseResp.StatusCode != 0 {
	//	//return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	//}
	return resp, nil
}
func UserLogin(ctx context.Context, req *user.LoginRequest) (*user.LoginResponse, error) {
	//发出rpc请求
	resp, err := userClient.Login(ctx, req)
	//处理错误
	if err != nil {
		return nil, err
	}
	// TODO 待修正
	//if resp.BaseResp.StatusCode != 0 {
	//	return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	//}
	return resp, nil
}
func UserInfo(ctx context.Context, req *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	//发出rpc请求
	resp, err := userClient.UserInfo(ctx, req)
	//处理错误
	if err != nil {
		return nil, err
	}
	// TODO 待修正
	//if resp.BaseResp.StatusCode != 0 {
	//	return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	//}
	return resp, nil
}
