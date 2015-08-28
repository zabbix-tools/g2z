package g2z

/*
#include <stdint.h>
#include "module.h"
#include "log.h"

// entry point for all agent items
int zbx_module_route_item(AGENT_REQUEST *request, AGENT_RESULT *result)
{
	// call export go function g2z.route_item()
	return route_item(request,result);
}
*/
import "C"
