# RSL_ Log Package Documentation

<img src="https://img.shields.io/badge/RedStone Skin-Launcher_Modules-red" /> <img src="https://img.shields.io/badge/MineCraft_Launcher-RedStone_Launcher-brightgreen" /> <img src="https://img.shields.io/badge/RSL_Log-1.0.0_beta-brightgreen" />

>

* A logging package based on Zap secondary encapsulation.
* It is mainly used for the development of RedStone Launcher, as well as other applications.

#### Latest Version: [v1.0.0_beta](#latest-version-100_beta)

> 中文文档见 [RSL_Log 中文文档](README.md)

### Index
- [RSL_ Log Package Documentation](#rsl_-log-package-documentation)
      - [Latest Version: v0.1.0_beta](#latest-version-v010_beta)
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

> Prompt: Although the function will return an error, the internal error of the log package will be printed automatically, and it is unnecessary to capture it again \
> (Of course, if necessary, capture it.)


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
### Others
> Print Log Package Version
* RSL_Log.LogVersion()
> Get Log Package Version
* RSL_Log.GetVersion()


## Variables
### About Paths
* LogFolder = "./LauncherLog/"
> This path is the default output path of the log file. If it does not exist, it will be created automatically.\
> Modify it to output the log file to a different path.

### Automation settings

* LogtoConsole = true
> Output the information written to the log file **to the console at the same time**.

## Examples
### Code
```
package main

import "github.com/TimeOoout/RSL_Log"

func main() {
	RSL_Log.InitLauncherLogger()
	RSL_Log.LogVersion()
	RSL_Log.ClearLogs()
}

```
### OutPut
```
...
[INFO] 2023-01-07_21-34-40 | Log    : Init Logger successfully!
[INFO] 2023-01-07_21-34-40 | Log    : RSL_Log version:  [v1.0.0_beta]
[INFO] 2023-01-07_21-34-40 | Log    : Logs cleared successfully!   
```
