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

// Login .
// @router /douyin/user [GET]
func Login(ctx context.Context, c *app.RequestContext) {
	var req user.LoginRequest
	// err: 请求数据格式有误
	if err := c.BindAndValidate(&req); err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	// err: 空否校验
	if len(req.Username) == 0 || len(req.Password) == 0 {
		c.String(consts.StatusBadRequest, "帐号或密码未填写完整")
		return
	}

	// rpc请求
	resp, err := rpc.UserLogin(ctx, &rpcUser.LoginRequest{
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

	// 返回成功结果
	c.JSON(consts.StatusOK, user.LoginResponse{
		StatusCode: 0,
		StatusMsg:  "",
		UserID:     resp.Id,
		Token:      token,
	})
}
