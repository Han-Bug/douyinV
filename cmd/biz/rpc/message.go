package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"log"
	"tiktok/config"
	"tiktok/message/kitex_gen/api"
	"tiktok/message/kitex_gen/api/message"
	"time"
)

var messageClient message.Client

func initMessage() {
	mc, err := message.NewClient("message", client.WithHostPorts(config.ServerMsgUrlHP))
	if err != nil {
		panic("连接message服务失败！" + err.Error())
	}
	messageClient = mc
}

func Chat(fromUID, toUID int) *api.ChatResp {
	req := &api.ChatReq{
		FromUserId: int64(fromUID),
		ToUserId:   int64(toUID),
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
	return resp
}

func Action(fromUID, toUID, actionType int, content string) *api.ActionResp {
	req := &api.ActionReq{
		FromUserId: int64(fromUID),
		ToUserId:   int64(toUID),
		ActionType: int64(actionType),
		Content:    content,
	}
	resp, err := messageClient.Action(context.Background(), req, callopt.WithRPCTimeout(3*time.Second), callopt.WithConnectTimeout(3*time.Second))
	if err != nil {
		msg := "发送消息失败！"
		log.Println(msg + err.Error())
		return &api.ActionResp{
			StatusCode: -1,
			StatusMsg:  &msg,
		}
	}
	return resp
}
