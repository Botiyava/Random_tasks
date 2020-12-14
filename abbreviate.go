package main

import (
	"strings"
)

/*
 https://www.codewars.com/kata/57eadb7ecd143f4c9c0000a3
 *Write a function to convert a name into initials. This kata strictly takes two words with one space in between them.
 *The output should be two capital letters with a dot separating them.
 */

func AbbrevName(name string) string{
	nameLetter := strings.Split(name, "")
	for i, val := range nameLetter{
		if val == " "{
			nameLetter[0] = strings.ToUpper(nameLetter[0])
			nameLetter[i + 1] = strings.ToUpper(nameLetter[i + 1])
			nameLetter = append(nameLetter[:1],".")
			nameLetter = append(nameLetter ,nameLetter[i + 1:i + 2]...)
			//nameLetter = append(nameLetter, ".")  Why I don't need it?
		}
	}
	name = strings.Join(nameLetter, "")
	return name
}

