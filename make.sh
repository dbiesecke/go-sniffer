#!/bin/sh
CGO_LDFLAGS="-L../gobox/lib/libpcap/libpcap.a"
CGO_CFLAGS="-I. -I../gobox/lib/libpcap"
LIBPCAP_PATH=../gobox/lib/libpcap
LIBPCAP=../gobox/lib/libpcap/libpcap.a
CGO_ENABLED=1 go build  -tags "release pcap" -a -installsuffix cgo -o sniffer && ldd sniffer
#CGO_ENABLED=1 go build  -tags "release pcap"  -o sniffer main.go && ldd sniffer 