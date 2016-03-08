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

all: g2z dummy

g2z: g2z.go darwin.go doc.go linux.go log.go lld.go module.go router.go
	go build -x

dummy:
	cd dummy && $(MAKE)

clean:
	go clean -i -x
	rm -vf g2z.h

clean-dummy:
	cd dummy && $(MAKE) clean

clean-all: clean clean-dummy
	
docker-build:
	docker build \
		--build-arg "http_proxy=$(http_proxy)" \
		--build-arg "https_proxy=$(https_proxy)" \
		--build-arg "no_proxy=$(no_proxy)" \
		-t cavaliercoder/g2z .

docker-run: docker-build
	docker run --rm -it \
		-v $(PWD):/usr/src/g2z \
		-w /usr/src/g2z \
		--privileged \
		cavaliercoder/g2z

.PHONY: all g2z dummy clean clean-dummy clean-all docker-build docker-run
