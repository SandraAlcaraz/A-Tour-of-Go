//Make the IPAddr type implement fmt.Stringer to print the address as a dotted quad.

package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func (ip IPAddr) String() string {
	result := fmt.Sprintf("%v", ip[0])
	for _, x := range ip[1:] {
		result += fmt.Sprintf(".%v", x)
	}
	return result
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
