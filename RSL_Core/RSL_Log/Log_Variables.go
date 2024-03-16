package RSL_Log

import (
	"go.uber.org/zap"
	"os"
)

var LogFolder = "./LauncherLog/"
var LauncherLogger *zap.Logger
var LogFile *os.File
var LogtoConsole = true
var inited = false
var version = "v1.0.0_beta"
var LogObj = "Log    : "
