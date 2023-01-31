package handlers

import (
	"context"
	"douyin_mini/cmd/api/rpc"
	"douyin_mini/kitex_gen/user"
	"douyin_mini/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
)

// GetUserInfo 获取用户信息
func GetUserInfo(ctx context.Context, c *app.RequestContext) {
	ToUserIdStr := c.Param("user_id")
	ToUserId, err := strconv.ParseInt(ToUserIdStr, 10, 64)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	var UserIds []int64
	UserIds[0] = ToUserId
	rpc.GetUserInfo(ctx, &user.UserInfoRequest{UserIds: UserIds})
}
