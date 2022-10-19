package RSL_Log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func InitLauncherLogger() {
	initLoggerError := initLogger()
	if initLoggerError != nil {
		msgInfoInitLogger(initLoggerError.Error())
	} else {
		inited = true
		logInfoInitLogger()
	}
}

func ClearLogs() error {
	// 读取当前目录中的所有文件和子目录
	files, err := ioutil.ReadDir(LogFolder)
	if err != nil {
		return err
	}
	// 获取文件，并输出它们的名字
	length := len(files)
	current := 0
	for _, file := range files {
		if length > 3 {
			if file.Name()[len(file.Name())-4:len(file.Name())] == ".log" {
				current++
				if current <= length-3 {
					err = os.Remove(LogFolder + file.Name())
					if err != nil {
						return err
					}
				}
			}
		}

	}

	return nil
}

func GetVersion() {
	LogInfo("RSL_Log version:%s", version)
}

func initLogger() error {
	if _, iSErrA := os.Stat(LogFolder); iSErrA != nil {
		if os.IsNotExist(iSErrA) {
			iSErrB := os.Mkdir(LogFolder, 0666)
			if iSErrB != nil {
				return iSErrB
			}
		}
	}
	var iSErrC error
	LogFile, iSErrC = os.Create(LogFolder + strings.Replace(
		strings.Replace(time.Now().Format("2006-01-02 15:04:05")+".log",
			" ", "_", -1),
		":", "-", -1))
	if iSErrC != nil {
		return iSErrC
	}
	return nil
}

func LogInfo(format string, args ...interface{}) {
	if inited == false {
		PrintError("Logger not initialized!")
	} else {
		encode := zap.NewProductionEncoderConfig()
		encode.EncodeTime = zapcore.ISO8601TimeEncoder
		encode.EncodeLevel = zapcore.CapitalLevelEncoder
		LauncherLogger = zap.New(zapcore.NewCore(zapcore.NewJSONEncoder(encode),
			zapcore.AddSync(LogFile),
			zapcore.InfoLevel), zap.AddCaller())

		sugarLogger := LauncherLogger.Sugar()
		if args != nil {
			sugarLogger.Infof(format, args)
		} else {
			sugarLogger.Infof(format)
		}
		if LogtoConsole == true {
			if args != nil {
				PrintInfo(strings.Replace(
					strings.Replace(time.Now().Format("2006-01-02 15:04:05")+" | ",
						" ", "_", 1),
					":", "-", -1)+format, args)
			} else {
				PrintInfo(strings.Replace(
					strings.Replace(time.Now().Format("2006-01-02 15:04:05")+" | ",
						" ", "_", 1),
					":", "-", -1) + format)
			}
		}
	}
}

func LogWarning(format string, args ...interface{}) {
	if inited == false {
		PrintError("Logger not initialized!")
	} else {
		encode := zap.NewProductionEncoderConfig()
		encode.EncodeTime = zapcore.ISO8601TimeEncoder
		encode.EncodeLevel = zapcore.CapitalLevelEncoder
		LauncherLogger = zap.New(zapcore.NewCore(zapcore.NewJSONEncoder(encode),
			zapcore.AddSync(LogFile),
			zapcore.WarnLevel), zap.AddCaller())

		sugarLogger := LauncherLogger.Sugar()
		if args != nil {
			sugarLogger.Warnf(format, args)
		} else {
			sugarLogger.Warnf(format)
		}
		if LogtoConsole == true {
			if args != nil {
				PrintWarning(strings.Replace(
					strings.Replace(time.Now().Format("2006-01-02 15:04:05")+" | ",
						" ", "_", 1),
					":", "-", -1)+format, args)
			} else {
				PrintWarning(strings.Replace(
					strings.Replace(time.Now().Format("2006-01-02 15:04:05")+" | ",
						" ", "_", 1),
					":", "-", -1) + format)
			}
		}
	}
}

func LogError(format string, args ...interface{}) {
	if inited == false {
		PrintError("Logger not initialized!")
	} else {
		encode := zap.NewProductionEncoderConfig()
		encode.EncodeTime = zapcore.ISO8601TimeEncoder
		encode.EncodeLevel = zapcore.CapitalLevelEncoder
		LauncherLogger = zap.New(zapcore.NewCore(zapcore.NewJSONEncoder(encode),
			zapcore.AddSync(LogFile),
			zapcore.ErrorLevel), zap.AddCaller())

		sugarLogger := LauncherLogger.Sugar()
		if args != nil {
			sugarLogger.Errorf(format, args)
		} else {
			sugarLogger.Errorf(format)
		}
		if LogtoConsole == true {
			if args != nil {
				PrintError(strings.Replace(
					strings.Replace(time.Now().Format("2006-01-02 15:04:05")+" | ",
						" ", "_", 1),
					":", "-", -1)+format, args)
			} else {
				PrintError(strings.Replace(
					strings.Replace(time.Now().Format("2006-01-02 15:04:05")+" | ",
						" ", "_", 1),
					":", "-", -1) + format)
			}
		}
	}
}

func LogDebug(format string, args ...interface{}) {
	if inited == false {
		PrintError("Logger not initialized!")
	} else {
		encode := zap.NewProductionEncoderConfig()
		encode.EncodeTime = zapcore.ISO8601TimeEncoder
		encode.EncodeLevel = zapcore.CapitalLevelEncoder
		LauncherLogger = zap.New(zapcore.NewCore(zapcore.NewJSONEncoder(encode),
			zapcore.AddSync(LogFile),
			zapcore.DebugLevel), zap.AddCaller())

		sugarLogger := LauncherLogger.Sugar()
		if args != nil {
			sugarLogger.Debugf(format, args)
		} else {
			sugarLogger.Debugf(format)
		}
		if LogtoConsole == true {
			if args != nil {
				PrintDebug(strings.Replace(
					strings.Replace(time.Now().Format("2006-01-02 15:04:05")+" | ",
						" ", "_", 1),
					":", "-", -1)+format, args)
			} else {
				PrintDebug(strings.Replace(
					strings.Replace(time.Now().Format("2006-01-02 15:04:05")+" | ",
						" ", "_", 1),
					":", "-", -1) + format)
			}
		}
	}
}

//Log:Love_皓月
