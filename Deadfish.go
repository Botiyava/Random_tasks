/*
https://www.codewars.com/kata/51e0007c1f9378fa810002a9/train/go
 *Write a simple parser that will parse and run Deadfish.
 *
 *Deadfish has 4 commands, each 1 character long:
 *
 *i increments the value (initially 0)
 *d decrements the value
 *s squares the value
 *o outputs the value into the return array
 *Invalid characters should be ignored.
 */
package kata

import (
	"fmt"
	"strings"
)

func Parse(data string) []int{
buf := 0
slice := []int{}
str := strings.Split(data, "")
for i := 0; i < len(str); i++{
	switch str[i]{
	case "i":
		buf += 1
	case "d":
		buf -= 1
	case "s":
		buf *= buf
	case "o":

		slice = append(slice, buf)
	}
}
	
return slice
}
