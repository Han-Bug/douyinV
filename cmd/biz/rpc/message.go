package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	tracer "github.com/kitex-contrib/tracer-opentracing"
	"log"
	"tiktok/config"
	"tiktok/message/kitex_gen/api"
	"tiktok/message/kitex_gen/api/message"
	"time"
)

var messageClient message.Client

var ChatOnlineMap = make(map[int64]int64)

func initMessage() {
	mc, err := message.NewClient(
		"message",
		client.WithHostPorts(config.ServerMsgUrlHP),
		client.WithSuite(tracer.NewDefaultClientSuite()),
	)
	if err != nil {
		panic("连接message服务失败！" + err.Error())
	}
	messageClient = mc
}

func Chat(fromUID, toUID int) *api.ChatResp {
	fUIDi64 := int64(fromUID)
	if ChatOnlineMap[fUIDi64] == -1 {
		return &api.ChatResp{
			StatusCode:  0,
			MessageList: nil,
		}
	}
	req := &api.ChatReq{
		FromUserId: fUIDi64,
		ToUserId:   int64(toUID),
		LatestTime: ChatOnlineMap[fUIDi64],
	}
	resp, err := messageClient.Chat(context.Background(), req, callopt.WithRPCTimeout(3*time.Second), callopt.WithConnectTimeout(3*time.Second))
	if err != nil {
		msg := "获取聊天记录失败！"
		log.Println(msg + err.Error())
		return &api.ChatResp{
			StatusCode: -1,
			StatusMsg:  &msg,
		}
	}
	if len(resp.MessageList) > 0 {
		ChatOnlineMap[fUIDi64] = resp.MessageList[len(resp.MessageList)-1].CreateTime
	}
	return resp
}

func Action(fromUID, toUID, actionType int, content string) *api.ActionResp {
	fUIDi64 := int64(fromUID)
	req := &api.ActionReq{
		FromUserId: fUIDi64,
		ToUserId:   int64(toUID),
		ActionType: int64(actionType),
		Content:    content,
	}
	ChatOnlineMap[fUIDi64] = -1
	resp, err := messageClient.Action(context.Background(), req, callopt.WithRPCTimeout(3*time.Second), callopt.WithConnectTimeout(3*time.Second))
	if err != nil {
		msg := "发送消息失败！"
		log.Println(msg + err.Error())
		return &api.ActionResp{
			StatusCode: -1,
			StatusMsg:  &msg,
		}
	}
	ChatOnlineMap[fUIDi64] = time.Now().Unix()
	return resp
}

func ChatLatest(fromUID, toUID int) string {
	req := &api.ChatLatestReq{
		FromUserId: int64(fromUID),
		ToUserId:   int64(toUID),
	}
	resp, err := messageClient.ChatLatest(context.Background(), req, callopt.WithRPCTimeout(3*time.Second), callopt.WithConnectTimeout(3*time.Second))
	if err != nil {
		return ""
	}
	return resp.Message.Content
}
