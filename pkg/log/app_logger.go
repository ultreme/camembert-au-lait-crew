package log

import "io"

type AppLogger interface {
	SetDebug(bool)

	SetOutWriter(io.Writer)
	SetErrWriter(io.Writer)

	Debug(args ...interface{})
	Debugf(fmgString string, args ...interface{})

	Error(args ...interface{})
	Errorf(fmgString string, args ...interface{})

	Fatal(args ...interface{})
	Fatalf(fmgString string, args ...interface{})

	Info(args ...interface{})
	Infof(fmgString string, args ...interface{})

	Warn(args ...interface{})
	Warnf(fmgString string, args ...interface{})
}
