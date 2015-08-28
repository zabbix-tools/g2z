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

// logf formats according to a format specifier and writes to the Zabbix log file.
func logf(level int, format string, a ...interface{}) {
	C.g2z_log(C.int(level), C.CString(fmt.Sprintf(format, a...)))
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
