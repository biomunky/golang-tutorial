package main

import "fmt"

var a int
var (
	b int
	c float64
)

func odd_or_even(n int) {
	var mod = n % 2
	switch mod {
	case 0:
		fmt.Printf("%d is even\n", mod)
	case 1:
		fmt.Printf("%d is odd\n", mod)
	}
}

func main() {
	a = 10
	b = 20
	fmt.Println(a - b)
	c = 20.1
	fmt.Println(c * 2.0)

	s := "hello"
	sa := []rune(s)
	fmt.Println(sa)

	for i,c := range(s) {
		fmt.Println(i, " -> ", c)
	}


	for i := 0; i < 10; i++ {
		odd_or_even(i)
	}


}
