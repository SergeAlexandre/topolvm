package vslog

import (
	"fmt"
	"github.com/cybozu-go/log"
	"strings"
)

/*
 Our aim is to understand scheduler behavior. For this, we implement this logging wrapper to:
 - Isolate our stuff from underlying logging system.
 - Allow multiline output
 */

func Printf(format string, a ...interface{})  {
	lines := strings.Split(fmt.Sprintf(format, a...), "\n")
	for _, line :=  range lines {
		if line != "" {
			_ = log.Debug(line, nil)
		}
	}
}

func IsEnabled() bool {
	return log.Enabled(log.LvDebug)
}
