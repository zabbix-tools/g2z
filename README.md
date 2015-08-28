# g2z - Zabbix module adapter for Go

This project aims to provide [Go](https://golang.org/) bindings for creating
native Zabbix modules.

Zabbix modules are an effective way of extending the Zabbix agent and server to
monitor resources which are not natively support by Zabbix. 

There are currently two ways to to extend Zabbix:

* [User Parameters](https://www.zabbix.com/documentation/2.4/manual/config/items/userparameters)
* and [Loadable Modules](https://www.zabbix.com/documentation/2.4/manual/config/items/loadablemodules?s[]=module)

User Parameters simply map agent item keys to system calls. While this is by
far the easiest way to extend Zabbix, User Parameters require a process fork on
every call (a severe performance impact under load) and typically require a
script interpreter such as Perl or Ruby and their dependent framework modules.

Loadable Modules offer a significant performance increase (being native C
libraries) and *reduce* the overhead of dependencies. Unfortunately, modules
are rarely adopted because the effort and expertice required to write one in C
is a great deal more than writing a script in a higher level language like
Perl.

This project aims to deliver the best of both worlds; fast, native C libraries,
written in a easier high-level language (Go), with all the dependencies bundled
in to a standalone library file.

## Requirements

* Go v1.5.0+
* Zabbix v2.2.0+
* C build tools

## Usage

Here's a quick high-level run down on how to create a Zabbix agent module. An
[example module](https://github.com/cavaliercoder/g2z/blob/master/dummy/dummy.go)
is included with the g2z sources.

To begin, create a mandatory `main()` entry point to your library and import g2z:

```go
package main

import "github.com/cavaliercoder/g2z"

func main() {
	panic("THIS_SHOULD_NEVER_HAPPEN")
}

```

Write a Go function which accepts a `*g2z.AgentRequest` parameter and returns
either a `string`, `uint64`, `float64` or `g2z.DiscoveryData` as the first
parameter, and an `error` as the second return parameter.

Your functions signature must match one of:

* `g2z.StringItemHandlerFunc`
* `g2z.Uint64ItemHandlerFunc`
* `g2z.DoubleItemHandlerFunc`
* or `g2z.DiscoveryItemHandlerFunc`

```go
func Echo(request *g2z.AgentRequest) (string, error) {
	return strings.Join(request.Params, " "), nil
}

```

Create an `init()` function to register your functions as agent item keys:

```go
func init() {
	g2z.RegisterStringItem("go.echo", "Hello world!", Echo)
}

```

Compile your project with:

```bash
$ go build -buildmode=c-shared
```

Load your `.so` module file into Zabbix as per the
[Zabbix manual](https://www.zabbix.com/documentation/2.2/manual/config/items/loadablemodules#configuration_parameters).

Test your item keys with `zabbix_agentd -p` or
[zabbix_agent_bench](https://github.com/cavaliercoder/zabbix_agent_bench).

## License

g2z - Zabbix module adapter for Go
Copyright (C) 2015 - Ryan Armstrong <ryan@cavaliercoder.com>

This program is free software; you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation; either version 2 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, write to the Free Software
Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.
