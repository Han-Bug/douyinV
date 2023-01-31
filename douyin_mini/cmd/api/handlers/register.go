package handlers

import (
	"context"
	"douyin_mini/cmd/api/rpc"
	"douyin_mini/kitex_gen/user"
	"douyin_mini/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
)

// Register register user info
// 这也是一个处理http请求的函数
// 作用是注册
func Register(ctx context.Context, c *app.RequestContext) {
	var registerVar UserParam
	if err := c.Bind(&registerVar); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	if len(registerVar.Username) == 0 || len(registerVar.Password) == 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	err := rpc.CreateUser(context.Background(), &user.RegisterUserRequest{
		Username: registerVar.Username,
		Password: registerVar.Password,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}
