package main

import (
	"catlinie_test01/global"
	"catlinie_test01/internal/dao"
	"catlinie_test01/internal/routers"
	"catlinie_test01/pkg/logger"
	setting "catlinie_test01/pkg/setting"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"os"
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
	setupLogger()
	global.Logger.Info("init success")
}

// 进行各种准备工作并启动主路由
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
	err = s.ReadConfig("App", &global.AppSetting)
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

func setupLogger() {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	fmt.Println(fileName)
	global.Logger = logger.NewLogger(logger.NewCore(&logger.LoggerConfig{
		EncoderType:       global.AppSetting.EncoderType,
		EncoderConfigType: global.AppSetting.EncoderConfigType,
		Writers: []io.Writer{
			os.Stdout,
			&lumberjack.Logger{
				Filename:   fileName,
				MaxSize:    global.AppSetting.MaxSize,
				MaxAge:     global.AppSetting.MaxAge,
				MaxBackups: global.AppSetting.MaxBackups,
				Compress:   global.AppSetting.Compress,
			},
		},
	}))
}
