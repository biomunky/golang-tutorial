package main

import "fmt"

func main() {
        afunc := func(a, b int) int {
                return a + b
        }
        fmt.Println(afunc(2,2))

}
