package main

// #cgo CFLAGS:  -I../gobox/lib/libpcap -I../gobox/src/vendor/libpcap-libpcap-1.7.4/pcap/ -I../gobox/src/vendor/libpcap-libpcap-1.7.4/ -I../gobox/../../src/vendor/libpcap-libpcap-1.7.4/pcap/ -I../gobox/../../src/pcap/pcap
// #cgo LDFLAGS: -L../gobox/lib/libpcap/libpcap.a

import (
	"github.com/40t/go-sniffer/core"
)

func main() {
	core := core.New()
	core.Run()
}
