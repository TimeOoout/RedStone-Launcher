package RSL_Core

import (
	"github.com/TimeOoout/RSL_Core/RSL_Log"
)

func logInfoWelcome() {
	RSL_Log.LogInfo(LogObj + "Welcome to RSL_ Core!")
	RSL_Log.LogInfo(LogObj+"RSL_Core version:	%s", Version)
}
