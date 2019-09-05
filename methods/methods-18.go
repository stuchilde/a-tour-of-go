/*
	Exercise: Stringers
	Make the IPAddr type implement fmt.Stringer to print the address as a dotted quad.

	For instance, IPAddr{1, 2, 3, 4} should print as "1.2.3.4".
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

func (ips IPAddr) String() string {
	s := make([]string, len(ips))
	for i := range ips {
		s[i] = strconv.Itoa(int(ips[i]))
	}
	return fmt.Sprintf("%s", strings.Join(s, "."))
}
func main() {
	hosts := map[string]IPAddr{
		"googleDNS": {127, 0, 0, 1},
		"loopback":  {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
