module RedStoneLauncher

go 1.19

require (
	github.com/fatih/color v1.13.0 // indirect
	go.uber.org/zap v1.23.0 // indirect
)

require github.com/TimeOoout/RSL_Core v0.1.0

replace github.com/TimeOoout/RSL_Core => ./RSL_Core

require (
	github.com/mattn/go-colorable v0.1.9 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/sys v0.0.0-20210806184541-e5e7981a1069 // indirect
)
