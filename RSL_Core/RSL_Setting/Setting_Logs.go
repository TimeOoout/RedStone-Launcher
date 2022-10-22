package RSL_Setting

import "RedStoneLauncher/RSL_Core/RSL_Log"

// name: log-Info-InitLogger
func logExample() {
	if RSL_Log.AutoLogging == true {
		//Statements.
	}
}

func logInfoInitSetting() {
	if RSL_Log.AutoLogging == true {
		RSL_Log.LogInfo(LogObj + "Successfully init settings!")
	}
}

func logInfoGetSetting() {
	if RSL_Log.AutoLogging == true {
		RSL_Log.LogInfo(LogObj + "Successfully get settings!")
	}
}

func logWarningInitSettings(s string) {
	if RSL_Log.AutoLogging == true {
		RSL_Log.LogInfo(LogObj + "Failed to init settings!   Error: " + s)
	}
}

func logWarningGetSettings(s string) {
	if RSL_Log.AutoLogging == true {
		RSL_Log.LogInfo(LogObj + "Failed to get settings!   Error: " + s)
	}
}

func logWarningSetSettings(s string) {
	if RSL_Log.AutoLogging == true {
		RSL_Log.LogInfo(LogObj + "Failed to set settings!   Error: " + s)
	}
}
