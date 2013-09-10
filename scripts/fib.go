package main

import "fmt"

func fib() func() int {
	x, y := 0, 1
	return func() int {
		x, y = y, y+x
		return x
	}
}

func main() {
	f := fib()
	for i := 0; i < 20; i++ {
		fmt.Println(f())
	}
}
