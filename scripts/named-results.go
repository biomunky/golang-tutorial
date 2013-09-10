package main

import "fmt"

func halfAndDouble(number int) (half, double int) {
	half = number / 2
	double = number * number
	return
}

func main() {
	p1, p2 := halfAndDouble(20)
	fmt.Println(p1, p2)
}
