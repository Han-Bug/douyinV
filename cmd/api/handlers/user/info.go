package user

import (
	"context"
	"douyinV/cmd/api/rpc"
	"douyinV/cmd/user/kitex_gen/user"
	"douyinV/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
)

func GetInfo(ctx context.Context, c *app.RequestContext) {
	var tokenVar TokenParam
	if err := c.Bind(&tokenVar); err != nil {
		SendInfoResponse(c, errno.ConvertErr(err), nil)
		return
	}
	// 根据token有无而给予不同结果
	if len(tokenVar.Token) == 0 {
		SendInfoResponse(c, errno.ParamErr, nil)
		return
	}
	// TODO 解析token

	// TODO 从token获取id

	// 微服务请求
	resp, err := rpc.UserInfo(context.Background(), &user.UserInfoRequest{
		Id:   0,
		ToId: tokenVar.Id,
	})
	if err != nil {
		SendInfoResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendInfoResponse(c, errno.Success, &ResParam{
		Id:            resp.Id,
		Name:          resp.Name,
		FollowCount:   resp.FollowCount,
		FollowerCount: resp.FollowerCount,
		IsFollow:      resp.IsFollow,
	})
}
