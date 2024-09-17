package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var input string

	fmt.Println("Введите строку: ")
	fmt.Scan(&input)

	length := utf8.RuneCountInString(input)
	fmt.Println(length)
}
