package main

import (
	"fmt"
	"net"
)

func main() {
	// LookupHost IPv4/IPv6
	iprecords, _ := net.LookupIP("sabio.inc")
	for _, ip := range iprecords {
		fmt.Println(ip)
	}

	// LookupCNAME
	cname, _ := net.LookupCNAME("m.facebook.com")
	fmt.Println(cname)

	// LookupMX
	mxrecords, _ := net.LookupMX("sabio.inc")
	for _, mx := range mxrecords {
		fmt.Println(mx.Host, mx.Pref)
	}

	// LookupNS
	nsrecords, _ := net.LookupNS("sabio.inc")
	for _, ns := range nsrecords {
		fmt.Println(ns.Host)
	}

	// LookupTXT
	txtrecords, _ := net.LookupTXT("sabio.inc")
	for x, txt := range txtrecords {
		x = x + 1
		fmt.Println("Text Record", x, ": ", txt)
	}

	// LookupAddr
	addrrecords, _ := net.LookupAddr("sabio.inc")
	for _, addr := range addrrecords {
		fmt.Println(addr)
	}

}
