package main

import "fmt"

type Person struct {
        Firstname string
        Surname string
}

func main() {
        p := Person{"Bio", "Munky"}
        q := &p
        q.Firstname = "Bio"
        fmt.Println(q.Surname + ", " + q.Firstname)

        pr := new(Person)
        pr.Firstname = "Mr"
        pr.Surname = "Smith"

        fmt.Println(pr.Firstname, pr.Surname)

}
