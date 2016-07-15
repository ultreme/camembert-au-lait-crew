package log

import (
	"fmt"
	"io"
	"os"
)

type FmtAppLogger struct {
	outWriter io.Writer
	errWriter io.Writer
	debug     bool
}

func NewFmtAppLogger() FmtAppLogger {
	return FmtAppLogger{
		outWriter: os.Stdout,
		errWriter: os.Stderr,
		debug:     false,
	}
}

func (bl *FmtAppLogger) SetDebug(debug bool)        { bl.debug = debug }
func (bl *FmtAppLogger) SetOutWriter(out io.Writer) { bl.outWriter = out }
func (bl *FmtAppLogger) SetErrWriter(err io.Writer) { bl.errWriter = err }

func (bl *FmtAppLogger) Debug(args ...interface{}) {
	if bl.debug {
		fmt.Fprintln(bl.errWriter, args...)
	}
}

func (bl *FmtAppLogger) Debugf(fmtString string, args ...interface{}) {
	if bl.debug {
		fmt.Fprintf(bl.errWriter, fmtString+"\n", args...)
	}
}

func (bl *FmtAppLogger) Error(args ...interface{}) {
	fmt.Fprintln(bl.errWriter, args...)
}

func (bl *FmtAppLogger) Errorf(fmtString string, args ...interface{}) {
	fmt.Fprintf(bl.errWriter, fmtString+"\n", args...)
}

func (bl *FmtAppLogger) Fatal(args ...interface{}) {
	fmt.Fprintln(bl.errWriter, args...)
	os.Exit(-1)
}

func (bl *FmtAppLogger) Fatalf(fmtString string, args ...interface{}) {
	fmt.Fprintf(bl.errWriter, fmtString+"\n", args...)
	os.Exit(-1)
}

func (bl *FmtAppLogger) Info(args ...interface{}) {
	fmt.Fprintln(bl.outWriter, args...)
}

func (bl *FmtAppLogger) Infof(fmtString string, args ...interface{}) {
	fmt.Fprintf(bl.outWriter, fmtString+"\n", args...)
}

func (bl *FmtAppLogger) Warn(args ...interface{}) {
	fmt.Fprintln(bl.outWriter, args...)
}

func (bl *FmtAppLogger) Warnf(fmtString string, args ...interface{}) {
	fmt.Fprintf(bl.outWriter, fmtString+"\n", args...)
}
