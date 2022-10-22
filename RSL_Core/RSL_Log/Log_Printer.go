package RSL_Log

import (
	"github.com/fatih/color"
)

func PrintInfo(format string, args ...interface{}) {
	color.Cyan("[INFO] "+format+"\n", args...)
}

func PrintWarning(format string, args ...interface{}) {
	color.Yellow("[Warning] "+format+"\n", args...)
}

func PrintError(format string, args ...interface{}) {
	color.Red("[Error] "+format+"\n", args...)
}

func PrintDebug(format string, args ...interface{}) {
	color.Green("[Debug] "+format+"\n", args...)
}
