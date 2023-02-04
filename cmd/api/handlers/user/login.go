package user

import (
	"context"
	"douyinV/cmd/api/rpc"
	"douyinV/cmd/user/kitex_gen/user"
	"douyinV/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
)

func Login(ctx context.Context, c *app.RequestContext) {
	var loginVar UserParam
	if err := c.Bind(&loginVar); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	// 空否校验
	if len(loginVar.UserName) == 0 || len(loginVar.PassWord) == 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}
	// 微服务请求
	resp, err := rpc.UserLogin(context.Background(), &user.LoginRequest{
		Username: loginVar.UserName,
		Password: loginVar.PassWord,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	// TODO 生成token
	SendResponse(c, errno.Success, &TokenParam{
		Id:    resp.Id,
		Token: "token",
	})
}
