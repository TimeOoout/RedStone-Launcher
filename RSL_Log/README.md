# RSL_ Log Package Documentation

<img src="https://img.shields.io/badge/RedStone Skin-Launcher_Modules-red" />
<img src="https://img.shields.io/badge/MineCraft_Launcher-RedStone_Launcher-brightgreen" />
<img src="https://img.shields.io/badge/RSL_Log-1.0.0_beta-brightgreen" />

>

* A logging package based on Zap secondary encapsulation.
* It is mainly used for the development of RedStone Launcher, as well as other applications.

#### Latest Version: [1.0.0_beta](#latest-version-100_beta)

> 中文文档见 [RSL_Log 中文文档](README_CH.md)

### Index
- [RSL_ Log Package Documentation](#rsl_-log-package-documentation)
      - [Latest Version: 1.0.0_beta](#latest-version-100_beta)
    - [Index](#index)
  - [Functions](#functions)
    - [File Operations](#file-operations)
    - [Log Operations](#log-operations)
    - [Output operations](#output-operations)
  - [Variables](#variables)
    - [About Paths](#about-paths)
    - [Automation settings](#automation-settings)
  - [Examples](#examples)
    - [Code](#code)
    - [OutPut](#output)

## Functions
### File Operations
* InitLauncherLogger()
> Called before **any log operation**. \
> Other operations can be performed only after **initialization is successful**.
* ClearLogs() error
> All log files except the last three will be deleted.\
> (Files ending in ". log" in the log folder)
### Log Operations
> The input string and parameters will be recorded in the log **at different levels**.\
> (If successfully initialized)\
> By default, the logged logs are also output to the **console**.\
> You can adjust the parameters to prevent them from being output to the console.
* LogInfo(format string, args ...interface{})
* LogDebug(format string, args ...interface{})
* LogWarning(format string, args ...interface{})
* LogError(format string, args ...interface{})
### Output operations
> Output the input string to the console.\
> It can be called at any time **without initialization**.
* PrintInfo(format string, args ...interface{})
* PrintDebug(format string, args ...interface{})
* PrintWarning(format string, args ...interface{})
* PrintError(format string, args ...interface{})



## Variables
### About Paths
* LogFolder = "./LauncherLog/"
> This path is the default output path of the log file. If it does not exist, it will be created automatically.\
> Modify it to output the log file to a different path.

### Automation settings
* AutoLogging = true
> Automatically output some logs included in the package. \
> **(Opening is not recommended in other projects)**
* AutoMsg = true
> Automatically output some information contained in the package. \
> **(As in the previous one, it is not recommended to open it in other projects)**
* LogtoConsole = true
> Output the information written to the log file **to the console at the same time**.

## Examples
### Code
```
package main

import "RedStoneLauncher/RSL_Log"

func main() {
	RSL_Log.InitLauncherLogger()
	a := RSL_Log.ClearLogs()
	if a != nil {
		println(a.Error())
	}
	RSL_Log.GetVersion()
}

```
### OutPut
```
...
[INFO] Year-Month-Date_Hour-Min-Sec | Init Logger successfully! 
[INFO] Year-Month-Date_Hour-Min-Sec | RSL_Log version:[1.0.0_beta]

```