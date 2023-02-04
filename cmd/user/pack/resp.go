package pack

import (
	"douyinV/cmd/user/kitex_gen/user"
	"douyinV/pkg/errno"
	"errors"
	"time"
)

// BuildBaseResp build baseResp from error
// 将error转换成BaseResp
func BuildBaseResp(err error) *user.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}
	// 如果err的类型为自定义错误ErrNo，直接调用函数转换为baseResp
	if errors.As(err, &e) {
		return baseResp(e)
	}
	// 如果err类型不是ErrNo，根据err的信息生成ErrNo并转换为baseResp
	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

// 将ErrorNo转换成BaseResp
func baseResp(err errno.ErrNo) *user.BaseResp {
	return &user.BaseResp{StatusCode: err.Code, StatusMessage: err.Message, ServiceTime: time.Now().Unix()}
}
