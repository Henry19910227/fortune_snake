package err

import (
	"errors"
	"fmt"
	"game_server_slots_fortune_snake/constants"
)

type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"-"` // 原始錯誤，不回傳給用戶
}

func New(code int32, msg string, err error) *Error {
	return &Error{
		Code:    code,
		Message: msg,
		Err:     err,
	}
}

func (e *Error) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("code=%d, msg=%s, err=%v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("code=%d, msg=%s", e.Code, e.Message)
}

func Err(obj error) Error {
	var e *Error
	if ok := errors.As(obj, &e); !ok {
		return Error{
			Code:    constants.CodeInternalError,
			Message: e.Error(),
			Err:     obj,
		}
	}
	return *e
}
