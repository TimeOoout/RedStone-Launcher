package main

import (
	"RedStoneLauncher/RSL_Log"
	"RedStoneLauncher/RSL_Setting"
)

func main() {
	RSL_Log.InitLauncherLogger()
	a := RSL_Log.ClearLogs()
	if a != nil {
		println(a.Error())
	}
	RSL_Log.GetVersion()
	RSL_Setting.InitSettings()

}
