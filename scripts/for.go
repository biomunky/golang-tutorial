package main

import "fmt"

func main() {
        sum := 0
        for i := 0; i < 10; i++ {
                sum ++
        }

        suma := 1
        for suma < 100 {
                suma += suma
        }

        fmt.Println(sum)
        fmt.Println(suma)
}
