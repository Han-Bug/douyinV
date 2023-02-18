package rpc

import (
	"context"
	"douyinV/api/global"
	"douyinV/api/kitex_gen/user"
	"douyinV/api/kitex_gen/user/usersvr"
	"douyinV/api/pkg/errno"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"time"
)

const (
	MuxConnection = 1
)

var (
	Timeout        = 3 * time.Second
	ConnectTimeout = 50 * time.Millisecond
)

//var userClient usersvr.Client
// 错误上抛

func initUserRpc() error {
	r, err := etcd.NewEtcdResolver([]string{global.APIImpl.ApiConfig.EtcdConfig.EtcdAddress})
	if err != nil {
		return err
	}

	c, err := usersvr.NewClient(
		global.APIImpl.ApiConfig.EtcdConfig.UserServerName,
		//client.WithMiddleware(middleware.CommonMiddleware),
		//client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(MuxConnection),           // mux
		client.WithRPCTimeout(Timeout),                    // rpc timeout
		client.WithConnectTimeout(ConnectTimeout),         // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		return err
	}
	global.APIImpl.UserClient = c
	return nil

}

// 以下均为使用rpc客户端直接向服务端请求而无额外业务逻辑
// 直接调用global里的client

func UserRegister(ctx context.Context, req *user.RegisterRequest) (*user.RegisterResponse, error) {
	//发出rpc请求
	resp, err := global.APIImpl.UserClient.Register(ctx, req)
	//	TODO 处理RPC请求错误
	if err != nil {
		return nil, err
	}
	//	TODO 处理RPC结果错误
	if resp.BaseResp.StatusCode != 0 {
		// 此处是直接沿用来自RPC Server传来的错误代码与信息
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp, nil
}
func UserLogin(ctx context.Context, req *user.LoginRequest) (*user.LoginResponse, error) {
	//发出rpc请求
	resp, err := global.APIImpl.UserClient.Login(ctx, req)
	//	TODO 处理RPC请求错误
	if err != nil {
		return nil, err
	}
	//	TODO 处理RPC结果错误
	if resp.BaseResp.StatusCode != 0 {
		// 此处是直接沿用来自RPC Server传来的错误代码与信息
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp, nil
}
func UserInfo(ctx context.Context, req *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	//发出rpc请求
	resp, err := global.APIImpl.UserClient.UserInfo(ctx, req)
	//	TODO 处理RPC请求错误
	if err != nil {
		return nil, err
	}
	//	TODO 处理RPC结果错误
	if resp.BaseResp.StatusCode != 0 {
		// 此处是直接沿用来自RPC Server传来的错误代码与信息
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp, nil
}
