package RSL_Player

import "github.com/TimeOoout/RSL_Core/RSL_Log"

func logWarningReadDir(err error) {
	RSL_Log.LogWarning(LogObj + "Unable to get user list: Failed to read dir.   Error: " + err.Error())
}
