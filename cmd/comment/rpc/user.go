package rpc

import (
	"context"
	"douyinV/kitex_gen/user"
	"douyinV/kitex_gen/user/usersvr"
	"errors"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var (
	UserSeverName = ""

	UserClient usersvr.Client
)

func initUserRpc() error {
	r, err := etcd.NewEtcdResolver([]string{EtcdAddress})
	if err != nil {
		return err
	}
	c, err := usersvr.NewClient(
		UserSeverName,
		client.WithMuxConnection(MuxConnection),
		client.WithRPCTimeout(Timeout),                    // rpc timeout
		client.WithConnectTimeout(ConnectTimeout),         // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),
	)
	if err != nil {
		return err
	}
	UserClient = c
	return nil
}

func Info(ctx context.Context, req *user.InfoRequest) (*user.InfoResponse, error) {

	infoResp, err := UserClient.Info(ctx, req)
	if err != nil {
		return nil, err
	}
	if infoResp.StatusCode != 0 {
		return nil, errors.New(infoResp.StatusMessage)
	}
	return infoResp, nil
}
