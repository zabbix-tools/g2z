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
#include <stdio.h>
#include <stdint.h>
#include "module.h"

// wrapper for get_rparam macro to ease the burden of indexing a **char in go
static char *g2z_get_rparam(AGENT_REQUEST *request, int i) {
	return get_rparam(request, i);
}
*/
import "C"

import (
	"fmt"
)

// route_item is called by zbx_module_route_item(), the entrypoint for all registered items.
//
//export route_item
func route_item(request *C.AGENT_REQUEST, result *C.AGENT_RESULT) C.int {
	// marshall a C.AGENT_RESULT to g2z.AgentRequest
	req := &AgentRequest{
		Key:    C.GoString(request.key),
		Params: make([]string, request.nparam),
	}

	for i := 0; i < int(request.nparam); i++ {
		req.Params[i] = C.GoString(C.g2z_get_rparam(request, C.int(i)))
	}

	// get the item handler for the requested key
	item, ok := itemHandlers[req.Key]
	if !ok {
		// this should never happen
		LogCriticalf("Item appears to be registered but has no go handler: %s", req.Key)
		return C.SYSINFO_RET_FAIL
	}

	// call handler
	switch item.Callback.(type) {
	case StringItemHandlerFunc:
		if v, err := item.Callback.(StringItemHandlerFunc)(req); err != nil {
			setMessageResult(result, err.Error())
			return C.SYSINFO_RET_FAIL
		} else {
			result._type = C.AR_STRING
			result.str = C.CString(v)
		}

	case Uint64ItemHandlerFunc:
		if v, err := item.Callback.(Uint64ItemHandlerFunc)(req); err != nil {
			setMessageResult(result, err.Error())
			return C.SYSINFO_RET_FAIL
		} else {
			result._type = C.AR_UINT64
			result.ui64 = C.uint64_t(v)
		}

	case DoubleItemHandlerFunc:
		if v, err := item.Callback.(DoubleItemHandlerFunc)(req); err != nil {
			setMessageResult(result, err.Error())
			return C.SYSINFO_RET_FAIL
		} else {
			result._type = C.AR_DOUBLE
			result.dbl = C.double(v)
		}

	case DiscoveryItemHandlerFunc:
		if v, err := item.Callback.(DiscoveryItemHandlerFunc)(req); err != nil {
			setMessageResult(result, err.Error())
			return C.SYSINFO_RET_FAIL
		} else {
			result._type = C.AR_STRING
			result.str = C.CString(v.JSON())
		}
	}

	return C.SYSINFO_RET_OK
}

// setMessageResult adds an error message to an agent result struct
func setMessageResult(result *C.AGENT_RESULT, format string, a ...interface{}) {
	result._type = C.AR_MESSAGE
	result.msg = C.CString(fmt.Sprintf(format, a...))
}
