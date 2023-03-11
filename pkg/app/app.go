package app

import (
	"catlinie_test01/pkg/errcode"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Ctx *gin.Context
}

func NewResponse(c *gin.Context) *Response {
	return &Response{Ctx: c}
}

func (r *Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, data)
}

func (r *Response) ToErrorResponse(error *errcode.Error) {
	response := gin.H{
		"code": error.Code(),
		"msg":  error.Msg(),
	}
	details := error.Details()
	if len(details) > 0 {
		response["details"] = details
	}
	r.Ctx.JSON(error.StatusCode(), response)
}
