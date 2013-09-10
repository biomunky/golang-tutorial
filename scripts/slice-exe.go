package main

import (
        "code.google.com/p/go-tour/pic"
        "fmt"
)

func Pic(dx, dy int) [][]uint8 {
        pic := make([][]uint8, dx)
        for i := range pic {
                pic[i] = make([] uint8, dy)
                for j := range pic[i] {
                        pic[i][j] = uint8((i+j)/2)
                }
        }
        return pic
}

func main() {
        pic.Show(Pic)
}
