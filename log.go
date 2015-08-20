package main

/*

#include "log.h"

// non-variadic wrapper for zabbix_log
static void g2z_log(int level, char *format) {
	return zabbix_log(level, format);
}

*/
import "C"

import (
	"fmt"
)

const (
	LogLevelEmpty       = int(C.LOG_LEVEL_EMPTY)
	LogLevelCritical    = int(C.LOG_LEVEL_CRIT)
	LogLevelError       = int(C.LOG_LEVEL_ERR)
	LogLevelWarning     = int(C.LOG_LEVEL_WARNING)
	LogLevelDebug       = int(C.LOG_LEVEL_DEBUG)
	LogLevelInformation = int(C.LOG_LEVEL_INFORMATION)
)

func Log(level int, format string, a ...interface{}) {
	C.g2z_log(C.int(level), C.CString(fmt.Sprintf(format, a...)))
}
