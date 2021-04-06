package main

import (
	"fmt"
	"strings" // Only required for the alternative method
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	// Simple solution:
	// return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2],ip[3])

	// More interesting (needlessly "complicated") solution:
	s := make([]string, len(ip))
	for i, val := range ip {
		s[i] = fmt.Sprint(int(val))
	}
	return strings.Join(s, ".")
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
