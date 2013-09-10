package main

import "fmt"

func main () {
        p := []int{2,3,5,6,7}
        fmt.Println("p == ", p)

        for i := 0; i < len(p); i++ {
                fmt.Printf("p[%d] == %d\n", i, p[i])
        }

        fmt.Println(">", p[0:len(p)/2])

        fmt.Println(">", p[2:])

        fmt.Println(">", p[:])
}
