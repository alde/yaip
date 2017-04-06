package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/shiena/ansicolor"
)

var color = struct {
	grey, red, green, yellow, blue, magenta, cyan, white, reset string
}{
	grey:    "\x1b[91m",
	red:     "\x1b[91m",
	green:   "\x1b[92m",
	yellow:  "\x1b[93m",
	blue:    "\x1b[94m",
	magenta: "\x1b[95m",
	cyan:    "\x1b[96m",
	white:   "\x1b[97m",
	reset:   "\x1b[39m",
}

func main() {
	ifs, err := net.Interfaces()
	ipv6 := flag.Bool("ipv6", false, "include ipv6 addresses")
	flag.Parse()
	if err != nil {
		fmt.Fprint(os.Stderr, "Unable to list interfaces")
		os.Exit(2)
	}
	for _, netintf := range ifs {
		pretty(netintf, *ipv6)
	}
}

func pretty(ninf net.Interface, ipv6 bool) {
	addrs, _ := ninf.Addrs()
	var s []string
	for _, addr := range addrs {
		ip := net.ParseIP(addrToString(addr))
		if ip.To4() != nil {
			s = append(s, fmt.Sprintf("%s%s%s", color.cyan, ip, color.reset))
		} else if ip.To16() != nil && ipv6 {
			s = append(s, fmt.Sprintf("%s%s%s", color.blue, ip, color.reset))
		}
	}
	if len(s) == 0 {
		return
	}
	w := ansicolor.NewAnsiColorWriter(os.Stdout)
	fmt.Fprintf(w, "%s[%s %5s %s]%s ", color.green, color.yellow, ninf.Name, color.green, color.reset)
	fmt.Fprintf(w, "%s\n", strings.Join(s, " "))
}

func addrToString(addr net.Addr) string {
	return strings.Split(addr.String(), "/")[0]
}
