package user

import (
	"douyinV/api/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

//type Response struct {
//	Code    int64  `json:"status_code"`
//	Message string `json:"status_msg"`
//	UserId  int64  `json:"user_id"`
//	Token   string `json:"token"`
//}
// InfoResponse 用户信息请求的返回结果
//type InfoResponse struct {
//	Code         int64    `json:"status_code"`
//	Message      string   `json:"status_msg"`
//	UserResParam ResParam `json:"user"`
//}
//
//type ResParam struct {
//	Id            int64  `json:"id"`
//	Name          string `json:"name"`
//	FollowCount   int32  `json:"follow_count"`
//	FollowerCount int32  `json:"follower_count"`
//	IsFollow      bool   `json:"is_follow"`
//}
//type TokenParam struct {
//	Id    int64  `json:"user_id"`
//	Token string `json:"token"`
//}
//
//type LoginRequest struct {
//	UserName string `json:"username"`
//	PassWord string `json:"password"`
//}
//
//type RegisterRequest struct {
//}
//type InfoRequest struct {
//}

type BaseResponse struct {
	Code    int64
	Message string
	Data    interface{}
}

// SendResponse 返回结果
func SendResponse(c *app.RequestContext, err errno.ErrNo, data interface{}) {
	c.JSON(consts.StatusOK, BaseResponse{
		Code:    err.Code,
		Message: err.Message,
		Data:    data,
	})
}

//func SendInfoResponse(c *app.RequestContext, err error, resParam *ResParam) {
//	Err := errno.ConvertErr(err)
//	if resParam == nil {
//		c.JSON(consts.StatusOK, InfoResponse{
//			Code:         Err.Code,
//			Message:      Err.Message,
//			UserResParam: ResParam{},
//		})
//	} else {
//		c.JSON(consts.StatusOK, InfoResponse{
//			Code:         Err.Code,
//			Message:      Err.Message,
//			UserResParam: *resParam,
//		})
//	}
//}
