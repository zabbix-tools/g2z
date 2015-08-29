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

package main

import (
	"github.com/cavaliercoder/g2z"
	"runtime"
	"testing"
)

func TestInitModule(t *testing.T) {
	if err := InitModule(); err != nil {
		t.Errorf(err.Error())
	}
}

func TestUninitModule(t *testing.T) {
	if err := UninitModule(); err != nil {
		t.Error(err.Error())
	}
}

func TestPing(t *testing.T) {
	i, err := Ping(&g2z.AgentRequest{})
	if err != nil {
		t.Error(err.Error())
	}

	if i != 1 {
		t.Errorf("Expected Ping() to return 1, got %d", i)
	}
}

func TestEcho(t *testing.T) {
	r := &g2z.AgentRequest{
		Params: []string{
			"hello",
			"world",
		},
	}

	s, err := Echo(r)
	if err != nil {
		t.Error(err.Error())
	}

	if s != "hello world" {
		t.Errorf("Expected Echo() to return 'hello world', got: %s", s)
	}
}

func TestRandom(t *testing.T) {
	// test range between 1 and 2
	r := &g2z.AgentRequest{
		Params: []string{"1", "2"},
	}

	f, err := Random(r)

	// check for errors
	if err != nil {
		t.Error(err.Error())
	}

	// check value is in range
	if f < 1 || f > 2 {
		t.Errorf("Expected Random() to return a number between 1 and 2, got: %f", f)
	}

	// validate param count check
	r = &g2z.AgentRequest{}
	_, err = Random(r)
	if err == nil {
		t.Errorf("Expected Random() to fail without parameters")
	}

	// validate param range
	r = &g2z.AgentRequest{
		Params: []string{"2", "1"},
	}
	_, err = Random(r)
	if err == nil {
		t.Errorf("Expected Random() to fail with an invalid range")
	}
}

func TestDiscoverCpus(t *testing.T) {
	r := &g2z.AgentRequest{}
	d, err := DiscoverCpus(r)

	if err != nil {
		t.Error(err.Error())
	}

	if len(d) != runtime.NumCPU() {
		t.Errorf("Expected DiscoverCpus() to return %d CPUs, got %d", runtime.NumCPU(), len(d))
	}
}

func TestVersion(t *testing.T) {

	r := &g2z.AgentRequest{}
	s, err := Version(r)

	if err != nil {
		t.Error(err.Error())
	}

	if s != runtime.Version() {
		t.Errorf("Expected Version() to return '%s', got '%s'", runtime.Version(), s)
	}
}
