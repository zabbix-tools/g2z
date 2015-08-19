# g2z - Zabbix module adapter for Go

This project aims to provide a simple [Go](https://golang.org/) library which
may be used to create shared object library modules for Zabbix, written in Go.

__WARNING__ this project is not yet working/stable!

## Requirements

* Go v1.5.0+
* Zabbix v2.2.0+

## Usage

Import g2z into your Go projects with:

```go
import "g2z"
```

Compile your project with:

```bash
$ go build -x -buildmode=c-shared -o g2z.so
```

## License

Copyright 2015 Ryan Armstrong <ryan@cavaliercoder.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
