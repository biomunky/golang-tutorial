package main

import "fmt"

type Vertex struct {
        Lat, Long float64
}

//var m map[string]Vertex

func main() {

        m := make(map[string]Vertex)
        m["foo"] = Vertex{
                40.654,
                30.0001,
        }
        fmt.Println(m["foo"])

        _, ok := m["cat"]
        fmt.Println(ok)

}
