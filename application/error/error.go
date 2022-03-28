package error

import (
	jsoniter "github.com/json-iterator/go"
)

type ErrorCode int64

type Error struct {
	Code    ErrorCode
	Message string
	Reason  string
}

func (e *Error) Error() string {
	err, _ := jsoniter.Marshal(e)
	return string(err)
}

func (e *Error) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"code":    e.Code,
		"message": e.Message,
	}
}

func NewError(errCode ErrorCode, err error) *Error {
	return &Error{
		Code:    errCode,
		Message: GetErrorMsg(errCode),
		Reason:  err.Error(),
	}
}
