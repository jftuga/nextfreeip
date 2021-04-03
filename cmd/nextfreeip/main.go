
/*
nextfreeip
-John Taylor
Mar 9 2021

Get the next consecutive IP address that is not found in DNS when given a CIDR address

Acknowledgments:
https://stackoverflow.com/a/60542265/452281

 */

package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

const pgmName string = "nextfreeip"
const pgmVersion string = "1.1.1"
const pgmURL string = "https://github.com/jftuga/nextfreeip"

func usage(pgm string) {
	fmt.Println()
	fmt.Printf("%s: Get the next consecutive IP address that is not found in DNS\n", pgmName)
	fmt.Printf("version: %s\n", pgmVersion)
	fmt.Println(pgmURL)
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Printf("%s [ cidr address ]\n", pgm)
	fmt.Println()
	fmt.Println("Example:")
	fmt.Printf("%s 192.168.1.4/27\n", pgm)
	fmt.Println()
	fmt.Println("Note: The program stops searching after checking the `x.y.z.255` address.")
	fmt.Println("      It assumes a /24 netmask when unspecified.")
	fmt.Println()
}

// resolveIP - given a IP address, return the hostname
func resolveIP(ip string) bool {
	allNames, err := net.LookupAddr(ip)
	if err != nil {
		if strings.Contains(err.Error(), "DNS name does not exist") {
			return false
		}
		log.Fatalf("Error while looking up: %s\n%s\n", ip, err)
	}
	for _, name := range allNames {
		fmt.Printf("%s\t%s\n", ip, name)
		return true
	}
	return false
}

// addressToByte - convert IP address to a byte array representation
func addressToByte(ip net.IP) []byte{
	octets := strings.Split(ip.String(),".")
	var b []byte
	for _, n := range octets {
		i, err := strconv.Atoi(n)
		if err != nil {
			log.Fatalf("Invalid IP address: %s\n%s\n", ip, err)
		}
		b = append(b,byte(i))
	}
	return b
}

// intToAddress - convert a integer into an IP address
func intToAddress(i uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, i)
	return ip
}

func main() {
	if len(os.Args) <= 1 {
		usage(os.Args[0])
		return
	}

	cidr := os.Args[1]
	if !strings.Contains(cidr,"/") {
		// default to /24 if no netmask is given
		cidr += "/24"
	}

	ip, ipv4Net, err := net.ParseCIDR(cidr)
	if err != nil {
		log.Fatalln(err)
	}

	onBoundary := false
	if net.IP.String(ipv4Net.IP) == net.IP.String(ip) {
		onBoundary = true
	}

	// given from command line
	first := addressToByte(ip)
	startIPv4Net := &net.IPNet{
		IP:   first,
		Mask: ipv4Net.Mask,
	}

	// convert starting IP to uint32
	start := binary.BigEndian.Uint32(startIPv4Net.IP)
	// convert IPNet struct mask and address to uint32
	mask := binary.BigEndian.Uint32(ipv4Net.Mask)
	// find the final address
	finish := (start & mask) | (mask ^ 0xffffffff)

	if onBoundary {
		start += 1
		fmt.Printf("%s\t%s\n", ip, "SKIPPED - Network Boundary")
	}

	// perform a DNS lookup on each IP address until the end of the netmask
	var i uint32
	for i = start; i <= finish; i++ {
		z := intToAddress(i).String()
		found := resolveIP(z)
		if !found {
			fmt.Println()
			fmt.Println(z + " is not in DNS")
			return
		}
	}

	fmt.Println()
	fmt.Printf("Stopped searching at: %s\n", intToAddress(i-1))
}
