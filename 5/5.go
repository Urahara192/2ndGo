package main

import "fmt"

func main() {
	var width, height float64
	fmt.Println("Введите ширину: ")
	fmt.Scan(&width)
	fmt.Println("Введите высоту: ")
	fmt.Scan(&height)

	area := width * height

	fmt.Printf("Площадь прямоугольника: %.2f\n", area)
}
