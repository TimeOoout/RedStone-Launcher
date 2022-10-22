package RSL_Log

import (
	"fmt"
	"github.com/fatih/color"
)

func PrintInfo(format string, args ...interface{}) {
	fmt.Printf(color.CyanString("[INFO] "+format+"\n", args...))
}

func PrintWarning(format string, args ...interface{}) {
	fmt.Printf(color.YellowString("[Warning] "+format+"\n", args...))
}

func PrintError(format string, args ...interface{}) {
	fmt.Printf(color.RedString("[Error] "+format+"\n", args...))
}

func PrintDebug(format string, args ...interface{}) {
	fmt.Printf(color.GreenString("[Debug] "+format+"\n", args...))
}
