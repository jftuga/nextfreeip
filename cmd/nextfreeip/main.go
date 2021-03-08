package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

const pgmName string = "nextfreeip"
const pgmVersion string = "1.0.0"
const pgmURL string = "https://github.com/jftuga/nextfreeip"

func usage(pgm string) {
	fmt.Println()
	fmt.Printf("%s: Get the next consecutive IP address that is not found in DNS\n", pgmName)
	fmt.Printf("version: %s\n", pgmVersion)
	fmt.Println(pgmURL)
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Printf("%s [ ip-address ]\n", pgm)
	fmt.Println()
	fmt.Println("Note: The program stops searching after checking the `x.y.z.255` address.")
	fmt.Println()
}

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

// getLastOctet - return the last octet of an IP in numeric format
func getLastOctet(ip string) int {
	slots := strings.Split(ip, ".")
	last := slots[3]
	i, err := strconv.Atoi(last)
	if err != nil {
		log.Fatalf("Unable to convert string value to int: %s\n%s\n", last, err)
	}
	return i
}

// getTriplet - return the first 3 octets of and IP
func getTriplet(ip string) string {
	s := strings.Split(ip, ".")
	return s[0] + "." + s[1] + "." + s[2]
}

func main() {
	if len(os.Args) <= 1 {
		usage(os.Args[0])
		return
	}

	ip := os.Args[1]
	triplet := getTriplet(ip)
	last := getLastOctet(ip)

	found := false
	i := 0
	for i = last; i <= 255; i++ {
		found = resolveIP(fmt.Sprintf("%s.%d", triplet, i))
		if !found {
			fmt.Printf("%s.%d is not in DNS\n", triplet, i)
			return
		}
	}

	fmt.Println()
	fmt.Printf("Stopped searching at: %s.%d\n", triplet, i-1)
}
