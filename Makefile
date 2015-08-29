# g2z - Zabbix module adapter for Go
# Copyright (C) 2015 - Ryan Armstrong <ryan@cavaliercoder.com>
#
# This program is free software; you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation; either version 2 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with this program; if not, write to the Free Software
# Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.

all: dummy/dummy.so

dummy/dummy.so: g2z.go module.go lld.go log.go router.go log.h module.h zbxtypes.h dummy/dummy.go
	cd dummy && go build -x -buildmode=c-shared -o dummy.so

clean:
	rm -f g2z.so g2z.h
	cd dummy && $(MAKE) clean

docker-build:
	docker build -t cavaliercoder/g2z .

docker-run: docker-build
	docker run --rm -it \
		-p 6060:6060 \
		-p 10050:10050 \
		-v $(PWD):/usr/src/g2z \
		-w /usr/src/g2z \
		cavaliercoder/g2z

.PHONY: all clean docker-build docker-run
