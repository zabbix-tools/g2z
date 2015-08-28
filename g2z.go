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

// g2z provides Go bindings for creating a native Zabbix module.
package g2z

// item is a go wrapper for a C.ZBX_METRIC
type item struct {
	// Key is the registered item key.
	Key string

	// TestParams is a comma separated list of parameters passed to the registered key when the
	// Zabbix agent executes a test run (via `zabbix_agentd -p`).
	TestParams string

	// Callback is the function to be called by route_item() when the registered key is queried by
	// Zabbix.
	Callback interface{}
}

// InitHandlerFunc is the function signature for functions which may be registered for callback
// when a Zabbix agent module is loaded when zbx_module_init() is called.
type InitHandlerFunc func() error

// UninitHandlerFunc is the function signature for functions which may be registered for callback
// when a Zabbix agent module is unloaded when zbx_module_uninit() is called.
type UninitHandlerFunc func() error

// StringItemHandlerFunc is the function signature for functions which may be registered as an
// item key handler and return a string value.
type StringItemHandlerFunc func(*AgentRequest) (string, error)

// Uint64ItemHandlerFunc is the function signature for functions which may be registered as an
// item key handler and return an unsigned integer value.
type Uint64ItemHandlerFunc func(*AgentRequest) (uint64, error)

// DoubleItemHandlerFunc is the function signature for functions which may be registered as an
// item key handler and return a double precision floating point integer value.
type DoubleItemHandlerFunc func(*AgentRequest) (float64, error)

// DiscoveryItemHandlerFunc is the function signature for functions which may be registered as an
// item key handler and return JSON encoded low-level discovery data.
type DiscoveryItemHandlerFunc func(*AgentRequest) (DiscoveryData, error)

// Timeout is set by Zabbix (via zbx_module_item_timeout()) to the number of seconds that all
// registered item handlers should obey as a strict timeout.
//
// This timeout is not enforced and should be voluntarily implemented.
var Timeout int

// initHandlers stores all registered InitHandlerFuncs
var initHandlers = []InitHandlerFunc{}

// uninitHandlers stores all registered UninitHandlerFuncs
var uninitHandlers = []UninitHandlerFunc{}

// itemHandlers stores a map of registered item keys to callback functions
var itemHandlers map[string]item = make(map[string]item, 0)

// RegisterInitHandler should be called from init() to register an InitHandlerFunc which will be
// called when Zabbix calls zbx_module_init().
func RegisterInitHandler(fn InitHandlerFunc) {
	LogDebugf("Registering init handler")
	initHandlers = append(initHandlers, fn)
}

// RegisterUninitHandler should be called from init() to register an UninitHandlerFunc which will
// be called when Zabbix calls zbx_module_uninit().
func RegisterUninitHandler(fn UninitHandlerFunc) {
	LogDebugf("Registering uninit handler")
	uninitHandlers = append(uninitHandlers, fn)
}

// registerItems registers an agent item key, its test parameters and a callback function with
// Zabbix.
func registerItem(key string, testParams string, callback interface{}) {
	k := item{
		Key:        key,
		TestParams: testParams,
		Callback:   callback,
	}

	itemHandlers[key] = k
}

// RegisterStringItem registers an agent item key, its test parameters and a callback function with
// Zabbix for a key with returns a string.
//
// This function should only be called from `init()`
func RegisterStringItem(key string, testParams string, callback StringItemHandlerFunc) {
	registerItem(key, testParams, callback)
}

// RegisterUint64Item registers an agent item key, its test parameters and a callback function with
// Zabbix for a key with returns an unsigned integer.
//
// This function should only be called from `init()`
func RegisterUint64Item(key string, testParams string, callback Uint64ItemHandlerFunc) {
	registerItem(key, testParams, callback)
}

// RegisterDoubleItem registers an agent item key, its test parameters and a callback function with
// Zabbix for a key with returns a double precision floating point integer.
//
// This function should only be called from `init()`
func RegisterDoubleItem(key string, testParams string, callback DoubleItemHandlerFunc) {
	registerItem(key, testParams, callback)
}

// RegisterDiscoveryItem registers an agent item key, its test parameters and a callback function with
// Zabbix for a key with returns DiscoveryData.
//
// This function should only be called from `init()`
func RegisterDiscoveryItem(key string, testParams string, callback DiscoveryItemHandlerFunc) {
	registerItem(key, testParams, callback)
}
