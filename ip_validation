//https://www.codewars.com/kata/515decfd9dcfc23bb6000006
//Write an algorithm that will identify valid IPv4 addresses in dot-decimal format. IPs should be considered valid if they consist of four octets, with values between 0 and 255, inclusive.
//
//Input to the function is guaranteed to be a single string.
//Examples
//
//Valid inputs:
//
//1.2.3.4
//123.45.67.89
//
//Invalid inputs:
//
//1.2.3
//1.2.3.4.5
//123.456.78.90
//123.045.067.089

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func is_valid_ip(ip string) bool {
	var ipIntMas []int
	ipStringMas := strings.Split(ip, ".")
	for _, elem := range ipStringMas {
		currentOctetNumber, err := strconv.Atoi(elem)

		if currentOctetNumber > 255 || currentOctetNumber < 0 || err != nil {
			return false
		}
		ipIntMas = append(ipIntMas, currentOctetNumber)
	}
	if len(ipIntMas) < 4 || len(ipIntMas) > 4 {
		return false
	}
	for i, _ := range ipStringMas {
		ipStringMas[i] = strconv.Itoa(ipIntMas[i])
	}
	result := strings.Join(ipStringMas, ".")

	if result != ip {
		return false
	}
	return true
}
func main() {

	input := "123.045.067.089"
	fmt.Println(is_valid_ip(input))

}
