package RSL_Log

import (
	"go.uber.org/zap"
	"os"
)

var LogFolder = "./LauncherLog/"
var LauncherLogger *zap.Logger
var LogFile *os.File
var AutoLogging = true
var AutoMsg = true
var LogtoConsole = true
var inited = false
var version = "1.0.0_beta"
var LogObj = "Log    : "
