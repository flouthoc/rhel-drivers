package log

import (
	"fmt"
	stdlog "log"
	"os"
	"sync"
)

var (
	Quiet   bool
	Verbose bool
	Debug   bool
)

var configureStdLogOnce sync.Once

func configureStdLog() {
	stdlog.SetOutput(os.Stderr)
	stdlog.SetFlags(stdlog.LstdFlags | stdlog.Lshortfile)
}

func stdPrintf(format string, args ...any) {
	configureStdLogOnce.Do(configureStdLog)
	stdlog.Output(3, fmt.Sprintf(format, args...))
}

func Debugf(format string, args ...any) {
	if Debug {
		stdPrintf(format, args...)
	}
}

func Logf(format string, args ...any) {
	if Debug {
		stdPrintf(format, args...)
		return
	}
	if Quiet {
		return
	}
	if !Verbose {
		return
	}
	fmt.Fprintf(os.Stderr, format+"\n", args...)
}

func Infof(format string, args ...any) {
	if Debug {
		stdPrintf(format, args...)
		return
	}
	if Quiet {
		return
	}
	fmt.Fprintf(os.Stderr, format+"\n", args...)
}

func Warnf(format string, args ...any) {
	if Debug {
		stdPrintf(format, args...)
		return
	}
	if Quiet {
		return
	}
	fmt.Fprintf(os.Stderr, "WARNING: "+format+"\n", args...)
}

func Errorf(format string, args ...any) {
	if Debug {
		stdPrintf(format, args...)
		return
	}
	fmt.Fprintf(os.Stderr, "ERROR: "+format+"\n", args...)
}
