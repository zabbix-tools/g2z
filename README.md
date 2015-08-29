# g2z - Zabbix module adapter for Go [![GoDoc](https://godoc.org/github.com/cavaliercoder/g2z?status.svg)](http://godoc.org/github.com/cavaliercoder/g2z)

This project provides [Go](https://golang.org/) bindings for creating native
Zabbix modules.

Zabbix modules are an effective way of extending the Zabbix agent and server to
monitor resources which are not natively supported by Zabbix. 

There are currently two ways to extend Zabbix:

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
written in an easier high-level language (Go), with all the dependencies bundled
into a standalone library file.

The findings of some performance tests are listed in [performance.md](performance.md).


## Requirements

* Go v1.5.0+
* Zabbix v2.2.0+
* C build tools (only tested on GCC)


## Installation

Once you have installed [Go lang](https://golang.org/doc/install), and
configured your `$GOPATH`, simply run:

	$ go get github.com/cavaliercoder/g2z


## Usage

Here's a quick high-level run down on how to create a Zabbix agent module. For 
further guidance, there is full API documentation available on
[godoc.org](http://godoc.org/github.com/cavaliercoder/g2z) and an
[example module](https://github.com/cavaliercoder/g2z/blob/master/dummy/dummy.go)
included in the g2z sources which implements the dummy C module published by
Zabbix.

To begin, create a mandatory `main()` entry point to your library and import
g2z. The `main()` function will never be called but is a requirement for
building shared libraries in Go.

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

```go
func Echo(request *g2z.AgentRequest) (string, error) {
    return strings.Join(request.Params, " "), nil
}

```

Create an `init()` function to register your functions as agent item keys. The
`init()` function is executed by the Go runtime when your module is loaded into
Zabbix via `dlopen()`. You should not execute any other calls in this function,
other than registering your items and init/uninit handlers.

```go
func init() {
    g2z.RegisterStringItem("go.echo", "Hello world!", Echo)
}

```

There are a few different item types you may register. Each requires an agent
item key name, some test parameters and a handler function for Zabbix to call
when it receives a request for the registered item key.

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
