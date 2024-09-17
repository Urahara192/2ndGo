package main

import "fmt"

func main() {
	var num1, num2 float64

	fmt.Print("Введите первое число: ")
	fmt.Scan(&num1)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&num2)

	average := (num1 + num2) / 2

	fmt.Printf("Среднее значение: (%.2f + %.2f) / 2 = %.2f\n", num1, num2, average)
}
