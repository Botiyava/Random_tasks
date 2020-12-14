/*
https://www.codewars.com/kata/56eb0be52caf798c630013c0/
 *Description:
 *Friday 13th or Black Friday is considered as unlucky day. Calculate how many unlucky days are in the given year.
 *Find the number of Friday 13th in the given year.
 *Input: Year in Gregorian calendar as integer.
 *Output: Number of Black Fridays in the year as an integer.
 */
package kata
import (
	"fmt"
	"time"
)
func UnluckyDays(year int) int{
	var counter int
	for  i := 1; i <= 12; i++ {
		if time.Date(year, time.Month(i), 13, 0, 0, 0, 0, time.UTC).Weekday() == 5{
			counter++
			fmt.Printf("%v will be in %v\n", counter, time.Month(i))
		}
	}
	return counter
}
