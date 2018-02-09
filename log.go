/*
 * g2z - Zabbix module adapter for Go
 * Copyright (C) 2015 - Ryan Armstrong <ryan@cavaliercoder.com>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 2 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.
 */

package g2z

/*
#include <stdlib.h>
#include "log.h"

// ignore missing symbol if not loaded via Zabbix
#pragma weak    __zbx_zabbix_log

// non-variadic wrapper for C.zabbix_log
static void g2z_log(int level, const char *format)
{
	void (*fptr)(int, const char*, ...);

	// check if zabbix_log() is resolvable
	if ((fptr = zabbix_log) != 0)
		(*fptr)(level, format);
}
*/
import "C"

import (
	"fmt"
	"unsafe"
)

// logf formats according to a format specifier and writes to the Zabbix log file.
func logf(level int, format string, a ...interface{}) {
	// TODO: Add runtime check for configured log level
	str := C.CString(fmt.Sprintf(format, a...))
	C.g2z_log(C.int(level), str)
	C.free(unsafe.Pointer(str))
}

// LogCriticalf formats according to a format specifier and writes to the Zabbix log file with a
// critical message.
func LogCriticalf(format string, a ...interface{}) {
	logf(C.LOG_LEVEL_CRIT, format, a...)
}

// LogErrorf formats according to a format specifier and writes to the Zabbix log file with an
// error message.
func LogErrorf(format string, a ...interface{}) {
	logf(C.LOG_LEVEL_ERR, format, a...)
}

// LogWarningf formats according to a format specifier and writes to the Zabbix log file with a
// warning message.
func LogWarningf(format string, a ...interface{}) {
	logf(C.LOG_LEVEL_WARNING, format, a...)
}

// LogDebugf formats according to a format specifier and writes to the Zabbix log file with a
// debug message.
func LogDebugf(format string, a ...interface{}) {
	logf(C.LOG_LEVEL_DEBUG, format, a...)
}

// LogInfof formats according to a format specifier and writes to the Zabbix log file with an
// informational message.
func LogInfof(format string, a ...interface{}) {
	logf(C.LOG_LEVEL_INFORMATION, format, a...)
}
