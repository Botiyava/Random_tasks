/*
A bookseller has lots of books classified in 26 categories labeled A, B, ... Z. 
Each book has a code c of 3, 4, 5 or more characters. The  characters of a code is a capital letters which defines the book category.

In the bookseller's stocklist each code c is followed by a space and by a positive integer n (int n >= 0) 
which indicates the quantity of books of this code in stock.
Your task is to find all the books of L with codes belonging to each category of M and to sum their quantity according to each category. 
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func StockList(listArt []string, listCat []string) string {
result := map[string]int{}
for i := 0; i < len(listArt); i++{
	var quantityString string
	currentString := strings.Split(listArt[i], "")
	for i, elem := range currentString {
		if elem == " "{
			quantityString = strings.Join(currentString[i+1:], "")
			currentString = currentString[:i]
		}
	}
	quantity, err := strconv.Atoi(quantityString)
	if err != nil{
		fmt.Println(err)
	}
	for _, elem := range currentString{
		result[elem] += quantity
	}

}
var str string
for i, elem := range listCat{
	if i != len(listCat) - 1 {
		str += fmt.Sprintf("(%v : %v) - ", elem, result[elem])
	}else{
		str += fmt.Sprintf("(%v : %v)", elem, result[elem])
	}

}

return str
}
func main() {
	L := []string{"ABART 2000", "CDXEF 5000", "BKWRK 25", "BTSQZ 89", "DRTYM 60"}
	M := []string{"A", "B", "C", "W"}
	fmt.Println(StockList(L, M))


}
