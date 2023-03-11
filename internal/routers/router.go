package routers

import (
	_ "catlinie_test01/docs"
	v1 "catlinie_test01/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// NewRouter 主要路由注册
func NewRouter() *gin.Engine {
	engine := gin.Default()
	root := engine.Group("/")

	userAPI := v1.NewUser()
	tasksAPI := v1.NewTasks()
	{
		//login
		root.GET("/user", userAPI.Login)
		//register
		root.POST("/user", userAPI.Register)
		root.DELETE("/user", userAPI.Logout)
	}

	menu := root.Group("/menu")
	{
		menu.POST("/timer", tasksAPI.UploadTimerData)
		menu.GET("/info", userAPI.GetInfo)
		menu.POST("/info", userAPI.SetInfo)
	}

	statistics := menu.Group("/statistics")
	{
		statistics.GET("/tpop", tasksAPI.GetTPOPData)
		// 已知当天的日与星期数
		// for：（星期数+7）%7
		statistics.GET("/tpoc", tasksAPI.GetTPOCData)
		statistics.GET("/tdow", tasksAPI.GetTDOWData)
		statistics.GET("/money", tasksAPI.GetMoneyData)
	}

	root.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return engine
}
