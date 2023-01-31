package pack

import (
	"douyin_mini/kitex_gen/user"
	"douyin_mini/pkg/errno"
	"errors"
	"time"
)

func BuildBaseResp(err error) *user.BaseResp {
	if err != nil {
		return baseResp(errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}
	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)

}

// 将errno转换成BaseResp
func baseResp(err errno.ErrNo) *user.BaseResp {
	return &user.BaseResp{
		StatusCode:    err.ErrCode,
		StatusMessage: err.ErrMsg,
		ServiceTime:   time.Now().Unix(),
	}
}
