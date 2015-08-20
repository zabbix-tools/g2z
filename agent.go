package main

/*
// zabbix agent headers
#include <stdio.h>
#include <stdint.h>
#include "module.h"

// wrapper for get_rparam to ease the burden of indexing a **char in Go
static char *g2z_get_rparam(AGENT_REQUEST *request, int i) {
	return get_rparam(request, i);
}

*/
import "C"

type AgentRequest struct {
	Key    string
	Params []string
}

func marshallAgentRequest(r *C.AGENT_REQUEST) *AgentRequest {
	out := &AgentRequest{
		Key:    C.GoString(r.key),
		Params: make([]string, r.nparam),
	}

	for i := 0; i < int(r.nparam); i++ {
		out.Params[i] = C.GoString(C.g2z_get_rparam(r, C.int(i)))
	}

	return out
}

func uint64Response(result *C.AGENT_RESULT, val uint64) {
	result._type = C.AR_UINT64
	result.ui64 = C.zbx_uint64_t(val)
}

func doubleResponse(result *C.AGENT_RESULT, val float64) {
	result._type = C.AR_DOUBLE
	result.dbl = C.double(val)
}

func stringResponse(result *C.AGENT_RESULT, val string) {
	result._type = C.AR_STRING
	result.str = C.CString(val)
}
