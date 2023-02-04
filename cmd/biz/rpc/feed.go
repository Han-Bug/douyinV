package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"log"
	"tiktok/config"
	"tiktok/feed/kitex_gen/api"
	"tiktok/feed/kitex_gen/api/feed"
	"time"
)

var feedClient feed.Client

func initFeed() {
	fc, err := feed.NewClient("feed", client.WithHostPorts(config.ServerFeedUrlHP))
	if err != nil {
		panic("连接feed服务失败！" + err.Error())
	}
	feedClient = fc
}

func Feed() *api.FeedResp {
	req := &api.FeedReq{}
	resp, err := feedClient.Feed(context.Background(), req, callopt.WithRPCTimeout(3*time.Second), callopt.WithConnectTimeout(3*time.Second))
	if err != nil {
		msg := "获取视频流失败！"
		log.Println(msg + err.Error())
		return &api.FeedResp{
			StatusCode: -1,
			StatusMsg:  &msg,
		}
	}
	return resp
}
