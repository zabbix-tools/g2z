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

type item struct {
	Key        string
	TestParams string
	Callback   interface{}
}

type InitHandlerFunc func() error
type UninitHandlerFunc func() error

type StringItemHandlerFunc func(*AgentRequest) (string, error)
type Uint64ItemHandlerFunc func(*AgentRequest) (uint64, error)
type DoubleItemHandlerFunc func(*AgentRequest) (float64, error)
type DiscoveryItemHandlerFunc func(*AgentRequest) (DiscoveryData, error)

var initHandlers = []InitHandlerFunc{}
var uninitHandlers = []UninitHandlerFunc{}
var itemHandlers map[string]item = make(map[string]item, 0)

func RegisterInitHandler(fn InitHandlerFunc) {
	LogDebugf("Registering init handler")
	initHandlers = append(initHandlers, fn)
}

func RegisterUninitHandler(fn UninitHandlerFunc) {
	LogDebugf("Registering uninit handler")
	uninitHandlers = append(uninitHandlers, fn)
}

func registerItem(key string, testParams string, callback interface{}) {
	k := item{
		Key:        key,
		TestParams: testParams,
		Callback:   callback,
	}

	itemHandlers[key] = k
}

func RegisterStringItem(key string, testParams string, callback StringItemHandlerFunc) {
	registerItem(key, testParams, callback)
}

func RegisterUint64Item(key string, testParams string, callback Uint64ItemHandlerFunc) {
	registerItem(key, testParams, callback)
}

func RegisterDoubleItem(key string, testParams string, callback DoubleItemHandlerFunc) {
	registerItem(key, testParams, callback)
}

func RegisterDiscoveryItem(key string, testParams string, callback DiscoveryItemHandlerFunc) {
	registerItem(key, testParams, callback)
}
