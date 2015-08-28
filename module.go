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
// some symbols (within the Zabbix agent) won't resolve at link-time
// we can ignore these and resolve at runtime
#cgo LDFLAGS: -Wl,--unresolved-symbols=ignore-in-object-files

// zabbix agent headers
#include <stdint.h>
#include "module.h"

// go binding for a pointer to an agent item callback
typedef int (*agent_item_callback)();

// item callback router function defined in cfuncs.go
int zbx_module_route_item(AGENT_REQUEST *request, AGENT_RESULT *result);
*/
import "C"

import (
	"unsafe"
)

// A AgentRequest represents an agent item request received from a Zabbix server or perhaps
// from `zabbix_get`.
type AgentRequest struct {
	// Key is the requested agent key (E.g. `agent.version`).
	Key string

	// Params is a slice of strings containing each parameter passed in the agent request for the
	// specified key (E.g. `dummy.echo[<param0>,<param1>]`)
	Params []string
}

//export zbx_module_api_version
func zbx_module_api_version() int {
	return C.ZBX_MODULE_API_VERSION_ONE
}

//export zbx_module_init
func zbx_module_init() int {
	// call all registered init hanlders
	for _, handler := range initHandlers {
		LogDebugf("Calling registered zbx_module_init() handler")
		if err := handler(); err != nil {
			LogCriticalf("%s", err.Error())
			return C.ZBX_MODULE_FAIL
		}
	}

	return C.ZBX_MODULE_OK
}

//export zbx_module_uninit
func zbx_module_uninit() int {
	// call all registered uninit hanlders
	for _, handler := range uninitHandlers {
		LogDebugf("Calling registered zbx_module_uninit() handler")
		if err := handler(); err != nil {
			LogCriticalf("%s", err.Error())
			return C.ZBX_MODULE_FAIL
		}
	}

	return C.ZBX_MODULE_OK
}

//export zbx_module_item_timeout
func zbx_module_item_timeout(timeout int) {
	// set global timeout var
	Timeout = timeout
}

//export zbx_module_item_list
func zbx_module_item_list() *C.ZBX_METRIC {
	// route all item key calls through zbx_module_route_item() -> route_item()
	callback := C.agent_item_callback(unsafe.Pointer(C.zbx_module_route_item))

	// create null-terminated array of metrics
	i := 0
	metrics := make([]C.ZBX_METRIC, len(itemHandlers)+1)
	for _, item := range itemHandlers {
		metrics[i] = C.ZBX_METRIC{
			key:        C.CString(item.Key),
			flags:      C.CF_HAVEPARAMS,
			function:   callback,
			test_param: C.CString(item.TestParams),
		}
		i++
	}

	return &metrics[0]
}
