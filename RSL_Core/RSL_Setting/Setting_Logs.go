package RSL_Setting

import "RedStoneLauncher/RSL_Core/RSL_Log"

// name: log-Info-InitLogger
func logExample() {
	//Statements.
}

func logInfoInitSetting() {
	RSL_Log.LogInfo(LogObj + "Successfully init settings!")
}

func logInfoGetSetting() {
	RSL_Log.LogInfo(LogObj + "Successfully get settings!")
}

func logWarningInitSettings(s string) {
	RSL_Log.LogInfo(LogObj + "Failed to init settings!   Error: " + s)
}

func logWarningGetSettings(s string) {
	RSL_Log.LogInfo(LogObj + "Failed to get settings!   Error: " + s)
}

func logWarningSetSettings(s string) {
	RSL_Log.LogInfo(LogObj + "Failed to set settings!   Error: " + s)
}
