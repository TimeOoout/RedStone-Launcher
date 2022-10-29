module RedStoneLauncher

go 1.19

require (
	github.com/fatih/color v1.13.0
	go.uber.org/zap v1.23.0
)


replace github.com/timeooout/RSL_Log => .\RSL_Log

require (
	github.com/mattn/go-colorable v0.1.9 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/sys v0.0.0-20210806184541-e5e7981a1069 // indirect
)
