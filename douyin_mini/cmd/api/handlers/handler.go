package handlers

//这个包里面主要定义了前端发过来请求的一个实体格式
//通用返回格式
//发送响应的通用函数
import (
	"douyin_mini/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type Response struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SendResponse pack response
// 这个是一个通用的返回的一个函数
// 传进来*app.RequestContext和data，这个函数能自动帮你写好json
func SendResponse(c *app.RequestContext, err error, data interface{}) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, Response{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
		Data:    data,
	})
}

// UserParam 这是一个前端传过来的请求数据的实体类
// body格式请求
// 传过来用户名和密码
type UserParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
