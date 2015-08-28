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

FROM golang:1.5

# install Zabbix agent
RUN \
	 wget http://repo.zabbix.com/zabbix/2.2/debian/pool/main/z/zabbix-release/zabbix-release_2.2-1+wheezy_all.deb && \
	 dpkg -i zabbix-release_2.2-1+wheezy_all.deb && \
	 apt-get update && \
	 apt-get install zabbix-agent

# install dummy module in Zabbix agent
RUN echo "LoadModulePath=/usr/src/g2z/dummy" >> /etc/zabbix/zabbix_agentd.conf && \
    echo "LoadModule=dummy.so" >> /etc/zabbix/zabbix_agentd.conf

# symlink g2z into GOPATH
RUN \
	mkdir -p /go/src/github.com/cavaliercoder/ && \
	ln -s /usr/src/g2z /go/src/github.com/cavaliercoder/g2z
