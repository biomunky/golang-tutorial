package main

import "fmt"

func main () {
        pow := [] int {1,2,3,4,5,6,7}

        for i, v := range pow {
                fmt.Printf("2**%d = %d\n", i+1, v*2)
        }
}
