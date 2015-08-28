FROM golang:1.5

RUN \
	 wget http://repo.zabbix.com/zabbix/2.2/debian/pool/main/z/zabbix-release/zabbix-release_2.2-1+wheezy_all.deb && \
	 dpkg -i zabbix-release_2.2-1+wheezy_all.deb && \
	 apt-get update && \
	 apt-get install zabbix-agent

RUN echo "LoadModulePath=/usr/src/g2z/dummy" >> /etc/zabbix/zabbix_agentd.conf && \
    echo "LoadModule=dummy.so" >> /etc/zabbix/zabbix_agentd.conf

RUN \
	mkdir -p /go/src/github.com/cavaliercoder/ && \
	ln -s /usr/src/g2z /go/src/github.com/cavaliercoder/g2z
