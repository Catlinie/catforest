package v1

import (
	"catlinie_test01/internal/service"
	"catlinie_test01/pkg/app"
	"catlinie_test01/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type User struct{}

func NewUser() User {
	return User{}
}

// Login
// @Summary 用户登陆
// @Produce json
// @Param account query string true "账号"
// @Param password query string true "密码"
// @Success 200 {object} app.Response "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router  /user [get]
func (u User) Login(c *gin.Context) {
	svc := service.NewService(c)
	rep := app.NewResponse(c)
	pwd, err := svc.GetLoginInfo(service.GetLoginInfoRequest{Account: c.Query("account"), Itype: c.Query("itype")})
	if err != nil {
		rep.ToErrorResponse(errcode.ServerError)
	}
	c.JSON(200, gin.H{
		"pwd": pwd,
	})

}

// Register
// @Summary 用户注册
// @Produce json
// @Param username body string true "用户昵称"
// @Param account body string true "账号"
// @Param gender body bool false "性别"
// @Param desc body string false "誓言"
// @Param password body string true "密码"
// @Param itype body string false "平台名称"
// @Param master_name body string false "主子名字"
// @Param url body string false "头像地址"
// @Success 200 {object} app.Response "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router  /user [post]
func (u User) Register(c *gin.Context) {
	svc := service.NewService(c)
	response := app.NewResponse(c)
	var request service.CreateUserInfoRequest
	if c.ShouldBind(&request) == nil {
		err := svc.CreateUserInfo(request)
		if err != nil {
			errReq := errcode.ServerError.WithDetails(err.Error())
			response.ToErrorResponse(errReq)
			return
		}
	} else {
		errReq := errcode.ServerError.WithDetails("Register.ShouldBind err")
		response.ToErrorResponse(errReq)
	}
}

// Logout
// @Summary 用户登出
// @Produce json
// @Success 200 {object} app.Response "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router  /user [delete]
func (u User) Logout(c *gin.Context) {

}

// GetInfo
// @Summary 获取用户信息
// @Produce json
// @Param account path string true "账号"
// @Param itype path string true "平台名称"
// @Success 200 {object} app.Response "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router  /menu/info [get]
func (u User) GetInfo(c *gin.Context) {
	svc := service.NewService(c)
	response := app.NewResponse(c)
	var request service.GetEntireInfoRequest
	if c.ShouldBind(&request) == nil {
		info, err := svc.GetEntireInfo(request)
		if err != nil {
			code := errcode.ServerError.WithDetails(err.Error())
			response.ToErrorResponse(code)
			return
		}
		response.ToResponse(info)
		return
	}
	code := errcode.ServerError.WithDetails("GetInfo.ShouldBind err")
	response.ToErrorResponse(code)

}

// SetInfo
// @Summary 修改用户信息
// @Produce json
// @Param username body string true "用户昵称"
// @Param gender body bool false "性别"
// @Param desc body string false "誓言"
// @Param master_name body string false "主子名字"
// @Param url body string false "头像地址"
// @Success 200 {object} app.Response "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router  /menu/info [post]
func (u User) SetInfo(c *gin.Context) {
	svc := service.NewService(c)
	response := app.NewResponse(c)
	request := service.SetUserInfoRequest{}
	if c.ShouldBind(&request) == nil {
		err := svc.SetUserInfo(request)
		if err != nil {
			code := errcode.ServerError.WithDetails(err.Error())
			response.ToErrorResponse(code)
			return
		}
		response.ToResponse("success")
		return
	}
	code := errcode.ServerError.WithDetails("SetInfo.ShouldBind err")
	response.ToErrorResponse(code)
}
