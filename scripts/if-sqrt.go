package main

import (
        "fmt"
        "math"
)

func sqrt(x float64) float64 {
        if x < 0 {
                return sqrt(-x)
        }
        return math.Sqrt(x)
}

func main () {
        fmt.Println(sqrt(2), sqrt(-4))
}
