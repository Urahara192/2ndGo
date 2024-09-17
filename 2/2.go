package main

import "fmt"

func main() {
	var a int
	fmt.Println("Введите число: ")
	fmt.Scan(&a)

	if a > 0 {
		fmt.Println("Positive")
	} else if a == 0 {
		fmt.Println("Zero")
	} else {
		fmt.Println("Negative")
	}
}
