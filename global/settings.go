package global

import (
	"catlinie_test01/pkg/logger"
	"catlinie_test01/pkg/setting"
)

var (
	ServerSetting   setting.ServerSettings
	DatabaseSetting setting.DatabaseSettings
	AppSetting      setting.AppSettings
	Logger          *logger.CatLogger
)
