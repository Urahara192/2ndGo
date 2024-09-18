package main

import "fmt"

type Rectangle struct {
	width  int
	height int
}

func main() {
	rectangle := Rectangle{5, 3}
	fmt.Println("Площадь прямоугольника: ", rectangle.width*rectangle.height)
}
