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

FROM golang:1.6

# install Zabbix agent
RUN \
    wget -q http://repo.zabbix.com/zabbix/2.2/debian/pool/main/z/zabbix-release/zabbix-release_2.2-1+wheezy_all.deb && \
    dpkg -i zabbix-release_2.2-1+wheezy_all.deb && \
    apt-get update -y && \
    apt-get install -y zabbix-agent zabbix-get zabbix-sender && \
    mkdir /var/run/zabbix && \
    chown zabbix:zabbix /var/run/zabbix

# install utilities
RUN apt-get install -y vim strace lsof

# install zabbix_agent_bench
RUN \
    wget -q http://sourceforge.net/projects/zabbixagentbench/files/linux/zabbix_agent_bench-0.4.0.x86_64.tar.gz && \
    tar -xzvf zabbix_agent_bench-0.4.0.x86_64.tar.gz && \
    mv zabbix_agent_bench-0.4.0.x86_64/zabbix_agent_bench /usr/bin/zabbix_agent_bench

# install dummy module in Zabbix agent
RUN \
    echo "LoadModulePath=/usr/src/g2z/dummy" >> /etc/zabbix/zabbix_agentd.conf && \
    echo "LoadModule=dummy.so" >> /etc/zabbix/zabbix_agentd.conf

# install UserParameters for benchmarking
RUN \
    echo "UserParameter=up.ping,/bin/echo 1" >> /etc/zabbix/zabbix_agentd.conf && \
    echo "UserParameter=up.echo[*],/bin/echo \$1 \$2 \$3 \$4" >> /etc/zabbix/zabbix_agentd.conf && \
    echo "#!/usr/bin/perl -w\nprint \"1\\\\n\";\n" >> /usr/bin/perl_ping.pl && chmod 755 /usr/bin/perl_ping.pl && \
    echo "UserParameter=perl.ping,/usr/bin/perl_ping.pl" >> /etc/zabbix/zabbix_agentd.conf

# symlink g2z into GOPATH
RUN \
    mkdir -p /go/src/github.com/cavaliercoder/ && \
    ln -s /usr/src/g2z /go/src/github.com/cavaliercoder/g2z
