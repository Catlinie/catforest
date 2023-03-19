package service

import (
	"catlinie_test01/internal/model"
	"time"
)

type GetLoginInfoRequest struct {
	Account string `form:"account" binding:"required"`
	//UnconfirmedPwd string
	Itype string `form:"itype" binding:"required"`
}

type CreateUserInfoRequest struct {
	Username   string `form:"username" binding:"required"`
	Desc       string `form:"desc" `
	MasterName string `form:"master_name"`
	Url        string `form:"url"`
	Gender     uint8  `form:"gender" binding:"oneof=0 1"`
	Account    string `form:"account" binding:"required"`
	Password   string `form:"password" binding:"required"`
	Itype      string `form:"itype" binding:"required,oneof=desktop web phone"`
}

type GetEntireInfoRequest struct {
	Account string `form:"account" binding:"required"`
	Itype   string `form:"itype" binding:"required,oneof=desktop web phone"`
}

type SetUserInfoRequest struct {
	Username   string `form:"username"`
	Desc       string `form:"desc"`
	MasterName string `form:"master_name"`
	Url        string `form:"url"`
	Gender     uint8  `form:"gender"`
	Password   string `form:"password"`
	Account    string `form:"account" binding:"required"`
	Itype      string `form:"itype" binding:"required,oneof=desktop web phone"`
}

func (s *Service) GetLoginInfo(param *GetLoginInfoRequest) (string, error) {
	return s.Engine.GetLoginInfo(param.Account, param.Itype)
}

func (s *Service) CreateUserInfo(param *CreateUserInfoRequest) error {
	return s.Engine.CreateUserInfo(
		&model.User{
			Cid:          100000,
			Username:     param.Username,
			Desc:         param.Desc,
			MasterName:   param.MasterName,
			Url:          param.Url,
			Gender:       param.Gender,
			Times:        0,
			Bills:        0,
			RegisterTime: time.Now().Unix(),
			LoginTime:    0,
			IsBaned:      0,
		}, param.Account, param.Password, param.Itype,
	)
}

func (s *Service) GetEntireInfo(param *GetEntireInfoRequest) (*model.User, error) {
	return s.Engine.GetEntireInfo(param.Account, param.Itype)
}

func (s *Service) SetUserInfo(param *SetUserInfoRequest) error {
	return s.Engine.SetUserInfo(
		param.Username,
		param.Desc,
		param.MasterName,
		param.Url,
		param.Gender,
	)
}
