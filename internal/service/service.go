package service

import (
	"catlinie_test01/global"
	"catlinie_test01/internal/dao"
	"context"
	"github.com/gin-gonic/gin"
)

type Service struct {
	Ctx    context.Context
	Engine *dao.Dao
}

func NewService(c *gin.Context) *Service {
	return &Service{
		Ctx:    c,
		Engine: dao.New(global.DBEngine),
	}
}
