package rpc

import (
	"douyinV/kitex_gen/favorite/favoritesvr"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var (
	FavoriteSeverName = ""
	FavoriteClient    favoritesvr.Client
)

func initFavoriteRpc() error {
	r, err := etcd.NewEtcdResolver([]string{EtcdAddress})
	if err != nil {
		hlog.Fatal("etcdResolver 初始化失败")
		return err
	}
	c, err := favoritesvr.NewClient(
		FavoriteSeverName,
		client.WithMuxConnection(MuxConnection),
		client.WithRPCTimeout(Timeout),                    // rpc timeout
		client.WithConnectTimeout(ConnectTimeout),         // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),
	)
	if err != nil {
		hlog.Fatal("%s 初始化失败", FavoriteSeverName)
		return err
	}
	FavoriteClient = c
	return nil
}
