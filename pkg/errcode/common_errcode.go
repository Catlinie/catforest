package errcode

var (
	// Success 200
	Success = NewError(10000001, "鉴权失败，找不到对应的AppKey和AppSecret")
	// InvalidParams 400
	InvalidParams = NewError(10000002, "鉴权失败，找不到对应的AppKey和AppSecret")
	// NotFound 404
	NotFound = NewError(10000003, "鉴权失败，找不到对应的AppKey和AppSecret")
	// LowAuthorize 403
	LowAuthorize = NewError(10000004, "鉴权失败，找不到对应的AppKey和AppSecret")
	// TooManyRequest 429
	TooManyRequest = NewError(10000005, "鉴权失败，找不到对应的AppKey和AppSecret")
	// ServerError 500
	ServerError               = NewError(10000006, "服务器内部错误")
	UnauthorizedAuthNotExist  = NewError(10000007, "鉴权失败，找不到对应的AppKey和AppSecret")
	UnauthorizedTokenError    = NewError(10000008, "鉴权失败，Token错误")
	UnauthorizedTokenTimeout  = NewError(10000009, "鉴权失败，Token超时")
	UnauthorizedTokenGenerate = NewError(10000010, "鉴权失败，Token生成失败")
)
