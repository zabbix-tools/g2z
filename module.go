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
// zabbix agent headers
#include <stdlib.h>
#include <stdint.h>
#include "module.h"

// go binding for a pointer to an agent item callback
typedef int (*agent_item_handler)(AGENT_REQUEST*, AGENT_RESULT*);

// item callback router function defined in router.go
int route_item(AGENT_REQUEST *request, AGENT_RESULT *result);

// create new metric list
static ZBX_METRIC *new_metric_list(const size_t n) {
	return (ZBX_METRIC*)calloc(sizeof(ZBX_METRIC), n + 1);
}

// append a metric to a metric list
static void append_metric(ZBX_METRIC *list, const ZBX_METRIC *n) {
	while(NULL != list->key)
		list++;

	memcpy(list, n, sizeof(ZBX_METRIC));
}

*/
import "C"

import (
	"runtime"
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
	LogDebugf("Initializing g2z module")
	LogDebugf(" - runtime version:\t%s", runtime.Version())

	// call all registered init hanlders
	handlerCount := len(initHandlers)
	for i, handler := range initHandlers {
		LogDebugf("Calling registered zbx_module_init() handler %d of %d", i+1, handlerCount)
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
	handlerCount := len(uninitHandlers)
	for i, handler := range uninitHandlers {
		LogDebugf("Calling registered zbx_module_uninit() handler %d of %d", i+1, handlerCount)
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
	LogDebugf("Setting module timeout to %d seconds", timeout)
	Timeout = timeout
}

//export zbx_module_item_list
func zbx_module_item_list() *C.ZBX_METRIC {
	LogDebugf("Registering %d item handlers", len(itemHandlers))

	// route all item key calls through route_item()
	router := C.agent_item_handler(unsafe.Pointer(C.route_item))

	// create null-terminated array of C.ZBX_METRICS
	metrics := C.new_metric_list(C.size_t(len(itemHandlers))) // never freed
	for _, item := range itemHandlers {
		m := C.ZBX_METRIC{
			key:        C.CString(item.Key), // freed by Zabbix
			flags:      C.CF_HAVEPARAMS,
			function:   router,
			test_param: C.CString(item.TestParams), // freed by Zabbix
		} // freed by Go GC

		C.append_metric(metrics, &m)
	}

	return metrics
}
