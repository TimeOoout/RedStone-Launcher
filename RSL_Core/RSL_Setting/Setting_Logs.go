package RSL_Setting

import "github.com/TimeOoout/RSL_Core/RSL_Log"

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
	RSL_Log.LogInfo(LogObj + "Failed to init settings! The default settings will be used.    Error: " + s)
}

func logWarningGetSettings(s string) {
	RSL_Log.LogInfo(LogObj + "Failed to get settings! The current settings will continue to be used.   Error: " + s)
}

func logWarningSetSettings(s string) {
	RSL_Log.LogInfo(LogObj + "Failed to set settings! All changes will not be saved.   Error: " + s)
}
