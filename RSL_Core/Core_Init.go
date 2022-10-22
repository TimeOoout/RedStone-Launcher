package RSL_Core

import (
	"RedStoneLauncher/RSL_Core/RSL_Log"
	"RedStoneLauncher/RSL_Core/RSL_Setting"
)

func InitLauncher() {
	RSL_Log.InitLauncherLogger()
	a := RSL_Log.ClearLogs()
	if a != nil {
		println(a.Error())
	}
	RSL_Log.GetVersion()
	RSL_Setting.InitSettings()
}
