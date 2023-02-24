package rpc

import (
	"douyinV/kitex_gen/comment/commentsvr"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var (
	CommentSeverName = ""

	CommentClient commentsvr.Client
)

func initCommentRpc() error {
	r, err := etcd.NewEtcdResolver([]string{EtcdAddress})
	if err != nil {
		return err
	}
	c, err := commentsvr.NewClient(
		CommentSeverName,
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
	CommentClient = c
	return nil
}
