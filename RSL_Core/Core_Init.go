package RSL_Core

import (
	"github.com/TimeOoout/RSL_Core/RSL_Log"
	"github.com/TimeOoout/RSL_Core/RSL_Setting"
)

func InitLauncher() InitErrors {
	//返回的错误集合
	errors := InitErrors{
		InitLogger:    "",
		InitSetting:   "",
		InitClearLogs: "",
	}

	//打印日志包的版本
	RSL_Log.GetVersion()
	//初始化日志记录器
	initLogger(&errors)
	//初始信息
	logInfoWelcome()
	//初始化设置信息
	initSettings(&errors)

	return errors
}

func initLogger(e *InitErrors) {
	//初始化日志记录器
	err := RSL_Log.InitLauncherLogger()
	if err != nil {
		//记录错误
		e.InitLogger = err.Error()
	}
	RSL_Log.LogVersion()
}

func initSettings(e *InitErrors) {
	//初始化设置
	err := RSL_Setting.InitSettings()
	if err != nil {
		//记录错误
		err = RSL_Log.ClearLogs(RSL_Setting.CurrentConfig.LauncherSettings.Log.RemainLogs)
		//默认自动清理日志
		if err != nil {
			//记录错误
			e.InitClearLogs = err.Error()
		}
	} else {
		//检查是否自动清理日志
		if RSL_Setting.CurrentConfig.LauncherSettings.Log.AutoClearLog == true {
			err = RSL_Log.ClearLogs(RSL_Setting.CurrentConfig.LauncherSettings.Log.RemainLogs)
			if err != nil {
				//记录错误
				e.InitClearLogs = err.Error()
			}
		}
	}
}
