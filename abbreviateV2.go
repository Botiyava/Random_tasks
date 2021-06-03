//Более сложная задача и более лакончиное решение к abbreviate.go
/*
Напишите программу, которая принимает на вход фразу и составляет аббревиатуру по первым буквам слов:

    Today I learned → TIL
    Высшее учебное заведение → ВУЗ
    Кот обладает талантом → КОТ

Если слово начинается не с буквы, игнорируйте его:

    Ар 2 Ди #2 → АД

Разделителями слов считаются только пробельные символы. Дефис, дробь и прочие можно не учитывать:

    Анна-Мария Волхонская → АВ

*/

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var phrase string
  fmt.Scan(&phrase)
	words := strings.Fields(phrase)
	var abbr []rune
	for _, word := range words {
		letter := []rune(word)[0]
		if !unicode.IsLetter(letter) {
			abbr = append(abbr, unicode.ToUpper(letter))
		}
	}

	fmt.Println(string(abbr))
}
