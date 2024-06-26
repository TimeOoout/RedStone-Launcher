# RSL_ Log 中文文档

<img src="https://img.shields.io/badge/RedStone Skin-Launcher_Modules-red" /> <img src="https://img.shields.io/badge/MineCraft_Launcher-RedStone_Launcher-brightgreen" /> <img src="https://img.shields.io/badge/RSL_Log-1.0.0_beta-brightgreen" />

>

* 一个基于Zap二次封装的日志记录包。
* 主要用于RedStone Launcher的开发

> 当然也可以用于开发其它应用。

#### 最新版本: [v1.0.0_beta](#latest-version-100_beta)

> For English documents, see [RSL_ Log Package Documentation](README_EN.md)

### 目录
- [RSL_ Log 中文文档](#rsl_-log-中文文档)
      - [最新版本: v0.1.0_beta](#最新版本-v010_beta)
    - [目录](#目录)
  - [函数](#函数)
    - [文件操作](#文件操作)
    - [日志操作](#日志操作)
    - [输出操作](#输出操作)
  - [变量](#变量)
    - [关于路径](#关于路径)
    - [自动化设置](#自动化设置)
  - [案例](#案例)
    - [代码](#代码)
    - [输出](#输出)

## 函数

> 提示：函数虽然会返回错误，但日志包的内部错误均会自行打印，无需再次捕捉 \
> （当然有需要的话捕捉未尝不可）

### 文件操作
* InitLauncherLogger()
> 在 **任何其它日志操作前**调用. \
> 只有当 **初始化成功**后其它日志操作才能被执行.
* ClearLogs() error
> 将除了最新三个日志文件外的其它日志文件（日志文件夹内以".log"结尾的文件）**全部删除**.

### 日志操作
> 输入的字符串将会以**不同的等级**被记录到日志中去.\
> (如果初始化成功的话)\
> 默认情况下，记录到日志中的信息**也会被输出到控制台**.\
> 可以通过后面调整参数的方法来防止将其输出到控制台.
* LogInfo(format string, args ...interface{})
* LogDebug(format string, args ...interface{})
* LogWarning(format string, args ...interface{})
* LogError(format string, args ...interface{})

### 输出操作
> 将输入的数据输出到控制台.\
> 它可以 **在未初始化的情况下**被调用.
* PrintInfo(format string, args ...interface{})
* PrintDebug(format string, args ...interface{})
* PrintWarning(format string, args ...interface{})
* PrintError(format string, args ...interface{})

### 其它
> 打印日志包版本
* RSL_Log.LogVersion()
> 获取日志包版本
* RSL_Log.GetVersion()


## 变量
### 关于路径
* LogFolder = "./LauncherLog/"
> 此路径是日志文件的默认输出路径。如果它不存在，将自动创建.\
> 修改它以将日志文件输出到其他路径.

### 自动化设置

* LogtoConsole = true
> 将记录的日志信息 **同时输出到控制台**.

## 案例
### 代码
```
package main

import "github.com/TimeOoout/RSL_Log"

func main() {
	RSL_Log.InitLauncherLogger()
	RSL_Log.LogVersion()
	RSL_Log.ClearLogs()
}

```
### 输出
```
...
[INFO] 2023-01-07_21-34-40 | Log    : Init Logger successfully!
[INFO] 2023-01-07_21-34-40 | Log    : RSL_Log version:  [v1.0.0_beta]
[INFO] 2023-01-07_21-34-40 | Log    : Logs cleared successfully!   
```