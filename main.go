package main

import (
	"catlinie_test01/global"
	"catlinie_test01/internal/dao"
	"catlinie_test01/internal/routers"
	setting "catlinie_test01/pkg/setting"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func init() {
	err := setupSettings()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
}

//进行各种准备工作并启动主路由
// @title CatForest API
// @version 1.0
// @description 猫猫森林第一版
// @host localhost:8080
// @BasePath /api/v1
func main() {
	engine := routers.NewRouter()
	engine.Run(global.ServerSetting.HttpPort)
}

func setupSettings() error {
	s, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = s.ReadConfig("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = s.ReadConfig("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupDBEngine() error {
	db, err := dao.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	global.DBEngine = db
	return nil
}
