/*
 *
 *Given two arrays a and b write a function comp(a, b) (compSame(a, b) in Clojure)
 *that checks whether the two arrays have the "same" elements, with the same multiplicities.
 *"Same" means, here, that the elements in b are the elements in a squared, regardless of the order.
 *https://www.codewars.com/kata/550498447451fbbd7600041c
 */
package kata

import (
	"fmt"
	"reflect"
	"sort"
)

func Comp(array1 []int, array2 []int) bool {
	fmt.Println(array1)
	fmt.Println(array2)
	ok := true
	if len(array1) != len(array2) {
		ok = false
		return ok
	}
	if array1 == nil || array2 == nil {
		return false
	}
	if reflect.DeepEqual(array1, array2) {
		return true
	}
	//k
	for i := 0; i < len(array1); i++ {
		array1[i] *= array1[i]
	}
	sort.Ints(array1)
	sort.Ints(array2)
	if reflect.DeepEqual(array1, array2) {
		ok = true
	} else {
		ok = false
	}
	return ok
}
