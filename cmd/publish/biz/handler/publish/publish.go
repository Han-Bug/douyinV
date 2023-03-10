// Code generated by hertz generator.

package publish

import (
	"context"
	publish "douyinV/cmd/publish/biz/model/publish"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Action .
// @router douyin/publish/action [POST]
func Action(ctx context.Context, c *app.RequestContext) {
	var err error
	var req publish.ActionRequest
	err = c.BindAndValidate(&req)
	// err: 请求参数未满足
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	// 将数据存至minio并获取链接

	resp := new(publish.ActionResponse)

	c.JSON(consts.StatusOK, resp)
}

// List .
// @router douyin/publish/list [GET]
func List(ctx context.Context, c *app.RequestContext) {
	var err error
	var req publish.ListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(publish.ListResponse)

	c.JSON(consts.StatusOK, resp)
}
