package errcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	code    int      `json:"code"`
	msg     string   `json:"msg"`
	details []string `json:"details"`
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}
	codes[code] = msg
	return &Error{code: code, msg: msg}
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.details = []string{}
	for _, d := range details {
		newError.details = append(newError.details, d)
	}
	return &newError
}

func (e *Error) ToString() string {
	return fmt.Sprintf("{errcode:%d, msg:%s}", e.Code(), e.Msg())
}

//func (e Error) ToStringf()  {
//
//}

func (e *Error) StatusCode() int {
	switch e.code {
	case Success.Code():
		return http.StatusOK
	case InvalidParams.Code():
		return http.StatusBadRequest
	case NotFound.Code():
		return http.StatusNotFound
	case LowAuthorize.Code():
		return http.StatusForbidden
	case TooManyRequest.Code():
		return http.StatusTooManyRequests
	case ServerError.Code():
		fallthrough
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
	case UnauthorizedTokenGenerate.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		return http.StatusUnauthorized
	}
	return http.StatusInternalServerError
}
