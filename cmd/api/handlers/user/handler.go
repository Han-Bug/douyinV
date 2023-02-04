package user

//这个包里面主要定义了前端发过来请求的一个实体格式
//通用返回格式
//发送响应的通用函数
import (
	"douyinV/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type Response struct {
	Code    int64  `json:"status_code"`
	Message string `json:"status_msg"`
	UserId  int64  `json:"user_id"`
	Token   string `json:"token"`
}
type InfoResponse struct {
	Code         int64    `json:"status_code"`
	Message      string   `json:"status_msg"`
	UserResParam ResParam `json:"user"`
}
type ResParam struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int32  `json:"follow_count"`
	FollowerCount int32  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}
type UserParam struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

type TokenParam struct {
	Id    int64  `json:"user_id"`
	Token string `json:"token"`
}

// SendResponse pack response
// 这个是一个通用的返回的一个函数
// 传进来*app.RequestContext和data，这个函数能自动帮你写好json
func SendResponse(c *app.RequestContext, err error, tokenParam *TokenParam) {
	Err := errno.ConvertErr(err)
	if tokenParam == nil {
		c.JSON(consts.StatusOK, Response{
			Code:    Err.Code,
			Message: Err.Message,
			UserId:  0,
			Token:   "",
		})
	} else {
		c.JSON(consts.StatusOK, Response{
			Code:    Err.Code,
			Message: Err.Message,
			UserId:  tokenParam.Id,
			Token:   tokenParam.Token,
		})
	}

}
func SendInfoResponse(c *app.RequestContext, err error, resParam *ResParam) {
	Err := errno.ConvertErr(err)
	if resParam == nil {
		c.JSON(consts.StatusOK, InfoResponse{
			Code:         Err.Code,
			Message:      Err.Message,
			UserResParam: ResParam{},
		})
	} else {
		c.JSON(consts.StatusOK, InfoResponse{
			Code:         Err.Code,
			Message:      Err.Message,
			UserResParam: *resParam,
		})
	}
}
