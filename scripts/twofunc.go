package main

import "fmt"

func swap(arg1 string, arg2 string) (string, string) {
        return arg2, arg1
}

func main() {
        a, b := swap("haz", "can")
        fmt.Println(a, b)
}
