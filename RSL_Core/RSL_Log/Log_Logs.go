package RSL_Log

import "strconv"

// name: log-Info-InitLogger
func logExample() {
	//Statements.
}

func logInfoInitLogger() {
	LogInfo(LogObj + "Init Logger successfully!")
}

func logWarningClearLog(err error) {
	LogWarning(LogObj + "Failed to clear all the logs.   Error: " + err.Error())
}

func logWarningClearLogArgs(arg int) {
	LogWarning(LogObj + "Wrong args entered!   Entered Args: " + strconv.Itoa(arg))
}

func logInfoClearLog() {
	LogInfo(LogObj + "Logs cleared successfully!")
}
