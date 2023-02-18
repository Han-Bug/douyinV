package user

import (
	"context"
	"douyinV/api/biz/model/user"
	"douyinV/api/global"
	rpcUser "douyinV/api/kitex_gen/user"
	"douyinV/api/pkg/errno"
	"douyinV/api/pkg/jwt"
	"douyinV/api/rpc"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Register .
// @router /douyin/user/register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var req user.RegisterRequest
	err := c.BindAndValidate(&req)

	if err != nil || len(req.Username) == 0 || len(req.Password) == 0 {
		// 请求数据不全
		c.String(consts.StatusBadRequest, "用户名或密码不全")
		return
	}
	// rpc请求
	resp, err := rpc.UserRegister(context.Background(), &rpcUser.RegisterRequest{
		Username: req.Username,
		Password: req.Password,
	})
	// err: rpc请求失败
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	// err: rpc返回错误码
	if resp.BaseResp.StatusCode != 0 {
		SendResponse(c, errno.NewErrNo(
			resp.BaseResp.StatusCode,
			resp.BaseResp.StatusMessage), nil)
		return
	}
	// 生成token

	// 生成token
	tokenData := jwt.TokenData{
		UserId: resp.Id,
	}
	jwtImpl := global.APIImpl.JWTImpl
	token, err := jwtImpl.MakeToken(tokenData)
	// err: 生成token失败
	if err != nil {
		c.String(consts.StatusBadRequest, "token生成失败")
		return
	}

	// 返回正确结果
	c.JSON(consts.StatusOK, user.RegisterResponse{
		StatusCode: 0,
		StatusMsg:  "",
		UserID:     resp.Id,
		Token:      token,
	})
}

//
//func Register(ctx context.Context, c *app.RequestContext) {
//	var req user.RegisterRequest
//	err := c.BindAndValidate(&req)
//	if err != nil || len(req.Username) == 0 || len(req.Password) == 0 {
//		// 请求数据不全（空数据）
//		c.String(consts.StatusBadRequest, err.Error())
//		return
//	}
//
//	// 微服务请求
//	var resp *user.RegisterResponse
//	resp, err := rpc.UserRegister(context.Background(), &user.RegisterRequest{
//		Username: registerVar.UserName,
//		Password: registerVar.PassWord,
//	})
//	if err != nil {
//		SendResponse(c, errno.ConvertErr(err), nil)
//		return
//	}
//	// TODO 生成token
//	SendResponse(c, errno.Success, &TokenParam{
//		Id:    resp.Id,
//		Token: "token",
//	})
//}
