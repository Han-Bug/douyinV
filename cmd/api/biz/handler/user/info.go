package user

import (
	"context"
	"douyinV/api/biz/model/user"
	"douyinV/api/global"
	rpcUser "douyinV/api/kitex_gen/user"
	"douyinV/api/pkg/errno"
	"douyinV/api/rpc"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"strconv"
)

// TODO 将c.String形式的错误输出部分转换为统一的错误码

// Info .
// @router /douyin/user/login [POST]
func Info(ctx context.Context, c *app.RequestContext) {
	var req user.InfoRequest
	// err: 请求数据格式有误
	if err := c.BindAndValidate(&req); err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	// 参数校验

	// 空否校验
	if len(req.UserID) == 0 {
		c.String(consts.StatusBadRequest, "无效的用户ID")
		return
	}

	toUserId, err := strconv.Atoi(req.UserID)
	// err: toUserId无法解析成int
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	// token校验
	var userId int64
	userId = 0

	if len(req.Token) != 0 {
		jwt := global.APIImpl.JWTImpl

		// 解析token
		tokenData, err := jwt.ParseToken(req.Token)
		// err: 无效的token
		// token无效时采用默认的匿名查看用户信息
		if err == nil {
			userId = tokenData.UserId

		}
	}

	// rpc请求
	resp, err := rpc.UserInfo(ctx, &rpcUser.UserInfoRequest{
		Id:   userId,
		ToId: int64(toUserId),
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

	// 返回成功结果
	c.JSON(consts.StatusOK, user.InfoResponse{
		StatusCode: 0,
		StatusMsg:  "",
		User: &user.UserParam{
			ID:            resp.Id,
			Name:          resp.Name,
			FollowCount:   resp.FollowCount,
			FollowerCount: resp.FollowerCount,
			IsFollow:      resp.IsFollow,
		},
	})
}
