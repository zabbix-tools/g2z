/*
 * https://github.com/jbuberel/buildmodeshared
 * https://docs.google.com/document/d/1nr-TQHw_er6GOQRsF6T43GGhFDelrAP0NqSS_00RgZQ/edit#heading=h.fwmrrio0df0i
 */
package main

/*
// some symbols (within the Zabbix agent) won't resolve at link-time
// we can ignore these and resolve at runtime
#cgo LDFLAGS: -Wl,--unresolved-symbols=ignore-in-object-files

#include <stdint.h>

// zabbix agent headers
#include "module.h"
#include "log.h"

// non-variadic wrapper for zabbix_log
static inline void g2z_log(int level, char *format) {
	return zabbix_log(level, format);
}

*/
import "C"

import "fmt"

const (
	ModuleOK   = int(C.ZBX_MODULE_OK)
	ModuleFail = int(C.ZBX_MODULE_FAIL)
)

const (
	LogLevelEmpty       = int(C.LOG_LEVEL_EMPTY)
	LogLevelCritical    = int(C.LOG_LEVEL_CRIT)
	LogLevelError       = int(C.LOG_LEVEL_ERR)
	LogLevelWarning     = int(C.LOG_LEVEL_WARNING)
	LogLevelDebug       = int(C.LOG_LEVEL_DEBUG)
	LogLevelInformation = int(C.LOG_LEVEL_INFORMATION)
)

type Metric struct {
	Key        string
	TestParams string
	HasParams  bool
}

func main() {

}

//export zbx_module_api_version
func zbx_module_api_version() int {
	return C.ZBX_MODULE_API_VERSION_ONE
}

//export zbx_module_init
func zbx_module_init() int {
	Log(LogLevelInformation, "Hello world!")

	return ModuleOK
}

//export zbx_module_item_timeout
func zbx_module_item_timeout(timeout int) {

}

//export zbx_module_item_list
func zbx_module_item_list() *C.struct_ZBX_METRIC {
	return &C.struct_ZBX_METRIC{}
}

//export zbx_module_uninit
func zbx_module_uninit() int {
	return ModuleOK
}

func Log(level int, format string, a ...interface{}) {
	C.g2z_log(C.int(level), C.CString(fmt.Sprintf(format, a...)))
}
