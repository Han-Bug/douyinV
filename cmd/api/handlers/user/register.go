package user

import (
	"context"
	"douyinV/cmd/api/rpc"
	"douyinV/cmd/user/kitex_gen/user"
	"douyinV/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
)

func Register(ctx context.Context, c *app.RequestContext) {
	var registerVar UserParam
	if err := c.Bind(&registerVar); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	// 空否校验
	if len(registerVar.UserName) == 0 || len(registerVar.PassWord) == 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}
	// 微服务请求
	var resp *user.RegisterResponse
	resp, err := rpc.UserRegister(context.Background(), &user.RegisterRequest{
		Username: registerVar.UserName,
		Password: registerVar.PassWord,
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
