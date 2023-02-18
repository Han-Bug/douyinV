package errno

import (
	"errors"
	"fmt"
)

type ErrNo struct {
	Code    int64
	Message string
}

func NewErrNo(code int64, msg string) ErrNo {
	return ErrNo{code, msg}
}
func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.Code, e.Message)
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.Message = msg
	return e
}

// ConvertErr 将error转换为ErrNo
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	// 如果已经是ErrNo了
	if errors.As(err, &Err) {
		return Err
	}
	s := ServiceErr
	s.Message = err.Error()
	return s
}
