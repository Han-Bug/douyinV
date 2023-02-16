package main

import (
	"context"
	"errors"
	"tiktok/dal/mysql"
	"tiktok/message/kitex_gen/api"
)

// MessageImpl implements the last service interface defined in the IDL.
type MessageImpl struct{}

// Chat implements the MessageImpl interface.
func (s *MessageImpl) Chat(ctx context.Context, req *api.ChatReq) (resp *api.ChatResp, err error) {
	msgs, err := mysql.GetMessages(int(req.FromUserId), int(req.ToUserId))
	if err != nil {
		return nil, err
	}

	var msgList []*api.Msg
	for _, v := range msgs {
		if v.CreatedAt.Unix() <= req.LatestTime {
			continue
		}
		msgList = append(msgList, &api.Msg{
			ToUserId:   int64(v.ToUserId),
			FromUserId: int64(v.FromUserId),
			Content:    v.Content,
			CreateTime: v.CreatedAt.Unix(),
		})
	}

	resp = &api.ChatResp{
		StatusCode:  0,
		StatusMsg:   nil,
		MessageList: msgList,
	}
	return
}

// Action implements the MessageImpl interface.
func (s *MessageImpl) Action(ctx context.Context, req *api.ActionReq) (resp *api.ActionResp, err error) {
	if req.ActionType != 1 {
		return nil, errors.New("不支持该操作类型")
	}

	err = mysql.NewMessage(int(req.FromUserId), int(req.ToUserId), req.Content)
	if err != nil {
		return
	}

	resp = &api.ActionResp{
		StatusCode: 0,
	}
	return
}
