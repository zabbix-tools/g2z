package main

/*
// zabbix agent headers
#include <stdint.h>
#include "module.h"

typedef int (*agent_item_callback)();

int zbx_module_route_item(AGENT_REQUEST *request, AGENT_RESULT *result);
*/
import "C"

import (
	"unsafe"
)

const (
	ModuleOK   = int(C.ZBX_MODULE_OK)
	ModuleFail = int(C.ZBX_MODULE_FAIL)
)

const (
	ReturnOK   = int(C.SYSINFO_RET_OK)
	ReturnFail = int(C.SYSINFO_RET_FAIL)
)

type Metric struct {
	Key        string
	TestParams string
	HasParams  bool
}

var Timeout int

//export zbx_module_api_version
func zbx_module_api_version() int {
	return C.ZBX_MODULE_API_VERSION_ONE
}

//export zbx_module_init
func zbx_module_init() int {
	Log(LogLevelInformation, "Hello world!")

	return ModuleOK
}

//export zbx_module_uninit
func zbx_module_uninit() C.int {
	return C.int(ModuleOK)
}

//export zbx_module_item_timeout
func zbx_module_item_timeout(timeout int) {
	Timeout = timeout
}

//export zbx_module_item_list
func zbx_module_item_list() *C.ZBX_METRIC {
	callback := C.agent_item_callback(unsafe.Pointer(C.zbx_module_route_item))

	return &[]C.ZBX_METRIC{
		{
			key:        C.CString("go"),
			flags:      C.CF_HAVEPARAMS,
			function:   callback,
			test_param: C.CString(""),
		},
		{},
	}[0]
}

//export route_item
func route_item(request *C.AGENT_REQUEST, result *C.AGENT_RESULT) C.int {
	r := marshallAgentRequest(request)

	for i, param := range r.Params {
		Log(LogLevelInformation, "Param %d: %s\n", i, param)
	}

	stringResponse(result, "Tits...Tits McGee")

	return C.SYSINFO_RET_OK
}
